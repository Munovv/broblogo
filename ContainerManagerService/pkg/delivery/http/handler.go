package http

import (
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	container *Container
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		container: NewContainerHandler(s.Container),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		container := api.Group("/container")
		{
			container.POST("Get")
			container.POST("Create")
			container.POST("Delete")
		}
	}

	return router
}
