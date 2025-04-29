package usecase

import (
	"microservices-e-commerce/cmd/user/service"
	"microservices-e-commerce/models"
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
	_, err := uc.UserService.CreateNewUser(user)
	if err != nil {
		return err
	}

	return nil
}
