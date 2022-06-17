package handler

import (
	response "github.com/Munovv/broblogo/internal/pkg/http"
	rest "github.com/Munovv/broblogo/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) createPost(ctx *gin.Context) {
	userId, err := h.getUserIdFromContext(ctx)
	if err != nil {
		response.NewError(ctx, http.StatusInternalServerError, "пустой идентификатор пользователя")
		return
	}

	var input rest.CreatePostInput

	if err = ctx.BindJSON(&input); err != nil {
		response.NewError(ctx, http.StatusBadRequest, "невалидная форма")
		return
	}

	if err = h.service.CreateItem(ctx, input, userId); err != nil {
		response.NewError(ctx, http.StatusInternalServerError, "ошибка создания поста")
	}

	ctx.Done()
}

func (h *handler) getPost(ctx *gin.Context) {

}

func (h *handler) getPosts(ctx *gin.Context) {

}

func (h *handler) updatePost(ctx *gin.Context) {

}

func (h *handler) deletePost(ctx *gin.Context) {

}
