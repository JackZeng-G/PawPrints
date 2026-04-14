package handler

import (
	"net/http"
	"path/filepath"
	"pawprints-server/internal/service"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	Service *service.UploadService
}

func NewUploadHandler(uploadDir string, maxWidth, maxHeight int) *UploadHandler {
	return &UploadHandler{
		Service: service.NewUploadService(uploadDir, maxWidth, maxHeight),
	}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	result, err := h.Service.Upload(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回完整 URL 路径
	result.URL = "/api/upload/" + result.URL
	if result.ThumbnailURL != "" {
		result.ThumbnailURL = "/api/upload/" + result.ThumbnailURL
	}

	c.JSON(http.StatusOK, result)
}

func (h *UploadHandler) GetFile(c *gin.Context) {
	filename := c.Param("filename")
	subDir := c.Param("subdir")
	if subDir != "" {
		filename = filepath.Join(subDir, filename)
	}
	filePath := filepath.Join(h.Service.UploadDir, filename)
	c.File(filePath)
}

func (h *UploadHandler) GetThumb(c *gin.Context) {
	filename := c.Param("filename")
	subDir := c.Param("subdir")
	if subDir != "" {
		filename = filepath.Join(subDir, filename)
	}
	filePath := filepath.Join(h.Service.UploadDir, "thumbs", filename)
	c.File(filePath)
}
