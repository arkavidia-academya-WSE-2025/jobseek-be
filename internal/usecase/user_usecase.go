package usecase

import (
	"context"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/model/converter"
	"fp-academya-be/internal/repository"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

// func (c *UserUseCase) verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {

// }

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("validation error: %v", err)
		return nil, fiber.ErrBadRequest
	}
	total, err := c.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		c.Log.Warnf("Failed to count by email: %v", err)
		return nil, fiber.ErrInternalServerError
	}
	if total > 0 {
		c.Log.Warnf("email already exist")
		return nil, fiber.ErrConflict
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to hash password: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Email:    request.Email,
		Username: request.Username,
		Password: string(password),
		Role:     "user",
	}
	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed to create user to database: %v", err)
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %v", err)
		return nil, fiber.ErrInternalServerError
	}
	return converter.UserToResponse(user), nil
}
