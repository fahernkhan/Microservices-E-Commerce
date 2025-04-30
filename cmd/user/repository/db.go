package repository

import (
	"errors"
	"microservices-e-commerce/models"

	"gorm.io/gorm"
)

func (r *UserRepositry) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Email tidak ditemukan, return nil tanpa error
			return nil, nil
		}
		// Error lain (koneksi database dll)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositry) InsertNewUser(user *models.User) (int64, error) {
	err := r.Database.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
