package service

import (
	"context"
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

func (svc *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := svc.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UserService) GetUserByUserID(ctx context.Context, userID int64) (*models.User, error) {
	user, err := svc.UserRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserService) CreateNewUser(ctx context.Context, user *models.User) (int64, error) {
	userID, err := svc.UserRepo.InsertNewUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
