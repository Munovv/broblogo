package handler

import (
	response "github.com/Munovv/broblogo/internal/pkg/http"
	rest "github.com/Munovv/broblogo/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input rest.SignUpInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "невалидная форма регистрации")
		return
	}

	if err := h.service.CreateUser(ctx, input.Username, input.Password); err != nil {
		response.NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, rest.SignUpOutput{
		Message: "success",
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input rest.SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(ctx, input.Username, input.Password)
	if err != nil {
		response.NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, rest.SignInOutput{
		Token: token,
	})
}

func (h *Handler) verify(ctx *gin.Context) {
	var input rest.VerifyInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.ParseToken(input.Token)
	if err != nil {
		response.NewError(ctx, http.StatusForbidden, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, rest.VerifyOutput{
		UserId: userId,
	})
}
