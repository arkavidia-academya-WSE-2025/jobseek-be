package model

import (
	"time"

	"github.com/google/uuid"
)

type JobResponse struct {
	ID           *uuid.UUID    `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Requirements string        `json:"requirements,omitempty"`
	Location     string        `json:"location,omitempty"`
	Salary       int           `json:"salary,omitempty"`
	CreatedAt    *time.Time    `json:"created_at,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	Recruiter    *UserResponse `json:"recruiter,omitempty"`
}

type JobRequest struct {
	RecruiterID  *uuid.UUID `json:"-" validate:"required,max=100"`
	Title        string     `json:"title" validate:"required,max=100"`
	Description  string     `json:"description" validate:"required,max=1000"`
	Requirements string     `json:"requirements" validate:"required,max=1000"`
	Location     string     `json:"location" validate:"required,max=100"`
	Salary       int        `json:"salary,omitempty"`
}

type SearchJobRequest struct {
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int    `json:"salary,omitempty"`
	Page         int    `json:"page,omitempty" validate:"min=1"`
	Size         int    `json:"size,omitempty" validate:"min=1,max=100"`
}

type GetJobRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type UpdateJobRequest struct {
	ID           string `json:"id" validate:"required,max=100"`
	UserID       string `json:"-" validate:"required,max=100"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int    `json:"salary,omitempty"`
}

type DeleteJobRequest struct {
	ID     string `json:"id" validate:"required,max=100"`
	UserID string `json:"-" validate:"required,max=100"`
}
