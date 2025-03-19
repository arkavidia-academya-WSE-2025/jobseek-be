package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

func JobToResponse(job *entity.Job) *model.JobResponse {
	return &model.JobResponse{
		ID:           &job.ID,
		Title:        job.Title,
		Description:  job.Description,
		Requirements: job.Requirements,
		Location:     job.Location,
		Salary:       job.Salary,
		CreatedAt:    &job.CreatedAt,
		UpdatedAt:    &job.UpdatedAt,
		Recruiter:    UserToResponse(&job.Recruiter),
	}
}
