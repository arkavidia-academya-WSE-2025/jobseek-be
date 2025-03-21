package usecase

import (
	"context"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/model/converter"
	"fp-academya-be/internal/repository"
	"math"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	MessageRepository *repository.MessageRepository
	UserRepository    *repository.UserRepository
}

func NewMessageUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	messageRepository *repository.MessageRepository,
	userRepository *repository.UserRepository,
) *MessageUseCase {
	return &MessageUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		MessageRepository: messageRepository,
		UserRepository:    userRepository,
	}
}

// SendMessage sends a message from one user to another
func (c *MessageUseCase) SendMessage(ctx context.Context, senderID string, request *model.SendMessageRequest) (*model.MessageResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Parse UUIDs
	senderUUID, err := uuid.Parse(senderID)
	if err != nil {
		c.Log.Warnf("Invalid sender ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	receiverUUID, err := uuid.Parse(request.ReceiverID)
	if err != nil {
		c.Log.Warnf("Invalid receiver ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Check if receiver exists
	receiver := new(entity.User)
	if err := c.UserRepository.FindById(tx, receiver, receiverUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Log.Warnf("Receiver not found: %+v", err)
			return nil, fiber.ErrNotFound
		}
		c.Log.Warnf("Failed to find receiver: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Create message
	message := &entity.Message{
		SenderID:   senderUUID,
		ReceiverID: receiverUUID,
		Content:    request.Content,
		IsRead:     false,
	}

	if err := c.MessageRepository.Create(tx, message); err != nil {
		c.Log.Warnf("Failed to create message: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Fetch complete message with sender and receiver details
	if err := tx.Preload("Sender").Preload("Receiver").First(message, message.ID).Error; err != nil {
		c.Log.Warnf("Failed to load message relationships: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.MessageToResponse(message), nil
}

// GetConversation gets messages between the current user and another user
func (c *MessageUseCase) GetConversation(ctx context.Context, userID string, request *model.GetMessagesRequest) ([]*model.MessageResponse, *model.PageMetadata, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request: %+v", err)
		return nil, nil, fiber.ErrBadRequest
	}

	// Set default pagination values if not provided
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Size <= 0 {
		request.Size = 10
	}

	// Parse UUIDs
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, nil, fiber.ErrBadRequest
	}

	otherUserUUID, err := uuid.Parse(request.WithUserID)
	if err != nil {
		c.Log.Warnf("Invalid other user ID: %+v", err)
		return nil, nil, fiber.ErrBadRequest
	}

	// Check if other user exists
	otherUser := new(entity.User)
	if err := c.UserRepository.FindById(tx, otherUser, otherUserUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Log.Warnf("Other user not found: %+v", err)
			return nil, nil, fiber.ErrNotFound
		}
		c.Log.Warnf("Failed to find other user: %+v", err)
		return nil, nil, fiber.ErrInternalServerError
	}

	// Get messages
	messages, total, err := c.MessageRepository.FindConversation(tx, userUUID, otherUserUUID, request.Page, request.Size)
	if err != nil {
		c.Log.Warnf("Failed to find messages: %+v", err)
		return nil, nil, fiber.ErrInternalServerError
	}

	// Auto-mark messages as read
	for _, message := range messages {
		if message.ReceiverID == userUUID && !message.IsRead {
			message.IsRead = true
			if err := c.MessageRepository.Update(tx, message); err != nil {
				c.Log.Warnf("Failed to mark message as read: %+v", err)
				// Continue anyway, not critical
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, nil, fiber.ErrInternalServerError
	}

	// Create pagination metadata
	totalPages := int64(math.Ceil(float64(total) / float64(request.Size)))
	metadata := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: totalPages,
	}

	return converter.MessagesToResponses(messages), metadata, nil
}

// MarkAsRead marks a message as read
func (c *MessageUseCase) MarkAsRead(ctx context.Context, userID string, request *model.MarkAsReadRequest) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return false, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request: %+v", err)
		return false, fiber.ErrBadRequest
	}

	// Parse UUIDs
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return false, fiber.ErrBadRequest
	}

	messageUUID, err := uuid.Parse(request.MessageID)
	if err != nil {
		c.Log.Warnf("Invalid message ID: %+v", err)
		return false, fiber.ErrBadRequest
	}

	// Check if message exists and belongs to the user
	message := new(entity.Message)
	if err := c.MessageRepository.FindById(tx, message, messageUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Log.Warnf("Message not found: %+v", err)
			return false, fiber.ErrNotFound
		}
		c.Log.Warnf("Failed to find message: %+v", err)
		return false, fiber.ErrInternalServerError
	}

	// Check if user is the receiver
	if message.ReceiverID != userUUID {
		c.Log.Warnf("User is not the receiver of this message")
		return false, fiber.ErrForbidden
	}

	// Mark as read
	if err := c.MessageRepository.MarkAsRead(tx, messageUUID); err != nil {
		c.Log.Warnf("Failed to mark message as read: %+v", err)
		return false, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

// GetUnreadCount gets the count of unread messages for a user
func (c *MessageUseCase) GetUnreadCount(ctx context.Context, userID string) (int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return 0, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Parse UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return 0, fiber.ErrBadRequest
	}

	// Get unread count
	count, err := c.MessageRepository.FindUnreadCount(tx, userUUID)
	if err != nil {
		c.Log.Warnf("Failed to get unread count: %+v", err)
		return 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return 0, fiber.ErrInternalServerError
	}

	return count, nil
} 