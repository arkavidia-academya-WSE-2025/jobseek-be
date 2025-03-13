package route

import (
	"fp-academya-be/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/register", c.UserController.Register)
	c.App.Post("/api/login", c.UserController.Login)
	c.App.Get("/api/current", c.UserController.Current)
	c.App.Post("/api/logout", c.UserController.Logout)
}
