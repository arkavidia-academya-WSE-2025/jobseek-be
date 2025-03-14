package model

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Role      string     `json:"role,omitempty" validate:"required,oneof=admin job_seeker recruiter"`
	Token     string     `json:"token,omitempty"`
	IsPremium bool       `json:"is_premium,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required" max="255" json:"token"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
	Role     string `json:"role" validate:"required,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
