package controllers

import (
	"net/http"
	"rsas-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// GetBedsAvailability mengambil data dari SIMRS dan mengirimkannya ke frontend
func GetBedsAvailability(c *gin.Context) {
	// Memanggil service SIMRS
	data, err := services.FetchBedAvailability()
	
	if err != nil {
		// Jika API SIMRS mati, kita berikan pesan error yang rapi
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"message": "Gagal mengambil data dari SIMRS",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": data,
	})
}
