package router

import (
	"pawprints-server/internal/handler"
	"pawprints-server/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	r.Use(middleware.CORS())

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	catHandler := handler.NewCategoryHandler(db)
	r.GET("/api/categories", catHandler.List)
	r.GET("/api/categories/:id/breeds", catHandler.Breeds)
}
