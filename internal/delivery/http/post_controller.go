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

type PostController struct {
	Log     *logrus.Logger
	Usecase *usecase.PostUseCase
}

func NewPostController(usecase *usecase.PostUseCase, logger *logrus.Logger) *PostController {
	return &PostController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *PostController) Create(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)
	// Parse request body
	request := new(model.PostRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}
	parsedUUID, err := uuid.Parse(auth.ID)
	if err != nil {
		c.Log.Warnf("Invalid UUID format for user ID: %v", err)
		return fiber.ErrUnauthorized
	}
	request.UserId = &parsedUUID
	// Create post
	postResponse, err := c.Usecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to create post")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.PostReponse]{Data: postResponse})
}

func (c *PostController) List(ctx *fiber.Ctx) error {
	request := &model.SearchPostRequest{
		Title:   ctx.Query("title"),
		Content: ctx.Query("content"),
		Page:    ctx.QueryInt("page", 1),
		Size:    ctx.QueryInt("size", 10),
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
	return ctx.JSON(model.WebResponse[[]model.PostReponse]{
		Data:   responses,
		Paging: paging,
	})

}
