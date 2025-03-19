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
	"gorm.io/gorm"
)

type JobUsecase struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	JobRepository *repository.JobRepository
}

func NewJobUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, jobRepository *repository.JobRepository) *JobUsecase {
	return &JobUsecase{
		DB:            db,
		Log:           log,
		Validate:      validate,
		JobRepository: jobRepository,
	}
}

func (c *JobUsecase) Create(ctx context.Context, request *model.JobRequest) (*model.JobResponse, error) {
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

	// Create job entity
	job := &entity.Job{
		RecruiterID:  *request.RecruiterID,
		Title:        request.Title,
		Description:  request.Description,
		Requirements: request.Requirements,
		Location:     request.Location,
		Salary:       request.Salary,
	}

	// Insert into DB
	if err := c.JobRepository.Create(tx, job); err != nil {
		c.Log.Warnf("Failed to create job: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Could not create job")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Return created job
	return converter.JobToResponse(job), nil
}

func (c *JobUsecase) Get(ctx context.Context, request *model.GetJobRequest) (*model.JobResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}
	job := new(entity.Job)
	if err := c.JobRepository.FindById(tx, job, request.ID); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobToResponse(job), nil

}

func (c *JobUsecase) Search(ctx context.Context, request *model.SearchJobRequest) ([]model.JobResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warnf("Invalid request body")
		return nil, 0, fiber.ErrBadRequest
	}
	jobs, total, err := c.JobRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search job")
		return nil, 0, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.JobResponse, len(jobs))
	for i, job := range jobs {
		responses[i] = *converter.JobToResponse(&job)
	}
	return responses, total, nil
}
