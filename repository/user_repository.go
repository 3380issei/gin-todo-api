package repository

import (
	"todo-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(user *model.User, username string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByUsername(user *model.User, username string) error {
	if err := ur.db.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
