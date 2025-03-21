package api

import (
	"github.com/AbdulmalikGiwa/backend-api/internal/api/handlers"
	"github.com/AbdulmalikGiwa/backend-api/internal/api/middleware"
	"github.com/AbdulmalikGiwa/backend-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the API routes
func SetupRouter(
	authHandler *handlers.AuthHandler,
	jwtService *jwt.JWTService,
) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/api/auth/signup", authHandler.Signup)
	router.POST("/api/auth/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		protected.GET("/profile", authHandler.GetProfile)
	}

	return router
}
