package usecase

import (
	"context"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/model/converter"
	"fp-academya-be/internal/repository"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	PostRepository *repository.PostRepository
}

func NewPostUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, postRepository *repository.PostRepository) *PostUseCase {
	return &PostUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		PostRepository: postRepository,
	}
}

func (c *PostUseCase) Create(ctx context.Context, request *model.PostRequest, userId uuid.UUID) (*model.PostReponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}
	if err := c.Validate.Var(userId, "required"); err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Create post entity
	post := &entity.Post{
		UserID:  userId,
		Title:   request.Title,
		Content: request.Content,
	}

	// Insert into DB
	if err := c.PostRepository.Create(tx, post); err != nil {
		c.Log.Warnf("Failed to create post: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Could not create post")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Return created post
	return converter.PostToResponse(post), nil
}
