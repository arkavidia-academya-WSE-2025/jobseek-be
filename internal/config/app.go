package config

import (
	"fp-academya-be/internal/delivery/http"
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
	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	//setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	//setup middleware
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
