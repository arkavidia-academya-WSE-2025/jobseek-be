package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyProfile struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	PhotoURL    string    `gorm:"column:photo_url"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	// Foreign key relationship
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (c *CompanyProfile) TableName() string {
	return "company_profiles"
}