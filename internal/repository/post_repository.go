package repository

import (
	"fp-academya-be/internal/entity"

	"github.com/sirupsen/logrus"
)

type PostRepository struct {
	Repository[entity.Post]
	Log *logrus.Logger
}

func NewPostRepository(log *logrus.Logger) *PostRepository {
	return &PostRepository{
		Log: log,
	}
}
