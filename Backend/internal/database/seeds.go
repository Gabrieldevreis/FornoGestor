package database

import (
	"log"

	"gorm.io/gorm"

	"github.com/Gabrieldevreis/FornoGestor/internal/models"
	"github.com/Gabrieldevreis/FornoGestor/internal/utils"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	// Seed admin user
	hashedPassword, _ := utils.HashPassword("admin123")
	admin := models.User{
		Name:     "Administrador",
		Email:    "admin@fornogestor.com",
		Password: hashedPassword,
		Role:     models.RoleAdmin,
		Active:   true,
	}
	db.Create(&admin)

	log.Println("Database seeded successfully")
}
