// Add to struct
type RouteConfig struct {
	App                *fiber.App
	UserController     *http.UserController
	PostController     *http.PostController
	ProfileController  *http.ProfileController
	AuthMiddleware     fiber.Handler
}

// Add to SetupAuthRoute method
func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)
	c.App.Post("/api/posts/new", c.PostController.Create)
	
	// Add profile routes
	c.App.Get("/api/profile/jobseeker", c.ProfileController.GetJobseekerProfile)
	c.App.Put("/api/profile/jobseeker", c.ProfileController.UpdateJobseekerProfile)
	c.App.Get("/api/profile/company", c.ProfileController.GetCompanyProfile)
	c.App.Put("/api/profile/company", c.ProfileController.UpdateCompanyProfile)
}