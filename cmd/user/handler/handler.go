package handler

import (
	"microservices-e-commerce/cmd/user/usecase"
	"microservices-e-commerce/infrastructure/log"
	"microservices-e-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var param models.RegisterParameter
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			// "error_message": err.Error(), gak valid itampilin di user
			"error_message": "invalid input parameter",
		})
		return
	}

	if len(param.Password) < 8 ||
		len(param.ConfirmPassword) < 8 {
		log.Logger.Info("Invalid Input")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password must longer than 8 characters",
		})
		return
	}

	if param.Password != param.ConfirmPassword {
		log.Logger.Info("Invalid Credentials")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password and Confirm Password Not Match",
		})
		return
	}

	user, err := h.UserUsecase.GetUserByEmail(param.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	// masuk ke use case resgister
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Email already exists!",
		})
		return
	}

	err = h.UserUsecase.RegisterUser(&models.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User sucessfulllly registered!",
	})
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
