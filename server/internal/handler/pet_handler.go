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

type PetHandler struct {
	Service *service.PetService
}

func NewPetHandler(db *gorm.DB) *PetHandler {
	return &PetHandler{
		Service: service.NewPetService(repository.NewPetRepository(db)),
	}
}

func (h *PetHandler) List(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	pets, total, err := h.Service.List(status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pets, "total": total})
}

func (h *PetHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pet, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) Create(c *gin.Context) {
	var pet model.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&pet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pet)
}

func (h *PetHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pet, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	if err := c.ShouldBindJSON(pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(pet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *PetHandler) SetMemorial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.SetMemorial(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "memorial mode set"})
}

func (h *PetHandler) GetPhotos(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	photos, err := h.Service.GetPhotos(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, photos)
}
