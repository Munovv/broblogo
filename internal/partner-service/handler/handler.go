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

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}
