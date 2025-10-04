package main

import (
	"log"
	"fornogestor/internal/config"
	"fornogestor/internal/database"
	"fornogestor/internal/routes"
	
	"github.com/gin-gonic/gin"
	_ "fornogestor/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FornoGestor API
// @version 1.0
// @description API para gestão de pizzarias
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@fornogestor.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	cfg := config.LoadConfig()
	
	db := database.Connect(cfg)
	database.Migrate(db)
	database.Seed(db)
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	routes.SetupRoutes(router, db, cfg)
	
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}