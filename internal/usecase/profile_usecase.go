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

type ProfileUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	JobseekerProfileRepo     *repository.JobseekerProfileRepository
	CompanyProfileRepo       *repository.CompanyProfileRepository
	UserRepository           *repository.UserRepository
}

func NewProfileUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	jobseekerProfileRepo *repository.JobseekerProfileRepository,
	companyProfileRepo *repository.CompanyProfileRepository,
	userRepository *repository.UserRepository,
) *ProfileUseCase {
	return &ProfileUseCase{
		DB:                       db,
		Log:                      log,
		Validate:                 validate,
		JobseekerProfileRepo:     jobseekerProfileRepo,
		CompanyProfileRepo:       companyProfileRepo,
		UserRepository:           userRepository,
	}
}

// GetJobseekerProfile gets a jobseeker profile by user ID
func (c *ProfileUseCase) GetJobseekerProfile(ctx context.Context, userID string) (*model.JobseekerProfileResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Parse UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find profile
	profile := new(entity.JobseekerProfile)
	if err := c.JobseekerProfileRepo.FindByUserID(tx, profile, userUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile if it doesn't exist
			profile.UserID = userUUID
			if err := c.JobseekerProfileRepo.Create(tx, profile); err != nil {
				c.Log.Warnf("Failed to create jobseeker profile: %+v", err)
				return nil, fiber.ErrInternalServerError
			}
		} else {
			c.Log.Warnf("Failed to find jobseeker profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobseekerProfileToResponse(profile), nil
}

// UpdateJobseekerProfile updates a jobseeker profile
func (c *ProfileUseCase) UpdateJobseekerProfile(ctx context.Context, userID string, request *model.UpdateJobseekerProfileRequest) (*model.JobseekerProfileResponse, error) {
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

	// Parse UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find profile
	profile := new(entity.JobseekerProfile)
	if err := c.JobseekerProfileRepo.FindByUserID(tx, profile, userUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile if it doesn't exist
			profile.UserID = userUUID
			profile.PhotoURL = request.PhotoURL
			profile.Headline = request.Headline
			if err := c.JobseekerProfileRepo.Create(tx, profile); err != nil {
				c.Log.Warnf("Failed to create jobseeker profile: %+v", err)
				return nil, fiber.ErrInternalServerError
			}
		} else {
			c.Log.Warnf("Failed to find jobseeker profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	} else {
		// Update profile
		profile.PhotoURL = request.PhotoURL
		profile.Headline = request.Headline
		if err := c.JobseekerProfileRepo.Update(tx, profile); err != nil {
			c.Log.Warnf("Failed to update jobseeker profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobseekerProfileToResponse(profile), nil
}

// GetCompanyProfile gets a company profile by user ID
func (c *ProfileUseCase) GetCompanyProfile(ctx context.Context, userID string) (*model.CompanyProfileResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Parse UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find profile
	profile := new(entity.CompanyProfile)
	if err := c.CompanyProfileRepo.FindByUserID(tx, profile, userUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile if it doesn't exist
			profile.UserID = userUUID
			profile.PhotoURL = ""
			profile.Description = ""
			if err := c.CompanyProfileRepo.Create(tx, profile); err != nil {
				c.Log.Warnf("Failed to create company profile: %+v", err)
				return nil, fiber.ErrInternalServerError
			}
		} else {
			c.Log.Warnf("Failed to find company profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CompanyProfileToResponse(profile), nil
}

// UpdateCompanyProfile updates a company profile
func (c *ProfileUseCase) UpdateCompanyProfile(ctx context.Context, userID string, request *model.UpdateCompanyProfileRequest) (*model.CompanyProfileResponse, error) {
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

	// Parse UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Log.Warnf("Invalid user ID: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find profile
	profile := new(entity.CompanyProfile)
	if err := c.CompanyProfileRepo.FindByUserID(tx, profile, userUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile if it doesn't exist
			profile.UserID = userUUID
			profile.PhotoURL = request.PhotoURL
			profile.Description = request.Description
			if err := c.CompanyProfileRepo.Create(tx, profile); err != nil {
				c.Log.Warnf("Failed to create company profile: %+v", err)
				return nil, fiber.ErrInternalServerError
			}
		} else {
			c.Log.Warnf("Failed to find company profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	} else {
		// Update profile
		profile.PhotoURL = request.PhotoURL
		profile.Description = request.Description
		if err := c.CompanyProfileRepo.Update(tx, profile); err != nil {
			c.Log.Warnf("Failed to update company profile: %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CompanyProfileToResponse(profile), nil
}