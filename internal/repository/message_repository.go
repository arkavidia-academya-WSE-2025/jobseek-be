package repository

import (
	"fp-academya-be/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageRepository struct {
	Repository[entity.Message]
	Log *logrus.Logger
}

func NewMessageRepository(log *logrus.Logger) *MessageRepository {
	return &MessageRepository{
		Log: log,
	}
}

// FindConversation gets messages between two users with pagination
func (r *MessageRepository) FindConversation(db *gorm.DB, userID, otherUserID uuid.UUID, page, size int) ([]*entity.Message, int64, error) {
	var messages []*entity.Message
	var total int64

	// Get total count of messages
	countQuery := db.Model(&entity.Message{}).
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", 
			userID, otherUserID, otherUserID, userID)
	
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query for messages with pagination
	offset := (page - 1) * size
	query := db.Preload("Sender").Preload("Receiver").
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", 
			userID, otherUserID, otherUserID, userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(size)
	
	if err := query.Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

// MarkAsRead marks a message as read
func (r *MessageRepository) MarkAsRead(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&entity.Message{}).Where("id = ?", id).Update("is_read", true).Error
}

// FindUnreadCount gets the count of unread messages for a user
func (r *MessageRepository) FindUnreadCount(db *gorm.DB, userID uuid.UUID) (int64, error) {
	var count int64
	err := db.Model(&entity.Message{}).
		Where("receiver_id = ? AND is_read = ?", userID, false).
		Count(&count).Error
	return count, err
} 