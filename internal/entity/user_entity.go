package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username  string    `gorm:"column:username;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	Role      string    `gorm:"column:role;"`
	Token     string    `gorm:"column:token"`
	CreatedAt int64     `gorm:"column:created_at;"`
	UpdatedAt int64     `gorm:"column:updated_at;"`
}

func (u *User) TableName() string {
	return "users"
}
