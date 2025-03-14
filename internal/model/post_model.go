package model

import (
	"time"

	"github.com/google/uuid"
)

type PostReponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	UserId    *uuid.UUID `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
