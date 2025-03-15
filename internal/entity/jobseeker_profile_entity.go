package entity

import (
	"time"

	"github.com/google/uuid"
)

type JobseekerProfile struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	PhotoURL  string    `gorm:"column:photo_url"`
	Headline  string    `gorm:"column:headline"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// Foreign key relationship
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (j *JobseekerProfile) TableName() string {
	return "jobseeker_profiles"
}