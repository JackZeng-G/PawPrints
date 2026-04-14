package embedfs

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var distFS embed.FS

func Register(r *gin.Engine) {
	sub, err := fs.Sub(distFS, "dist")
	if err != nil {
		return
	}

	fileServer := http.FileServer(http.FS(sub))

	// 静态资源 - 用 fileServer 处理 /assets/*
	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = "/assets" + c.Param("filepath")
		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	// 读取 index.html 内容
	indexHTML, _ := fs.ReadFile(sub, "index.html")

	// SPA: 所有非 /api 路由返回 index.html
	spaHandler := func(c *gin.Context) {
		path := c.Request.URL.Path
		// 尝试直接匹配静态文件
		if f, err := sub.Open(path[1:]); err == nil {
			f.Close()
			c.Request.URL.Path = path[1:]
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}
		// SPA fallback: 返回 index.html
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	}

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	})

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) >= 5 && path[:5] == "/api/" {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}
		spaHandler(c)
	})
}
