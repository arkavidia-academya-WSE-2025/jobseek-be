package http

import (
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PostController struct {
	Log         *logrus.Logger
	PostUsecase *usecase.PostUseCase
	UserUseCase *usecase.UserUseCase
}

func NewPostController(postUsecase *usecase.PostUseCase, userUsecase *usecase.UserUseCase, logger *logrus.Logger) *PostController {
	return &PostController{
		Log:         logger,
		PostUsecase: postUsecase,
		UserUseCase: userUsecase,
	}
}

func (c *PostController) Create(ctx *fiber.Ctx) error {
	// Parse request body
	request := new(model.PostRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Get authenticated user
	auth := middleware.GetUser(ctx)
	user := &model.GetUserRequest{
		ID: auth.ID,
	}

	// Fetch current user
	userResponse, err := c.UserUseCase.Current(ctx.UserContext(), user)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	// Create post
	postResponse, err := c.PostUsecase.Create(ctx.UserContext(), request, *userResponse.ID)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to create post")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.PostReponse]{Data: postResponse})
}
