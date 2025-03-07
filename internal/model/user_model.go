package model

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID `json:"id",omitempty`
	Username  string    `json:"username",omitempty`
	Email     string    `json:"email",omitempty`
	Token     string    `json:"token",omitempty`
	CreatedAt int64     `json:"created_at",omitempty`
	UpdatedAt int64     `json:"updated_at",omitempty`
}

type VerifyUserRequest struct {
	Token string `validate:"required" max="255" json:"token"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
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
