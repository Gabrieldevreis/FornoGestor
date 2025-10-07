package routes

import (
	"github.com/Gabrieldevreis/FornoGestor/internal/config"
	"github.com/Gabrieldevreis/FornoGestor/internal/middleware"
	"github.com/Gabrieldevreis/FornoGestor/internal/models"
	"github.com/Gabrieldevreis/FornoGestor/internal/repository"
	"github.com/Gabrieldevreis/FornoGestor/internal/service"
	"github.com/Gabrieldevreis/FornoGestor/internal/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	router.Use(middleware.CORSMiddleware())

	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, cfg)
	userService := service.NewUserService(userRepo)

	// Controllers
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)

	// API routes
	api := router.Group("/api/v1")

	// Public routes
	auth := api.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// Auth
		protected.POST("/auth/logout", authController.Logout)
		protected.GET("/auth/me", authController.Me)

		// Users (Admin only)
		users := protected.Group("/users")
		users.Use(middleware.RoleMiddleware(models.RoleAdmin))
		{
			users.GET("", userController.List)
			users.GET("/:id", userController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

	}
}
