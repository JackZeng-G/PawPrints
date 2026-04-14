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

type DiaryHandler struct {
	Service *service.DiaryService
}

func NewDiaryHandler(db *gorm.DB) *DiaryHandler {
	return &DiaryHandler{
		Service: service.NewDiaryService(repository.NewDiaryRepository(db)),
	}
}

func (h *DiaryHandler) List(c *gin.Context) {
	petID, _ := strconv.Atoi(c.DefaultQuery("pet_id", "0"))
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	entries, total, err := h.Service.List(uint(petID), keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": entries, "total": total})
}

func (h *DiaryHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	entry, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (h *DiaryHandler) Create(c *gin.Context) {
	var input service.CreateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entry, err := h.Service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, entry)
}

func (h *DiaryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input service.CreateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entry, err := h.Service.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (h *DiaryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *DiaryHandler) AddPhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var photo model.DiaryPhoto
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	photo.DiaryEntryID = uint(id)
	if err := h.Service.AddPhoto(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, photo)
}

func (h *DiaryHandler) DeletePhoto(c *gin.Context) {
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	if err := h.Service.DeletePhoto(uint(photoID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
