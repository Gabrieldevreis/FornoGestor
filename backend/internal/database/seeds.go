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

	// Seed categories
	categories := []models.ProductCategory{
		{Name: "Pizzas Tradicionais", Description: "Pizzas clássicas"},
		{Name: "Pizzas Especiais", Description: "Pizzas gourmet"},
		{Name: "Bebidas", Description: "Refrigerantes e sucos"},
		{Name: "Lanches", Description: "Sanduíches e porções"},
	}
	db.Create(&categories)

	// Seed financial categories
	finCategories := []models.FinancialCategory{
		{Name: "Vendas", Type: models.TypeRevenue},
		{Name: "Fornecedores", Type: models.TypeExpense},
		{Name: "Salários", Type: models.TypeExpense},
		{Name: "Aluguel", Type: models.TypeExpense},
		{Name: "Utilidades", Type: models.TypeExpense},
	}
	db.Create(&finCategories)

	// Seed tables
	tables := []models.Table{
		{Number: 1, Capacity: 4, Status: models.StatusAvailable},
		{Number: 2, Capacity: 4, Status: models.StatusAvailable},
		{Number: 3, Capacity: 6, Status: models.StatusAvailable},
		{Number: 4, Capacity: 2, Status: models.StatusAvailable},
		{Number: 5, Capacity: 8, Status: models.StatusAvailable},
	}
	db.Create(&tables)

	log.Println("Database seeded successfully")
}
