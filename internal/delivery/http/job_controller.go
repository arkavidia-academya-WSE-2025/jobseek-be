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

type JobController struct {
	Log     *logrus.Logger
	Usecase *usecase.JobUsecase
}

func NewJobController(usecase *usecase.JobUsecase, logger *logrus.Logger) *JobController {
	return &JobController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *JobController) Create(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)
	// Parse request body
	request := new(model.JobRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}
	parsedUUID, err := uuid.Parse(auth.ID)
	if err != nil {
		c.Log.Warnf("Invalid UUID format for user ID: %v", err)
		return fiber.ErrUnauthorized
	}
	request.RecruiterID = &parsedUUID
	// Create job
	jobResponse, err := c.Usecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to create post")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.JobResponse]{Data: jobResponse})
}

func (c *JobController) Get(ctx *fiber.Ctx) error {
	request := &model.GetJobRequest{
		ID: ctx.Params("id"),
	}
	response, err := c.Usecase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get post")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.JobResponse]{Data: response})
}

func (c *JobController) List(ctx *fiber.Ctx) error {
	request := &model.SearchJobRequest{
		Title:       ctx.Query("title"),
		Description: ctx.Query("description"),
		Location:    ctx.Query("location"),
		Salary:      ctx.QueryInt("salary"),
		Page:        ctx.QueryInt("page", 1),
		Size:        ctx.QueryInt("size", 10),
	}
	responses, total, err := c.Usecase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search post")
		return err
	}
	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}
	return ctx.JSON(model.WebResponse[[]model.JobResponse]{
		Data:   responses,
		Paging: paging,
	})

}
