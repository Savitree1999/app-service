package router

import (
	"github.com/Savitree1999/app-service/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	v1 "github.com/Savitree1999/app-service/internal/handler/v1"
)

func SetupRouter(cfg *config.Config, logger *zap.SugaredLogger, db *gorm.DB) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Middleware
	r.Use(gin.Recovery())
	r.Use(GinLoggerMiddleware(logger))
	r.Use(cors.Default()) // อนุญาตทุก origin (ใน dev)

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	apiV1 := r.Group("/api/v1")
	{
		v1.RegisterRoutes(apiV1, db, logger)
	}

	return r
}
