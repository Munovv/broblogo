package handler

import (
	"github.com/Munovv/broblogo/partner-service/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) createPartner(ctx *gin.Context) {
	var input models.CreatePartnerInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreatePartner(ctx, &input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *handler) getPartners(ctx *gin.Context) {
	partners, err := h.service.GetPartners(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, partners)
}

func (h *handler) deletePartner(ctx *gin.Context) {
	var input models.DeletePartnerInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := h.service.GetPartner(ctx, input.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err = h.service.DeletePartner(ctx, input.Id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
