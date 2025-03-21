package route

import (
	"fp-academya-be/internal/delivery/http"
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// RouteConfig holds controllers and middleware for route setup
type RouteConfig struct {
	App                   *fiber.App
	UserController        *http.UserController
	PostController        *http.PostController
	JobController         *http.JobController
	ApplicationController *http.ApplicationController
	ProfileController     *http.ProfileController
	MessageController     *http.MessageController
	AuthMiddleware        fiber.Handler
	UserUseCase           *usecase.UserUseCase
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	//users
	c.App.Post("/api/users/register", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)
	c.App.Get("/api/users/:id", c.UserController.Get)
	//posts
	c.App.Get("/api/posts", c.PostController.List)
	c.App.Get("/api/posts/:id", c.PostController.Get)

	//jobs
	c.App.Get("/api/jobs", c.JobController.List)
	c.App.Get("/api/jobs/:id", c.JobController.Get)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	//users
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)
	//posts
	c.App.Post("/api/posts", c.PostController.Create)

	// Profile routes with role middleware
	//job seeker
	c.App.Get("/api/profile/jobseeker", middleware.JobSeekerOnly(c.UserUseCase), c.ProfileController.GetJobseekerProfile)
	c.App.Put("/api/profile/jobseeker", middleware.JobSeekerOnly(c.UserUseCase), c.ProfileController.UpdateJobseekerProfile)
	//applications
	c.App.Post("/api/applications", middleware.JobSeekerOnly(c.UserUseCase), c.ApplicationController.Create)
	c.App.Put("/api/applications/:id", middleware.JobSeekerOnly(c.UserUseCase), c.ApplicationController.Update)
	c.App.Delete("/api/applications/:id", middleware.JobSeekerOnly(c.UserUseCase), c.ApplicationController.Delete)
	//company
	c.App.Get("/api/profile/company", middleware.CompanyOnly(c.UserUseCase), c.ProfileController.GetCompanyProfile)
	c.App.Put("/api/profile/company", middleware.CompanyOnly(c.UserUseCase), c.ProfileController.UpdateCompanyProfile)

	//jobs
	c.App.Post("/api/jobs", middleware.CompanyOnly(c.UserUseCase), c.JobController.Create)
	c.App.Delete("/api/jobs/:id", middleware.CompanyOnly(c.UserUseCase), c.JobController.Delete)
	c.App.Put("/api/jobs/:id", middleware.CompanyOnly(c.UserUseCase), c.JobController.Update)

	//applications
	c.App.Get("/api/applications", middleware.CompanyOnly(c.UserUseCase), c.ApplicationController.List)
	c.App.Get("/api/applications/:id", middleware.CompanyOnly(c.UserUseCase), c.ApplicationController.Get)

	// Messages routes - available to all authenticated users
	c.App.Post("/api/messages", c.MessageController.SendMessage)
	c.App.Post("/api/messages/conversation", c.MessageController.GetConversation)
	c.App.Post("/api/messages/mark-read", c.MessageController.MarkAsRead)
	c.App.Get("/api/messages/unread-count", c.MessageController.GetUnreadCount)
}
