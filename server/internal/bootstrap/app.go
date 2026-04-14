package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"

	"pawprints-server/internal/config"
	"pawprints-server/internal/router"
	"pawprints-server/pkg/embedfs"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp() *App {
	cfg := config.Load()

	// 确保数据库目录存在
	dir := filepath.Dir(cfg.Database.Path)
	os.MkdirAll(dir, 0755)

	db, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := Migrate(db); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	if err := Seed(db); err != nil {
		panic("failed to seed database: " + err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	router.Setup(r, db)

	// 嵌入前端 SPA
	embedfs.Register(r)

	return &App{
		Config: cfg,
		DB:     db,
		Router: r,
	}
}

func (a *App) Close() error {
	sqlDB, err := a.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (a *App) Run() error {
	return a.Router.Run(fmt.Sprintf(":%d", a.Config.Server.Port))
}
