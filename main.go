package main

import (
	"order-service/cmd/user/handler"
	"order-service/cmd/user/repository"
	"order-service/cmd/user/resource"
	"order-service/cmd/user/service"
	"order-service/cmd/user/usecase"
	"order-service/config"
	"order-service/infrastructure/log"
	"order-service/routes"

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
