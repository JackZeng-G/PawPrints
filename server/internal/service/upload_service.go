package service

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

type UploadService struct {
	UploadDir    string
	MaxWidth     int
	MaxHeight    int
	ThumbnailDir string
}

func NewUploadService(uploadDir string, maxWidth, maxHeight int) *UploadService {
	return &UploadService{
		UploadDir:    uploadDir,
		MaxWidth:     maxWidth,
		MaxHeight:    maxHeight,
		ThumbnailDir: filepath.Join(uploadDir, "thumbs"),
	}
}

type UploadResult struct {
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Filename     string `json:"filename"`
}

func (s *UploadService) Upload(file multipart.File, header *multipart.FileHeader) (*UploadResult, error) {
	now := time.Now()
	subDir := now.Format("2006-01")
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}

	filename := fmt.Sprintf("%d%s", now.UnixNano(), ext)

	origDir := filepath.Join(s.UploadDir, subDir)
	if err := os.MkdirAll(origDir, 0755); err != nil {
		return nil, err
	}

	origPath := filepath.Join(origDir, filename)
	dst, err := os.Create(origPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	thumbURL, err := s.generateThumbnail(origPath, subDir, filename)
	if err != nil {
		thumbURL = "" // 缩略图生成失败不阻止上传
	}

	return &UploadResult{
		URL:          filepath.Join(subDir, filename),
		ThumbnailURL: thumbURL,
		Filename:     filename,
	}, nil
}

func (s *UploadService) generateThumbnail(origPath, subDir, filename string) (string, error) {
	img, err := imaging.Open(origPath)
	if err != nil {
		return "", err
	}

	thumb := imaging.Fit(img, s.MaxWidth, s.MaxHeight, imaging.Lanczos)

	thumbDir := filepath.Join(s.ThumbnailDir, subDir)
	if err := os.MkdirAll(thumbDir, 0755); err != nil {
		return "", err
	}

	thumbPath := filepath.Join(thumbDir, filename)
	if err := saveImage(thumb, thumbPath); err != nil {
		return "", err
	}

	return filepath.Join("thumbs", subDir, filename), nil
}

func saveImage(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(f, img, &jpeg.Options{Quality: 80})
	case ".png":
		return png.Encode(f, img)
	default:
		return jpeg.Encode(f, img, &jpeg.Options{Quality: 80})
	}
}
