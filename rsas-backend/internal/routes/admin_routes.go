package routes

import (
	"rsas-backend/internal/controllers"
	"rsas-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	// Group rute yang memerlukan login (JWT)
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		// Modul Berita
		news := admin.Group("/news")
		{
			// Menggunakan string kosong "" agar match dengan /api/admin/news
			news.GET("", controllers.GetAllNews)
			news.GET("/:id", controllers.GetNewsByID)
			news.POST("", controllers.CreateNews)
			
			// Editor & Admin bisa edit dan publish
			news.PUT("/:id", middleware.RoleMiddleware("admin", "editor"), controllers.UpdateNews)
			news.PATCH("/:id/publish", middleware.RoleMiddleware("admin", "editor"), controllers.TogglePublishNews)
			
			// Hanya Admin yang bisa hapus
			news.DELETE("/:id", middleware.RoleMiddleware("admin"), controllers.DeleteNews)
		}
	}
}
