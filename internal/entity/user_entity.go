package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username  string    `gorm:"column:username;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	Role      string    `gorm:"column:role;"`
	IsPremium bool      `gorm:"column:is_premium;"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}

func (u *User) TableName() string {
	return "users"
}
