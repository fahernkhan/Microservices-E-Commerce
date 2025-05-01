package repository

import (
	"context"
	"errors"
	"microservices-e-commerce/models"

	"gorm.io/gorm"
)

func (r *UserRepositry) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.Database.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Email tidak ditemukan, return nil tanpa error //&user
			return nil, nil
		}
		// Error lain (koneksi database dll)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositry) FindByUserID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	err := r.Database.WithContext(ctx).Where("id = ?", userID).Last(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepositry) InsertNewUser(ctx context.Context, user *models.User) (int64, error) {
	err := r.Database.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
