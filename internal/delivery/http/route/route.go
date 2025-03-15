package route

import (
	"fp-academya-be/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	PostController *http.PostController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	//users
	c.App.Post("/api/users/register", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)
	//posts
	c.App.Get("/api/posts", c.PostController.List)
	c.App.Get("/api/posts/:id", c.PostController.Get)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	//users
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)
	//posts
	c.App.Post("/api/posts/new", c.PostController.Create)
}
