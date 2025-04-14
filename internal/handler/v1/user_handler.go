package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Savitree1999/app-service/internal/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	Serv *service.UserService
	Log  *zap.SugaredLogger
}

func NewUserHandler(db *gorm.DB, log *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		Serv: service.NewUserService(db, log),
		Log:  log,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req service.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Serv.CreateUser(req)
	if err != nil {
		if err == service.ErrEmailExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		h.Log.Errorf("CreateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
