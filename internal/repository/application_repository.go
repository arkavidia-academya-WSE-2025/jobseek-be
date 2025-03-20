package repository

import (
	"errors"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	Repository[entity.Application]
	Log *logrus.Logger
}

func NewApplicationRepository(log *logrus.Logger) *ApplicationRepository {
	return &ApplicationRepository{
		Log: log,
	}
}
func (r *ApplicationRepository) FindById(db *gorm.DB, post *entity.Application, id string) error {
	return db.Preload("Job").Preload("JobSeeker").Where("id = ?", id).Take(post).Error
}

func (r *ApplicationRepository) VerifyApplicationOwnership(db *gorm.DB, ApplicationId string, userID string) error {
	var application entity.Application
	err := db.Where("id = ? AND job_seeker_id = ?", ApplicationId, userID).Take(&application).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrForbidden // Job exists but doesn't belong to the user
		}
		return fiber.ErrInternalServerError // Database error
	}
	return nil
}

func (r *ApplicationRepository) Search(db *gorm.DB, request *model.SearchApplicationRequest) ([]entity.Application, int64, error) {
	var Applications []entity.Application
	if err := db.Preload("Job").Preload("JobSeeker").Scopes(r.FilterApplication(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&Applications).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Application{}).Scopes(r.FilterApplication(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return Applications, total, nil
}

func (r *ApplicationRepository) FilterApplication(request *model.SearchApplicationRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if request.FullName != "" {
			tx = tx.Where("full_name ILIKE ?", "%"+request.FullName+"%")
		}
		if request.Address != "" {
			tx = tx.Where("address ILIKE ?", "%"+request.Address+"%")
		}
		if request.ApplicationStatus != "" {
			tx = tx.Where("application_status ILIKE ?", "%"+request.ApplicationStatus+"%")
		}
		return tx
	}
}
