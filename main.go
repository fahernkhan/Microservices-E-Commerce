package main

import (
	"microservices-e-commerce/cmd/user/handler"
	"microservices-e-commerce/cmd/user/resource"
	"microservices-e-commerce/config"
	"microservices-e-commerce/infrastructure/log"
	"microservices-e-commerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	resource.InitRedis(&cfg)
	resource.InitDB(&cfg)

	log.SetupLogger()

	// userRepositry := repository.NewUserRepository(db, redis)
	userHandler := handler.NewUserHandler()

	port := cfg.App.Port // baca config yang kita load diawal
	router := gin.Default()
	routes.SetupRoutes(router, *userHandler)
	router.Run(":" + port)

	log.Logger.Printf("server running on port: %s", port)
}
