package config

import (
	"fp-academya-be/internal/delivery/http"
	"fp-academya-be/internal/delivery/http/middleware"
	"fp-academya-be/internal/delivery/http/route"
	"fp-academya-be/internal/repository"
	"fp-academya-be/internal/usecase"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func BootStrap(config *BootstrapConfig) {
	// Apply CORS middleware
	config.App.Use(middleware.CORSConfig())
	//setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	postRepository := repository.NewPostRepository(config.Log)
	jobRepository := repository.NewJobRepository(config.Log)
	applicationRepository := repository.NewApplicationRepository(config.Log)
	jobseekerProfileRepository := repository.NewJobseekerProfileRepository(config.Log)
	companyProfileRepository := repository.NewCompanyProfileRepository(config.Log)
	messageRepository := repository.NewMessageRepository(config.Log)

	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	postUseCase := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)
	jobUseCase := usecase.NewJobUseCase(config.DB, config.Log, config.Validate, jobRepository)
	applicationUseCase := usecase.NewApplicationUsecase(config.DB, config.Log, config.Validate, applicationRepository)
	profileUseCase := usecase.NewProfileUseCase(
		config.DB,
		config.Log,
		config.Validate,
		jobseekerProfileRepository,
		companyProfileRepository,
		userRepository,
	)
	messageUseCase := usecase.NewMessageUseCase(
		config.DB,
		config.Log,
		config.Validate,
		messageRepository,
		userRepository,
	)

	//setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	postController := http.NewPostController(postUseCase, config.Log)
	jobController := http.NewJobController(jobUseCase, config.Log)
	applicationController := http.NewApplicationController(applicationUseCase, config.Log)
	profileController := http.NewProfileController(profileUseCase, config.Log)
	messageController := http.NewMessageController(messageUseCase, config.Log)

	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                   config.App,
		UserController:        userController,
		PostController:        postController,
		JobController:         jobController,
		ApplicationController: applicationController,
		ProfileController:     profileController,
		MessageController:     messageController,
		AuthMiddleware:        authMiddleware,
		UserUseCase:           userUseCase,
	}
	routeConfig.Setup()
}
