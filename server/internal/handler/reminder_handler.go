package handler

import (
	"net/http"
	"strconv"
	"time"
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
	"pawprints-server/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReminderHandler struct {
	Service *service.ReminderService
}

func NewReminderHandler(db *gorm.DB) *ReminderHandler {
	return &ReminderHandler{
		Service: service.NewReminderService(repository.NewReminderRepository(db)),
	}
}

func (h *ReminderHandler) List(c *gin.Context) {
	completed := c.Query("completed") == "true"
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	reminders, total, err := h.Service.List(completed, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reminders, "total": total})
}

func (h *ReminderHandler) GetActive(c *gin.Context) {
	reminders, err := h.Service.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reminders)
}

func (h *ReminderHandler) Create(c *gin.Context) {
	var reminder model.Reminder
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reminder)
}

func (h *ReminderHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	reminder, err := h.Service.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}
	if err := c.ShouldBindJSON(reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reminder)
}

func (h *ReminderHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *ReminderHandler) Complete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Complete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "completed"})
}

func (h *ReminderHandler) Snooze(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Minutes int `json:"minutes"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		input.Minutes = 30 // 默认延后30分钟
	}
	if err := h.Service.Snooze(uint(id), time.Duration(input.Minutes)*time.Minute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "snoozed"})
}