package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Content    string    `gorm:"column:content;not null"`
	IsRead     bool      `gorm:"column:is_read;default:false"`
	SenderID   uuid.UUID `gorm:"type:uuid;not null;column:sender_id"`
	ReceiverID uuid.UUID `gorm:"type:uuid;not null;column:receiver_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	// Foreign key relationships
	Sender   User `gorm:"foreignKey:SenderID;references:ID;constraint:OnDelete:CASCADE"`
	Receiver User `gorm:"foreignKey:ReceiverID;references:ID;constraint:OnDelete:CASCADE"`
}

func (m *Message) TableName() string {
	return "messages"
} 