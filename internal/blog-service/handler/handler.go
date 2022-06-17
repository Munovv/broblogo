package handler

import (
	"context"
	"github.com/Munovv/broblogo/internal/blog-service/model"
	rest "github.com/Munovv/broblogo/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateItem(ctx context.Context, input rest.CreatePostInput, userId string) error
	GetItem(ctx context.Context, id string) (*model.Post, error)
	GetItems(ctx context.Context, userId string) ([]model.Post, error)
	UpdateItem(ctx context.Context, id string, input rest.UpdatePostInput) error
	DeleteItem(ctx context.Context, id string) error
}

type Agent interface {
	VerifyUser(in rest.AuthServiceRequest) (rest.AuthServiceResponse, error)
}

type handler struct {
	service Service
	agent   Agent
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.corsMiddleware())

	api := router.Group("/api", h.authMiddleware)
	{
		blog := api.Group("blog")
		{
			blog.POST("/", h.createPost)
			blog.GET("/:id", h.getPost)
			blog.GET("/", h.getPosts)
			blog.PUT("/:id", h.updatePost)
			blog.DELETE("/:id", h.deletePost)
		}
	}

	return router
}

func NewHandler(service Service, agent Agent) *handler {
	return &handler{
		service: service,
		agent:   agent,
	}
}
