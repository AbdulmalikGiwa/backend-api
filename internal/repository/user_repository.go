// internal/repository/user_repository.go
package repository

import (
	"Ahm/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user in the database
func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(id uint) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

// EmailExists checks if an email already exists
func (r *UserRepository) EmailExists(email string) bool {
	var count int64
	r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
