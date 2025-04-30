package usecase

import (
	"microservices-e-commerce/cmd/user/service"
	"microservices-e-commerce/infrastructure/log"
	"microservices-e-commerce/models"
	"microservices-e-commerce/utils"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	UserService service.UserService
}

func NewUserUsecase(userService service.UserService) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
	}
}

func (uc *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	user, err := uc.UserService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) RegisterUser(user *models.User) error {
	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("utils.HashPassword() got error %v", err)
		return err
	}

	// insert db
	user.Password = hashedPassword
	_, err = uc.UserService.CreateNewUser(user)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
			"name":  user.Name,
		}).Errorf("uc.UserService.CreateNewUser(user) got error %v", err)
		return err
	}

	return nil
}
