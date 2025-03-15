package model

import (
	"time"

	"github.com/google/uuid"
)

// JobseekerProfileResponse is the response model for jobseeker profiles
type JobseekerProfileResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	PhotoURL  string     `json:"photo_url,omitempty"`
	Headline  string     `json:"headline,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// UpdateJobseekerProfileRequest is the request model for updating jobseeker profiles
type UpdateJobseekerProfileRequest struct {
	PhotoURL string `json:"photo_url" validate:"omitempty,url"`
	Headline string `json:"headline" validate:"omitempty,max=255"`
}

// CompanyProfileResponse is the response model for company profiles
type CompanyProfileResponse struct {
	ID          *uuid.UUID `json:"id,omitempty"`
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	PhotoURL    string     `json:"photo_url,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// UpdateCompanyProfileRequest is the request model for updating company profiles
type UpdateCompanyProfileRequest struct {
	PhotoURL    string `json:"photo_url" validate:"omitempty,url"`
	Description string `json:"description" validate:"omitempty"`
}

// GetProfileRequest is used to request a profile by user ID
type GetProfileRequest struct {
	UserID string `json:"user_id" validate:"required"`
}