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
	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	postUseCasse := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)
	//setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	postController := http.NewPostController(postUseCasse, userUseCase, config.Log)
	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		PostController: postController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
