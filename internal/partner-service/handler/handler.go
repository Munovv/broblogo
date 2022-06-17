package handler

import (
	"context"
	"github.com/Munovv/broblogo/internal/partner-service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreatePartner(ctx context.Context, input *models.CreatePartnerInput) error
	GetPartner(ctx context.Context, id string) (*models.Partner, error)
	GetPartners(ctx context.Context) ([]*models.Partner, error)
	DeletePartner(ctx context.Context, id string) error
}

type handler struct {
	service Service
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.Cors())

	api := router.Group("/api")
	{
		partner := api.Group("/partner")
		{
			partner.POST("/", h.createPartner)
			partner.GET("/", h.getPartners)
			partner.DELETE("/", h.deletePartner)
		}
	}

	return router
}

func (h *handler) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}
