package handler

import (
	"github.com/Munovv/broblogo/internal/blog-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) createPost(ctx *gin.Context) {
	userId, err := h.getUserIdFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "пустой идентификатор пользователя")
		return
	}

	var input model.CreatePostInput

	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидная форма")
		return
	}

	if err = h.service.CreateItem(ctx, input, userId); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "ошибка создания поста")
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
