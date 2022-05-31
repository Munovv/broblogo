package rest

import (
	"github.com/Munovv/broblogo/auth-service/auth-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.SignUpInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидная форма регистрации")
		return
	}

	if err := h.service.CreateUser(ctx, input.Username, input.Password); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, model.SignUpOutput{
		Message: "success",
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input model.SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(ctx, input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, model.SignInOutput{
		Token: token,
	})
}

func (h *Handler) verify(ctx *gin.Context) {
	var input model.VerifyInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.ParseToken(input.Token)
	if err != nil {
		newErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, model.VerifyOutput{
		UserId: userId,
	})
}
