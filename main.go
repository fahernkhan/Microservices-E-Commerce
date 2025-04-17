package main

import (
	"microservices-e-commerce/config"
	"microservices-e-commerce/infrastructure/log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	log.SetupLogger()

	port := cfg.App.Port // baca config yang kita load diawal
	router := gin.Default()
	router.Run(":" + port)

	log.Logger.Printf("server running on port: %s", port)
}
