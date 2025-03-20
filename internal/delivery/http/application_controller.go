package http

import (
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ApplicationController struct {
	Log     *logrus.Logger
	Usecase *usecase.ApplicationUsecase
}

func NewApplicationController(usecase *usecase.ApplicationUsecase, logger *logrus.Logger) *ApplicationController {
	return &ApplicationController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *ApplicationController) Create(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)

	// Parse request body
	request := new(model.ApplicationRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Set JobSeeker ID
	parsedUUID, err := uuid.Parse(auth.ID)
	if err != nil {
		c.Log.Warnf("Invalid UUID format for user ID: %v", err)
		return fiber.ErrUnauthorized
	}
	request.JobSeekerID = &parsedUUID

	// Check if JobID is nil and parse it from form if needed
	if request.JobID == nil {
		jobIDStr := ctx.FormValue("job_id")
		if jobIDStr != "" {
			jobUUID, err := uuid.Parse(jobIDStr)
			if err != nil {
				c.Log.Warnf("Invalid UUID format for job_id: %v", err)
				return fiber.ErrBadRequest
			}
			request.JobID = &jobUUID
		}
	}

	// Create application
	appResponse, err := c.Usecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to create application")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.ApplicationResponse]{Data: appResponse})
}

func (c *ApplicationController) Get(ctx *fiber.Ctx) error {
	request := &model.GetApplicationRequest{
		ID: ctx.Params("id"),
	}
	response, err := c.Usecase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get post")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ApplicationResponse]{Data: response})
}

func (c *ApplicationController) List(ctx *fiber.Ctx) error {
	request := &model.SearchApplicationRequest{
		FullName:          ctx.Query("name"),
		Address:           ctx.Query("address"),
		ApplicationStatus: ctx.Query("status"),
		Page:              ctx.QueryInt("page", 1),
		Size:              ctx.QueryInt("size", 10),
	}
	responses, total, err := c.Usecase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search application")
		return err
	}
	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}
	return ctx.JSON(model.WebResponse[[]model.ApplicationResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *ApplicationController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	// Parse request body
	request := new(model.UpdateApplicationRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Get job ID from URL params
	applicationIdParam := ctx.Params("id")
	request.ID = applicationIdParam
	request.JobSeekerID = auth.ID

	// Call use case to update job
	updatedJob, err := c.Usecase.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(updatedJob)
}

func (c *ApplicationController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	request := &model.DeleteApplicationRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}
	response, err := c.Usecase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete post")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ApplicationResponse]{Data: response})
}
