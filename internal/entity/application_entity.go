package entity

import (
	"time"

	"github.com/google/uuid"
)

// ApplicationStatus defines the enum for application status
type ApplicationStatus string

const (
	Pending  ApplicationStatus = "pending"
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
)

type Application struct {
	ID                uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName          string            `gorm:"column:full_name;not null"`
	Address           string            `gorm:"column:address;not null"`
	ApplicationStatus ApplicationStatus `gorm:"column:application_status;type:app_status;default:'pending'"`
	CVPath            string            `gorm:"column:cv_path"`
	CreatedAt         time.Time         `gorm:"column:created_at;default:now()"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;default:now()"`
	JobID             uuid.UUID         `gorm:"column:job_id;not null"`
	JobSeekerID       uuid.UUID         `gorm:"column:job_seeker_id;not null"`
	//foreign key
	Job       Job  `gorm:"foreignKey:JobID;references:ID;constraint:OnDelete:CASCADE"`
	JobSeeker User `gorm:"foreignKey:JobSeekerID;references:ID;constraint:OnDelete:CASCADE"`
}
