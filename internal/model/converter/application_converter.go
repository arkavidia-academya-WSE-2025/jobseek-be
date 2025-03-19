package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

func ApplicationToResponse(application *entity.Application) *model.ApplicationResponse {
	return &model.ApplicationResponse{
		ID:                &application.ID,
		FullName:          application.FullName,
		Address:           application.Address,
		ApplicationStatus: string(application.ApplicationStatus),
		CVPath:            application.CVPath,
		CreatedAt:         application.CreatedAt,
		UpdatedAt:         application.UpdatedAt,
		Job:               JobToResponse(&application.Job),
		JobSeeker:         UserToResponse(&application.JobSeeker),
	}
}
