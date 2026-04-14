package router

import (
	"pawprints-server/internal/config"
	"pawprints-server/internal/handler"
	"pawprints-server/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	r.Use(middleware.CORS())

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 分类
	catHandler := handler.NewCategoryHandler(db)
	r.GET("/api/categories", catHandler.List)
	r.GET("/api/categories/:id/breeds", catHandler.Breeds)

	// 宠物
	petHandler := handler.NewPetHandler(db)
	r.GET("/api/pets", petHandler.List)
	r.POST("/api/pets", petHandler.Create)
	r.GET("/api/pets/:id", petHandler.Get)
	r.PUT("/api/pets/:id", petHandler.Update)
	r.DELETE("/api/pets/:id", petHandler.Delete)
	r.PUT("/api/pets/:id/memorial", petHandler.SetMemorial)
	r.GET("/api/pets/:id/photos", petHandler.GetPhotos)

	// 日记
	diaryHandler := handler.NewDiaryHandler(db)
	r.GET("/api/diaries", diaryHandler.List)
	r.POST("/api/diaries", diaryHandler.Create)
	r.GET("/api/diaries/:id", diaryHandler.Get)
	r.PUT("/api/diaries/:id", diaryHandler.Update)
	r.DELETE("/api/diaries/:id", diaryHandler.Delete)
	r.POST("/api/diaries/:id/photos", diaryHandler.AddPhoto)
	r.DELETE("/api/diaries/:id/photos/:photoId", diaryHandler.DeletePhoto)

	// 健康记录
	healthHandler := handler.NewHealthHandler(db)
	r.GET("/api/health", healthHandler.List)
	r.POST("/api/health", healthHandler.Create)
	r.GET("/api/health/:id", healthHandler.Get)
	r.PUT("/api/health/:id", healthHandler.Update)
	r.DELETE("/api/health/:id", healthHandler.Delete)
	r.GET("/api/health/:id/weight-chart", healthHandler.WeightChart)
	r.GET("/api/health/:id/upcoming", healthHandler.Upcoming)

	// 提醒
	reminderHandler := handler.NewReminderHandler(db)
	r.GET("/api/reminders", reminderHandler.List)
	r.GET("/api/reminders/active", reminderHandler.GetActive)
	r.POST("/api/reminders", reminderHandler.Create)
	r.PUT("/api/reminders/:id", reminderHandler.Update)
	r.DELETE("/api/reminders/:id", reminderHandler.Delete)
	r.PUT("/api/reminders/:id/complete", reminderHandler.Complete)
	r.PUT("/api/reminders/:id/snooze", reminderHandler.Snooze)

	// 上传
	cfg := config.Load()
	uploadHandler := handler.NewUploadHandler(cfg.Server.UploadDir, cfg.Thumbnail.MaxWidth, cfg.Thumbnail.MaxHeight)
	r.POST("/api/upload", uploadHandler.Upload)
	r.GET("/api/upload/*filepath", func(c *gin.Context) {
		c.File(cfg.Server.UploadDir + c.Param("filepath"))
	})
}
