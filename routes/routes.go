package routes

import (
	"microservices-e-commerce/cmd/user/handler"
	"microservices-e-commerce/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler, jwtSecret string) {
	// Public API
	router.Use(middleware.RequestLogger())
	router.GET("/ping", userHandler.Ping)
	router.POST("/v1/register", userHandler.Register)
	router.POST("/v1/login", userHandler.Login)

	// Private API
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	private := router.Group("/api")
	private.Use(authMiddleware)
	private.GET("/v1/user_info", userHandler.GetUserInfo)
}
