package database

import (
	"fmt"
	"log"

	"github.com/Gabrieldevreis/FornoGestor/internal/config"
	"github.com/Gabrieldevreis/FornoGestor/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Client{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Table{},
		&models.Order{},
		&models.OrderItem{},
		&models.FinancialCategory{},
		&models.FinancialEntry{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}
