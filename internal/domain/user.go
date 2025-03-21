package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"` // "-" ensures it won't be included in JSON responses
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Duration  `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SignupRequest struct {
	Name     string `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AuthResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
