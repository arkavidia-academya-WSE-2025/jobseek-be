package route

import (
	"fp-academya-be/internal/delivery/http"
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// Add to struct
type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	PostController    *http.PostController
	ProfileController *http.ProfileController
	MessageController *http.MessageController
	AuthMiddleware    fiber.Handler
	UserUseCase       *usecase.UserUseCase
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

	// Profile routes with role middleware
	// Job seeker profile routes
	c.App.Get("/api/profile/jobseeker", middleware.JobSeekerOnly(c.UserUseCase), c.ProfileController.GetJobseekerProfile)
	c.App.Put("/api/profile/jobseeker", middleware.JobSeekerOnly(c.UserUseCase), c.ProfileController.UpdateJobseekerProfile)
	
	// Company profile routes
	c.App.Get("/api/profile/company", middleware.CompanyOnly(c.UserUseCase), c.ProfileController.GetCompanyProfile)
	c.App.Put("/api/profile/company", middleware.CompanyOnly(c.UserUseCase), c.ProfileController.UpdateCompanyProfile)

	// Messages routes - available to all authenticated users
	c.App.Post("/api/messages", c.MessageController.SendMessage)
	c.App.Post("/api/messages/conversation", c.MessageController.GetConversation)
	c.App.Post("/api/messages/mark-read", c.MessageController.MarkAsRead)
	c.App.Get("/api/messages/unread-count", c.MessageController.GetUnreadCount)
}
