package repository

import (
	"errors"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JobRepository struct {
	Repository[entity.Job]
	Log *logrus.Logger
}

func NewJobRepository(log *logrus.Logger) *JobRepository {
	return &JobRepository{
		Log: log,
	}
}
func (r *JobRepository) FindById(db *gorm.DB, post *entity.Job, id string) error {
	return db.Preload("Recruiter").Where("id = ?", id).Take(post).Error
}

func (r *JobRepository) VerifyJobOwnership(db *gorm.DB, jobID string, userID string) error {
	var job entity.Job
	err := db.Where("id = ? AND recruiter_id = ?", jobID, userID).Take(&job).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrForbidden // Job exists but doesn't belong to the user
		}
		return fiber.ErrInternalServerError // Database error
	}
	return nil
}

func (r *JobRepository) Search(db *gorm.DB, request *model.SearchJobRequest) ([]entity.Job, int64, error) {
	var jobs []entity.Job
	if err := db.Scopes(r.FilterJob(request)).Preload("Recruiter").Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&jobs).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Job{}).Scopes(r.FilterJob(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return jobs, total, nil
}

func (r *JobRepository) FilterJob(request *model.SearchJobRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if request.Title != "" {
			tx = tx.Where("title ILIKE ?", "%"+request.Title+"%")
		}
		if request.Description != "" {
			tx = tx.Where("description ILIKE ?", "%"+request.Description+"%")
		}
		if request.Requirements != "" {
			tx = tx.Where("requirements ILIKE ?", "%"+request.Requirements+"%")
		}
		if request.Location != "" {
			tx = tx.Where("location ILIKE ?", "%"+request.Location+"%")
		}
		if request.Salary != 0 {
			tx = tx.Where("salary = ?", request.Salary)
		}
		return tx
	}
}
