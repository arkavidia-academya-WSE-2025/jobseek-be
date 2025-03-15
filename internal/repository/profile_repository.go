package repository

import (
	"fp-academya-be/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JobseekerProfileRepository struct {
	Repository[entity.JobseekerProfile]
	Log *logrus.Logger
}

func NewJobseekerProfileRepository(log *logrus.Logger) *JobseekerProfileRepository {
	return &JobseekerProfileRepository{
		Log: log,
	}
}

func (r *JobseekerProfileRepository) FindByUserID(db *gorm.DB, profile *entity.JobseekerProfile, userID uuid.UUID) error {
	return db.Where("user_id = ?", userID).First(profile).Error
}

type CompanyProfileRepository struct {
	Repository[entity.CompanyProfile]
	Log *logrus.Logger
}

func NewCompanyProfileRepository(log *logrus.Logger) *CompanyProfileRepository {
	return &CompanyProfileRepository{
		Log: log,
	}
}

func (r *CompanyProfileRepository) FindByUserID(db *gorm.DB, profile *entity.CompanyProfile, userID uuid.UUID) error {
	return db.Where("user_id = ?", userID).First(profile).Error
}