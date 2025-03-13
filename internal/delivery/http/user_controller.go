package http

import (
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	Usecase *usecase.UserUseCase
}

func NewUserController(usecase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}
	response, err := c.Usecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create user: %v", err)
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.Usecase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user: %v", err)
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.Usecase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.Usecase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
