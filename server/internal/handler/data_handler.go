package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"pawprints-server/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataHandler struct {
	DB *gorm.DB
}

func NewDataHandler(db *gorm.DB) *DataHandler {
	return &DataHandler{DB: db}
}

func (h *DataHandler) Import(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var importData map[string]interface{}
	if err := json.Unmarshal(data, &importData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// 简化实现：只导入宠物
	if pets, ok := importData["pets"].([]interface{}); ok {
		for _, p := range pets {
			petJSON, _ := json.Marshal(p)
			var pet model.Pet
			json.Unmarshal(petJSON, &pet)
			h.DB.Create(&pet)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "imported"})
}