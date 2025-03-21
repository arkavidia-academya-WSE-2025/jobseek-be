package model

import (
	"time"

	"github.com/google/uuid"
)

// MessageResponse is the response model for messages
type MessageResponse struct {
	ID         *uuid.UUID  `json:"id,omitempty"`
	Content    string      `json:"content,omitempty"`
	IsRead     bool        `json:"is_read,omitempty"`
	SenderID   *uuid.UUID  `json:"sender_id,omitempty"`
	ReceiverID *uuid.UUID  `json:"receiver_id,omitempty"`
	CreatedAt  *time.Time  `json:"created_at,omitempty"`
	Sender     *UserDetail `json:"sender,omitempty"`
	Receiver   *UserDetail `json:"receiver,omitempty"`
}

// UserDetail is a simplified user model for messages
type UserDetail struct {
	ID       *uuid.UUID `json:"id,omitempty"`
	Username string     `json:"username,omitempty"`
	Role     string     `json:"role,omitempty"`
}

// SendMessageRequest is the request model for sending a message
type SendMessageRequest struct {
	ReceiverID string `json:"receiver_id" validate:"required,uuid4"`
	Content    string `json:"content" validate:"required,min=1"`
}

// GetMessagesRequest is the request model for getting messages
type GetMessagesRequest struct {
	WithUserID string `json:"with_user_id" validate:"required,uuid4"`
	Page       int    `json:"page,omitempty"`
	Size       int    `json:"size,omitempty"`
}

// MarkAsReadRequest is the request model for marking a message as read
type MarkAsReadRequest struct {
	MessageID string `json:"message_id" validate:"required,uuid4"`
} 