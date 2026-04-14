package main

import (
	"log"
	"time"
	"pawprints-server/internal/bootstrap"
	"pawprints-server/internal/task"
)

func main() {
	app := bootstrap.NewApp()
	defer app.Close()

	// 启动提醒检查定时任务（每分钟）
	checker := task.NewReminderChecker(app.DB)
	checker.Start(1 * time.Minute)

	log.Printf("PawPrints Server starting on port %d...", app.Config.Server.Port)
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
