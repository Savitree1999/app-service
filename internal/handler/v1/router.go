package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, logger *zap.SugaredLogger) {
	// ตัวอย่าง route ย่อย
	// userRoutes := r.Group("/users")
	// userRoutes.GET("/", ...)
}
