package main

import (
	"fmt"
	"github.com/AbdulmalikGiwa/backend-api/internal/api"
	"github.com/AbdulmalikGiwa/backend-api/internal/api/handlers"
	"github.com/AbdulmalikGiwa/backend-api/internal/config"
	"github.com/AbdulmalikGiwa/backend-api/internal/domain"
	"github.com/AbdulmalikGiwa/backend-api/internal/repository"
	"github.com/AbdulmalikGiwa/backend-api/internal/service"
	"github.com/AbdulmalikGiwa/backend-api/pkg/jwt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database with retry mechanism
	var db *gorm.DB
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(cfg.PostgresConnectionString()), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		} else {
			log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
		}
	}

	log.Println("Successfully connected to database")

	// Auto migrate models
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize services and handlers
	jwtService := jwt.NewJWTService(cfg.JWT.Secret, cfg.JWT.Issuer)
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository, jwtService)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup router
	router := api.SetupRouter(authHandler, jwtService)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server running on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
