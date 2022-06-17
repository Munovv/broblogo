package handler

import (
	"github.com/Munovv/broblogo/internal/partner-service/models"
	response "github.com/Munovv/broblogo/internal/pkg/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) createPartner(ctx *gin.Context) {
	var input models.CreatePartnerInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreatePartner(ctx, &input); err != nil {
		response.NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *handler) getPartners(ctx *gin.Context) {
	partners, err := h.service.GetPartners(ctx)
	if err != nil {
		response.NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, partners)
}

func (h *handler) deletePartner(ctx *gin.Context) {
	var input models.DeletePartnerInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := h.service.GetPartner(ctx, input.Id)
	if err != nil {
		response.NewError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err = h.service.DeletePartner(ctx, input.Id); err != nil {
		response.NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
