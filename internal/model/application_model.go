package model

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationResponse struct {
	ID                *uuid.UUID    `json:"id"`
	FullName          string        `json:"full_name"`
	Address           string        `json:"address"`
	ApplicationStatus string        `json:"application_status"`
	CVPath            string        `json:"cv_path,omitempty"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Job               *JobResponse  `json:"job"`
	JobSeeker         *UserResponse `json:"job_seeker"`
}

type ApplicationRequest struct {
	FullName    string     `json:"full_name" form:"full_name" validate:"required"`
	Address     string     `json:"address" form:"address" validate:"required"`
	CVPath      string     `json:"cv_path,omitempty" form:"cv_path"`
	JobID       *uuid.UUID `json:"job_id" form:"job_id" validate:"required,max=100"`
	JobSeekerID *uuid.UUID `json:"-" form:"-" validate:"required,max=100"`
}

type SearchApplicationRequest struct {
	FullName          string `json:"full_name,omitempty"`
	Address           string `json:"address,omitempty"`
	ApplicationStatus string `json:"application_status,omitempty" validate:"omitempty,oneof=pending accepted rejected"`
	Page              int    `json:"page,omitempty" validate:"omitempty,min=1"`
	Size              int    `json:"size,omitempty" validate:"omitempty,min=1,max=100"`
}

type GetApplicationRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type UpdateApplicationRequest struct {
	ID          string `json:"id" validate:"required,max=100"`
	FullName    string `json:"full_name,omitempty" form:"full_name"`
	Address     string `json:"address,omitempty" form:"address"`
	CVPath      string `json:"cv_path,omitempty" form:"cv_path"`
	JobSeekerID string `json:"-"`
}

type DeleteApplicationRequest struct {
	ID     string `json:"id" validate:"required,max=100"`
	UserID string `json:"-" validate:"required,max=100"`
}
