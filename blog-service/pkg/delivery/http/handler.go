package http

import (
	"github.com/Munovv/broblogo/blog-service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*PostHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		PostHandler: NewPostHandler(service.Post),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		blog := api.Group("/blog")
		{
			blog.POST("/", h.PostHandler.CreatePost)
			blog.GET("/:id", h.PostHandler.GetPost)
			blog.GET("/", h.PostHandler.GetPosts)
			blog.PUT("/:id", h.PostHandler.EditPosts)
			blog.DELETE("/:id", h.PostHandler.RemovePost)
		}
	}

	return router
}
