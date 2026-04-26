package routes

import (
	"rsas-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine) {
	public := r.Group("/api/public")
	{
		// Endpoint untuk ketersediaan tempat tidur (Data dari SIMRS)
		public.GET("/beds", controllers.GetBedsAvailability)
		
		// Endpoint Berita untuk Publik
		public.GET("/news", controllers.GetAllNews)
		public.GET("/news/:slug", controllers.GetNewsBySlug)
	}
}
