package main

import (
	"log"

	"github.com/Savitree1999/app-service/internal/config"
	"github.com/Savitree1999/app-service/internal/db"
	"github.com/Savitree1999/app-service/internal/logger"
	"github.com/Savitree1999/app-service/internal/router"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Init logger
	sugar := logger.NewLogger(cfg.App.Env)
	sugar.Infof("Starting app in %s mode", cfg.App.Env)

	// Connect DB
	database, err := db.Connect(cfg)
	if err != nil {
		sugar.Fatalf("Failed to connect database: %v", err)
	}
	defer database.Close()

	// Init Gin
	r := router.SetupRouter(cfg, sugar, database)

	// Start server
	sugar.Infof("Listening on port %s", cfg.App.Port)
	if err := r.Run(":" + cfg.App.Port); err != nil {
		sugar.Fatalf("Server failed to start: %v", err)
	}
}
