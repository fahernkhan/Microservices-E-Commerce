package service

import (
	"microservices-e-commerce/cmd/user/repository"
	"microservices-e-commerce/models"
)

type UserService struct {
	UserRepo repository.UserRepositry
}

func NewUserService(userRepo repository.UserRepositry) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (svc *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := svc.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UserService) CreateNewUser(user *models.User) (int64, error) {
	userID, err := svc.UserRepo.InsertNewUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
