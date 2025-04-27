package service

import "microservices-e-commerce/cmd/user/repository"

type UserService struct {
	UserRepo repository.UserRepositry
}

func NewUserService(userRepo repository.UserRepositry) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}
