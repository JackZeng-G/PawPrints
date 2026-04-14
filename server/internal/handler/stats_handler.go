package handler

import (
	"net/http"
	"os"
	"strconv"
	"pawprints-server/internal/config"
	"pawprints-server/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StatsHandler struct {
	Service *service.StatsService
	Config  *config.Config
}

func NewStatsHandler(db *gorm.DB, cfg *config.Config) *StatsHandler {
	return &StatsHandler{Service: service.NewStatsService(db), Config: cfg}
}

func (h *StatsHandler) Overview(c *gin.Context) {
	stats, err := h.Service.Overview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *StatsHandler) Timeline(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))
	items, err := h.Service.Timeline(days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *StatsHandler) Export(c *gin.Context) {
	data, err := h.Service.ExportJSON()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *StatsHandler) Backup(c *gin.Context) {
	dbPath := h.Config.Database.Path
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Database file not found"})
		return
	}
	c.FileAttachment(dbPath, "pawprints-backup.db")
}
