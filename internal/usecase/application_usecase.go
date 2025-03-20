package usecase

import (
	"context"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/model/converter"
	"fp-academya-be/internal/repository"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ApplicationUsecase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	ApplicationRepository *repository.ApplicationRepository
}

func NewApplicationUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, applicationRepository *repository.ApplicationRepository) *ApplicationUsecase {
	return &ApplicationUsecase{
		DB:                    db,
		Log:                   log,
		Validate:              validate,
		ApplicationRepository: applicationRepository,
	}
}

func (c *ApplicationUsecase) Create(ctx context.Context, request *model.ApplicationRequest) (*model.ApplicationResponse, error) {
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

	// Create application entity and populate it with request data
	application := &entity.Application{
		FullName:          request.FullName, // Note: Make sure field names match with entity
		Address:           request.Address,
		CVPath:            request.CVPath,
		ApplicationStatus: "pending", // Default status
		JobID:             *request.JobID,
		JobSeekerID:       *request.JobSeekerID,
	}

	// Insert into DB
	if err := c.ApplicationRepository.Create(tx, application); err != nil {
		c.Log.Warnf("Failed to create application: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Could not create application")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Return created application
	return converter.ApplicationToResponse(application), nil
}

func (c *ApplicationUsecase) Update(ctx context.Context, request *model.UpdateApplicationRequest) (*model.ApplicationResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Fetch application by ID
	application := new(entity.Application)
	if err := c.ApplicationRepository.FindById(tx, application, request.ID); err != nil {
		c.Log.Warnf("Failed to find application : %+v", err)
		return nil, fiber.ErrNotFound
	}

	// Update application fields if provided
	if request.FullName != "" {
		application.FullName = request.FullName
	}
	if request.Address != "" {
		application.Address = request.Address
	}
	if request.CVPath != "" {
		application.CVPath = request.CVPath
	}
	application.UpdatedAt = time.Now()

	// Save updated application
	if err := c.ApplicationRepository.Update(tx, application); err != nil {
		c.Log.Warnf("Failed to update application : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ApplicationToResponse(application), nil
}

func (c *ApplicationUsecase) Delete(ctx context.Context, request *model.DeleteApplicationRequest) (*model.ApplicationResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Verify application ownership
	if err := c.ApplicationRepository.VerifyApplicationOwnership(tx, request.ID, request.UserID); err != nil {
		c.Log.Warnf("Unauthorized attempt to delete application : %+v", err)
		return nil, err // Returns fiber.ErrForbidden if user is not the owner
	}

	// Find application by ID
	application := new(entity.Application)
	if err := c.ApplicationRepository.FindById(tx, application, request.ID); err != nil {
		c.Log.Warnf("Failed to find application : %+v", err)
		return nil, fiber.ErrNotFound
	}

	// Delete application
	if err := c.ApplicationRepository.Delete(tx, application); err != nil {
		c.Log.Warnf("Failed to delete application : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ApplicationToResponse(application), nil
}

func (c *ApplicationUsecase) Get(ctx context.Context, request *model.GetApplicationRequest) (*model.ApplicationResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}
	application := new(entity.Application)
	if err := c.ApplicationRepository.FindById(tx, application, request.ID); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ApplicationToResponse(application), nil

}

func (c *ApplicationUsecase) Search(ctx context.Context, request *model.SearchApplicationRequest) ([]model.ApplicationResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warnf("Invalid request body")
		return nil, 0, fiber.ErrBadRequest
	}
	applications, total, err := c.ApplicationRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search job")
		return nil, 0, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.ApplicationResponse, len(applications))
	for i, application := range applications {
		responses[i] = *converter.ApplicationToResponse(&application)
	}
	return responses, total, nil
}
