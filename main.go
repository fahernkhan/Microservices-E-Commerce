package main

import (
	"microservices-e-commerce/cmd/user/handler"
	"microservices-e-commerce/cmd/user/repository"
	"microservices-e-commerce/cmd/user/resource"
	"microservices-e-commerce/cmd/user/service"
	"microservices-e-commerce/cmd/user/usecase"
	"microservices-e-commerce/config"
	"microservices-e-commerce/infrastructure/log"
	"microservices-e-commerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)

	log.SetupLogger()

	userRepository := repository.NewUserRepository(db, redis)
	userService := service.NewUserService(*userRepository)
	userUsecase := usecase.NewUserUsecase(*userService, cfg.Secret.JWTSecret)
	userHandler := handler.NewUserHandler(*userUsecase)

	port := cfg.App.Port // baca config yang kita load diawal
	router := gin.Default()
	routes.SetupRoutes(router, *userHandler, cfg.Secret.JWTSecret)
	router.Run(":" + port)

	log.Logger.Printf("server running on port: %s", port)
}
