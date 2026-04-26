package controllers

import (
	"fmt"
	"net/http"
	"rsas-backend/internal/config"
	"rsas-backend/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type NewsInput struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Category  string `json:"category"`
	Thumbnail string `json:"thumbnail"`
}

// CreateNews membuat berita baru
func CreateNews(c *gin.Context) {
	var input NewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil AuthorID dari token (via Middleware)
	userID := c.MustGet("user_id").(float64)

	// Buat Slug dari Title (Sederhana)
	slug := strings.ToLower(strings.ReplaceAll(input.Title, " ", "-"))
	
	// Cek duplikasi slug dan tambahkan suffix jika perlu
	originalSlug := slug
	count := 1
	for {
		var existingNews models.News
		if err := config.DB.Where("slug = ?", slug).First(&existingNews).Error; err != nil {
			// Jika tidak ditemukan (error), berarti slug tersedia
			break
		}
		// Jika ditemukan, tambah suffix
		slug = fmt.Sprintf("%s-%d", originalSlug, count)
		count++
	}

	news := models.News{
		Title:     input.Title,
		Slug:      slug,
		Content:   input.Content,
		Category:  input.Category,
		Thumbnail: input.Thumbnail,
		AuthorID:  uint(userID),
		IsDraft:   true, // Default adalah draft
	}

	if err := config.DB.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News created successfully", "data": news})
}

// GetAllNews mengambil semua berita
func GetAllNews(c *gin.Context) {
	var news []models.News
	query := config.DB.Preload("Author")

	// Jika dipanggil dari rute publik (tidak ada header Auth), hanya tampilkan yang bukan draft
	// Atau kita cek path-nya mengandung "/public"
	if strings.Contains(c.Request.URL.Path, "/public") {
		query = query.Where("is_draft = ?", false)
	}

	query.Order("created_at desc").Find(&news)

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// GetNewsByID mengambil satu berita berdasarkan ID
func GetNewsByID(c *gin.Context) {
	var news models.News
	id := c.Param("id")

	if err := config.DB.Preload("Author").First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// GetNewsBySlug mengambil satu berita berdasarkan slug
func GetNewsBySlug(c *gin.Context) {
	var news models.News
	slug := c.Param("slug")

	if err := config.DB.Preload("Author").Where("slug = ?", slug).First(&news).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// UpdateNews mengubah berita
func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	var input NewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&news).Updates(models.News{
		Title:     input.Title,
		Content:   input.Content,
		Category:  input.Category,
		Thumbnail: input.Thumbnail,
	})

	c.JSON(http.StatusOK, gin.H{"message": "News updated successfully", "data": news})
}

// DeleteNews menghapus berita
func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	config.DB.Delete(&news)

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}

// TogglePublishNews mengubah status publikasi berita
func TogglePublishNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	// Balikkan status is_draft
	news.IsDraft = !news.IsDraft

	if err := config.DB.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update publication status"})
		return
	}

	status := "published"
	if news.IsDraft {
		status = "drafted"
	}

	c.JSON(http.StatusOK, gin.H{"message": "News successfully " + status, "data": news})
}
