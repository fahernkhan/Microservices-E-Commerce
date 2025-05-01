package usecase

import (
	"context"
	"errors"
	"microservices-e-commerce/cmd/user/service"
	"microservices-e-commerce/infrastructure/log"
	"microservices-e-commerce/models"
	"microservices-e-commerce/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	UserService service.UserService
	JWTSecret   string
}

func NewUserUsecase(userService service.UserService, jwtsecret string) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
		JWTSecret:   jwtsecret,
	}
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) GetUserByUserID(ctx context.Context, userID int64) (*models.User, error) {
	user, err := uc.UserService.GetUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, user *models.User) error {
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
	_, err = uc.UserService.CreateNewUser(ctx, user)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
			"name":  user.Name,
		}).Errorf("uc.UserService.CreateNewUser(user) got error %v", err)
		return err
	}

	return nil
}

func (uc *UserUsecase) Login(ctx context.Context, param *models.LoginParameter) (string, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, param.Email)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("uc.UserService.GetUserByEmail got error: %v", err)
		return "", errors.New("Email not found")
	}

	if user == nil || user.ID == 0 {
		return "", errors.New("Email not found")
	}

	isMatch, err := utils.CheckPasswordHash(param.Password, user.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("utils.CheckPasswordHash got error: %v", err)
		return "", err
	}
	if !isMatch {
		return "", errors.New("Email atau password salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("token.SignedString got error: %v", err)
		return "", err
	}

	return tokenString, nil
}
