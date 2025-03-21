package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server struct {
		Port int
		Host string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	JWT struct {
		Secret string
		Issuer string
	}
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{}

	// Server
	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	config.Server.Port = port
	config.Server.Host = getEnv("SERVER_HOST", "")

	// Database
	config.Database.Host = getEnv("DB_HOST", "localhost")
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	config.Database.Port = dbPort
	config.Database.User = getEnv("DB_USER", "postgres")
	config.Database.Password = getEnv("DB_PASSWORD", "postgres")
	config.Database.DBName = getEnv("DB_NAME", "auth_api")
	config.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT
	config.JWT.Secret = getEnv("JWT_SECRET", "your-secret-key")
	config.JWT.Issuer = getEnv("JWT_ISSUER", "backend-api")

	return config, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// PostgresConnectionString returns the connection string for PostgreSQL
func (c *Config) PostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}
