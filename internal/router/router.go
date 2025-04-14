package router

import (
	"net/http"

	v1 "github.com/Savitree1999/app-service/internal/handler/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Savitree1999/app-service/internal/config"
	"github.com/Savitree1999/app-service/internal/middleware"
	reqLogger "github.com/Savitree1999/app-service/internal/logger"
)

func SetupRouter(cfg *config.Config, logger *zap.SugaredLogger, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Custom request logger middleware
	requestLogger, requestSugar, _ := reqLogger.NewRequestLogger()

	defer requestLogger.Sync()
	r.Use(middleware.RequestLoggerMiddleware(requestSugar))

	r.Use(cors.Default())

	v1Group := r.Group("/v1")

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// --- User handler ---
	userHandler := v1.NewUserHandler(db, logger)
	v1Group.POST("/users", userHandler.CreateUser)

	return r
}
