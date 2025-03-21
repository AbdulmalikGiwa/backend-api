package service

import (
	"Ahm/internal/domain"
	"Ahm/internal/repository"
	"Ahm/pkg/hasher"
	"Ahm/pkg/jwt"
	"errors"
)

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtService *jwt.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *jwt.JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

// Signup creates a new user
func (s *AuthService) Signup(req domain.SignupRequest) (domain.AuthResponse, error) {
	// Check if email already exists
	if s.userRepo.EmailExists(req.Email) {
		return domain.AuthResponse{}, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := hasher.HashPassword(req.Password)
	if err != nil {
		return domain.AuthResponse{}, errors.New("failed to hash password")
	}

	// Create user
	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return domain.AuthResponse{}, errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateToken(createdUser)
	if err != nil {
		return domain.AuthResponse{}, errors.New("failed to generate token")
	}

	// Return user and token
	return domain.AuthResponse{
		User:  createdUser,
		Token: token,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(req domain.LoginRequest) (domain.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return domain.AuthResponse{}, errors.New("invalid email or password")
	}

	// Check if password is correct
	if !hasher.CheckPasswordHash(req.Password, user.Password) {
		return domain.AuthResponse{}, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return domain.AuthResponse{}, errors.New("failed to generate token")
	}

	// Return user and token
	return domain.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(id uint) (domain.User, error) {
	return s.userRepo.FindByID(id)
}
