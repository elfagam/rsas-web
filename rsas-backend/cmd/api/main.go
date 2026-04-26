package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rsas-backend/internal/config"
	"rsas-backend/internal/routes"
	"rsas-backend/internal/models"
	"rsas-backend/internal/middleware"
	"rsas-backend/pkg/utils"
	"log"
)

func main() {
	// Initialize Database
	config.ConnectDatabase()

	// Auto-Seed Admin User
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", "admin_rsas").Count(&count)
	if count == 0 {
		hashedPassword, _ := utils.HashPassword("AdminRSAS2024!")
		admin := models.User{
			Username: "admin_rsas",
			Password: hashedPassword,
			FullName: "Administrator RSAS",
			Role:     "admin",
		}
		config.DB.Create(&admin)
		log.Println("Default Admin User created: admin_rsas / AdminRSAS2024!")
	}

	r := gin.Default()

	// Global Middlewares
	r.Use(middleware.CorsMiddleware())

	// Register Routes
	routes.AuthRoutes(r)
	routes.PublicRoutes(r)
	routes.AdminRoutes(r)

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "RSAS Backend API is running",
		})
	})

	r.Run(":8000") // Menjalankan server di port 8000
}
