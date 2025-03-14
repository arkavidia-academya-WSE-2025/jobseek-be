package http

import (
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProfileController struct {
	Log         *logrus.Logger
	ProfileUseCase *usecase.ProfileUseCase
}

func NewProfileController(profileUseCase *usecase.ProfileUseCase, logger *logrus.Logger) *ProfileController {
	return &ProfileController{
		Log:         logger,
		ProfileUseCase: profileUseCase,
	}
}

// GetJobseekerProfile gets the jobseeker profile of the authenticated user
func (c *ProfileController) GetJobseekerProfile(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)
	
	// Get profile
	profileResponse, err := c.ProfileUseCase.GetJobseekerProfile(ctx.UserContext(), auth.ID)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to get jobseeker profile")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.JobseekerProfileResponse]{Data: profileResponse})
}

// UpdateJobseekerProfile updates the jobseeker profile of the authenticated user
func (c *ProfileController) UpdateJobseekerProfile(ctx *fiber.Ctx) error {
	// Parse request body
	request := new(model.UpdateJobseekerProfileRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Get authenticated user
	auth := middleware.GetUser(ctx)
	
	// Update profile
	profileResponse, err := c.ProfileUseCase.UpdateJobseekerProfile(ctx.UserContext(), auth.ID, request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to update jobseeker profile")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.JobseekerProfileResponse]{Data: profileResponse})
}

// GetCompanyProfile gets the company profile of the authenticated user
func (c *ProfileController) GetCompanyProfile(ctx *fiber.Ctx) error {
	// Get authenticated user
	auth := middleware.GetUser(ctx)
	
	// Get profile
	profileResponse, err := c.ProfileUseCase.GetCompanyProfile(ctx.UserContext(), auth.ID)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to get company profile")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.CompanyProfileResponse]{Data: profileResponse})
}

// UpdateCompanyProfile updates the company profile of the authenticated user
func (c *ProfileController) UpdateCompanyProfile(ctx *fiber.Ctx) error {
	// Parse request body
	request := new(model.UpdateCompanyProfileRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	// Get authenticated user
	auth := middleware.GetUser(ctx)
	
	// Update profile
	profileResponse, err := c.ProfileUseCase.UpdateCompanyProfile(ctx.UserContext(), auth.ID, request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to update company profile")
		return err
	}

	// Return JSON response
	return ctx.JSON(model.WebResponse[*model.CompanyProfileResponse]{Data: profileResponse})
}