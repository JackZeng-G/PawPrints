package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pawprints-server/internal/model"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{DB: db}
}

func (h *CategoryHandler) List(c *gin.Context) {
	var categories []model.PetCategory
	h.DB.Order("sort_order").Find(&categories)
	c.JSON(200, categories)
}

func (h *CategoryHandler) Breeds(c *gin.Context) {
	categoryID := c.Param("id")
	var breeds []model.PetBreed
	h.DB.Where("category_id = ?", categoryID).Find(&breeds)
	c.JSON(200, breeds)
}
