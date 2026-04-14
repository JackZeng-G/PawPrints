package main

import (
	"log"
	"pawprints-server/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	defer app.Close()

	log.Printf("PawPrints Server starting on port %d...", app.Config.Server.Port)
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
