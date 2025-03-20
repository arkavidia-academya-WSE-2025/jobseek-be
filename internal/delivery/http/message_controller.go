package http

import (
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MessageController struct {
	Log     *logrus.Logger
	Usecase *usecase.MessageUseCase
}

func NewMessageController(usecase *usecase.MessageUseCase, logger *logrus.Logger) *MessageController {
	return &MessageController{
		Log:     logger,
		Usecase: usecase,
	}
}

// SendMessage sends a message to another user
func (c *MessageController) SendMessage(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)

	// Parse request body
	request := new(model.SendMessageRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Send message
	response, err := c.Usecase.SendMessage(ctx.UserContext(), auth.ID, request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to send message")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.MessageResponse]{Data: response})
}

// GetConversation gets messages between the current user and another user
func (c *MessageController) GetConversation(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)

	// Parse request body
	request := new(model.GetMessagesRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Set pagination parameters from query if provided
	if ctx.QueryInt("page", 0) > 0 {
		request.Page = ctx.QueryInt("page", 1)
	}
	if ctx.QueryInt("size", 0) > 0 {
		request.Size = ctx.QueryInt("size", 10)
	}

	// Get messages
	messages, paging, err := c.Usecase.GetConversation(ctx.UserContext(), auth.ID, request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to get conversation")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[[]*model.MessageResponse]{
		Data:   messages,
		Paging: paging,
	})
}

// MarkAsRead marks a message as read
func (c *MessageController) MarkAsRead(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)

	// Parse request body
	request := new(model.MarkAsReadRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Mark message as read
	success, err := c.Usecase.MarkAsRead(ctx.UserContext(), auth.ID, request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to mark message as read")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[bool]{Data: success})
}

// GetUnreadCount gets the count of unread messages for the current user
func (c *MessageController) GetUnreadCount(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)

	// Get unread count
	count, err := c.Usecase.GetUnreadCount(ctx.UserContext(), auth.ID)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to get unread count")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[int64]{Data: count})
} 