package main

import (
	"log"
	"rsas-backend/internal/config"
	"rsas-backend/internal/models"
	"rsas-backend/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	godotenv.Load("../../.env")

	// Connect to DB
	config.ConnectDatabase()

	// Cek apakah admin sudah ada
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", "admin_rsas").Count(&count)
	if count > 0 {
		log.Println("Admin user already exists")
		return
	}

	// Buat Admin
	hashedPassword, _ := utils.HashPassword("AdminRSAS2024!")
	admin := models.User{
		Username: "admin_rsas",
		Password: hashedPassword,
		FullName: "Administrator RSAS",
		Role:     "admin",
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		log.Fatal("Failed to create admin:", err)
	}

	log.Println("Admin user created successfully!")
}
