package services

import (
	"api-alemao/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) BuscarUsuarioPorEmail(email string) (models.User, error) {
	var existingUser models.User

	err := s.db.Where("email = ?", email).First(&existingUser).Error

	return existingUser, err
}
