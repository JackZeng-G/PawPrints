package handler

import (
	"net/http"
	"strconv"
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
	"pawprints-server/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	Service *service.HealthService
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		Service: service.NewHealthService(repository.NewHealthRepository(db)),
	}
}

func (h *HealthHandler) List(c *gin.Context) {
	petID, _ := strconv.Atoi(c.DefaultQuery("pet_id", "0"))
	recordType := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	records, total, err := h.Service.List(uint(petID), recordType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": records, "total": total})
}

func (h *HealthHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	record, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (h *HealthHandler) Create(c *gin.Context) {
	var record model.HealthRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, record)
}

func (h *HealthHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	record, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (h *HealthHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *HealthHandler) WeightChart(c *gin.Context) {
	petID, _ := strconv.Atoi(c.Param("id"))
	records, err := h.Service.GetWeightChart(uint(petID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func (h *HealthHandler) Upcoming(c *gin.Context) {
	petID, _ := strconv.Atoi(c.Param("id"))
	records, err := h.Service.GetUpcoming(uint(petID), 30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}