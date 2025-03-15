package middleware

import (
	"fp-academya-be/internal/model"
	"fp-academya-be/internal/usecase"
	"fp-academya-be/internal/helper"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUsecase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userUsecase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUsecase.Verify(ctx.UserContext(), request)

		if err != nil {
			userUsecase.Log.Warnf("Failed to verify user: %v", err)
			return fiber.ErrUnauthorized
		}
		userUsecase.Log.Debugf("User: %v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return helper.GetUser(ctx)
}

// RoleMiddleware creates a middleware that checks if the user has the required role
func RoleMiddleware(userUsecase *usecase.UserUseCase, requiredRole string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := GetUser(ctx)
		if auth == nil {
			return fiber.ErrUnauthorized
		}

		// Use role directly from auth object since it's already included from Verify
		if requiredRole == "recruiter" && auth.Role != "recruiter" {
			return fiber.NewError(fiber.StatusForbidden, "Access denied: company access only")
		} else if requiredRole == "job_seeker" && auth.Role != "job_seeker" {
			return fiber.NewError(fiber.StatusForbidden, "Access denied: job seeker access only")
		}

		return ctx.Next()
	}
}

// JobSeekerOnly middleware ensures only job seekers can access the route
func JobSeekerOnly(userUsecase *usecase.UserUseCase) fiber.Handler {
	return RoleMiddleware(userUsecase, "job_seeker")
}

// CompanyOnly middleware ensures only companies can access the route
func CompanyOnly(userUsecase *usecase.UserUseCase) fiber.Handler {
	return RoleMiddleware(userUsecase, "recruiter")
}
