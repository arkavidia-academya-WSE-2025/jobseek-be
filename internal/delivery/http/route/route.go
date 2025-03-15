package route

import (
	"fp-academya-be/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

// Add to struct
type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	PostController    *http.PostController
	ProfileController *http.ProfileController
	AuthMiddleware    fiber.Handler
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

// Add to SetupAuthRoute method
func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	//users
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)
	//posts
	c.App.Post("/api/posts", c.PostController.Create)

	// Add profile routes
	c.App.Get("/api/profile/jobseeker", c.ProfileController.GetJobseekerProfile)
	c.App.Put("/api/profile/jobseeker", c.ProfileController.UpdateJobseekerProfile)
	c.App.Get("/api/profile/company", c.ProfileController.GetCompanyProfile)
	c.App.Put("/api/profile/company", c.ProfileController.UpdateCompanyProfile)
}
