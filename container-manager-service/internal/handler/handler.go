package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Composer interface {
	Compose(ctx context.Context, service string) error
}

type handler struct {
	composer Composer
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		container := api.Group("/container")
		{
			container.POST("/compose", h.compose)
		}
	}

	return router
}

func NewHandler(composer Composer) *handler {
	return &handler{
		composer: composer,
	}
}
