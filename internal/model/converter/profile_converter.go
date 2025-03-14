package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

// JobseekerProfileToResponse converts a jobseeker profile entity to response
func JobseekerProfileToResponse(profile *entity.JobseekerProfile) *model.JobseekerProfileResponse {
	return &model.JobseekerProfileResponse{
		ID:        &profile.ID,
		UserID:    &profile.UserID,
		PhotoURL:  profile.PhotoURL,
		Headline:  profile.Headline,
		CreatedAt: &profile.CreatedAt,
		UpdatedAt: &profile.UpdatedAt,
	}
}

// CompanyProfileToResponse converts a company profile entity to response
func CompanyProfileToResponse(profile *entity.CompanyProfile) *model.CompanyProfileResponse {
	return &model.CompanyProfileResponse{
		ID:          &profile.ID,
		UserID:      &profile.UserID,
		PhotoURL:    profile.PhotoURL,
		Description: profile.Description,
		CreatedAt:   &profile.CreatedAt,
		UpdatedAt:   &profile.UpdatedAt,
	}
}