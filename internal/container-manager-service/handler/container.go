package handler

import (
	"fmt"
	response "github.com/Munovv/broblogo/internal/pkg/http"
	"github.com/Munovv/broblogo/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) compose(ctx *gin.Context) {
	var input models.ComposeInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.composer.Compose(input.Images); err != nil {
		fmt.Println(err.Error())
		response.NewError(ctx, http.StatusInternalServerError, "failed compose image: "+err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *handler) stop(ctx *gin.Context) {
	var input models.ComposeInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.composer.Down(input.Images); err != nil {
		fmt.Println(err.Error())
		response.NewError(ctx, http.StatusInternalServerError, "failed stop image: "+err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
