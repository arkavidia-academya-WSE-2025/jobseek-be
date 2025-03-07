package usecase

import (
	"context"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/repository"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
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

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) *model.UserResponse {
}
