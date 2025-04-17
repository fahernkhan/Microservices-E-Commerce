package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	// Public API
	router.GET("/ping", userHandler.ping)

	// Private API
}
