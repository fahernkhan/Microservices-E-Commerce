package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepositry struct {
	Database *gorm.DB
	Redis    *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) *UserRepositry {
	return &UserRepositry{
		Database: db,
		Redis:    redis,
	}
}
