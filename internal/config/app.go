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
	//setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	postRepository := repository.NewPostRepository(config.Log)
	jobseekerProfileRepository := repository.NewJobseekerProfileRepository(config.Log)
	companyProfileRepository := repository.NewCompanyProfileRepository(config.Log)
	messageRepository := repository.NewMessageRepository(config.Log)

	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	postUseCasse := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)
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
	postController := http.NewPostController(postUseCasse, config.Log)
	profileController := http.NewProfileController(profileUseCase, config.Log)
	messageController := http.NewMessageController(messageUseCase, config.Log)
	
	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		PostController:    postController,
		ProfileController: profileController,
		MessageController: messageController,
		AuthMiddleware:    authMiddleware,
		UserUseCase:       userUseCase,
	}
	routeConfig.Setup()
}
