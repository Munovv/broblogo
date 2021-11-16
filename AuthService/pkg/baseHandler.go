package pkg

import (
	"github.com/Munovv/broblogo/AuthService/pkg/user/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler *http.UserHandler
}

func NewHandler(userHandler *http.UserHandler) *Handler {
	return &Handler{
		userHandler: userHandler,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/GetUser", h.userHandler.GetUser)
			user.POST("/CreateUser", h.userHandler.CreateUser)
		}
	}

	return router
}
