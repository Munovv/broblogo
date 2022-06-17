package handler

import (
	"github.com/gin-gonic/gin"
)

type Composer interface {
	Compose(services []string) error
	Down(services []string) error
}

type handler struct {
	composer Composer
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.Cors())

	api := router.Group("/api")
	{
		container := api.Group("/container")
		{
			container.POST("/compose", h.compose)
			container.POST("/stop", h.stop)
		}
	}

	return router
}

func (h *handler) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewHandler(composer Composer) *handler {
	return &handler{
		composer: composer,
	}
}
