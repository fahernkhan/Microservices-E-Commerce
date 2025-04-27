package usecase

import "microservices-e-commerce/cmd/user/service"

type UserUsecase struct {
	UserService service.UserService
}

func NewUserUsecase(userService service.UserService) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
	}
}
