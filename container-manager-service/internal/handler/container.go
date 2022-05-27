package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) compose(c *gin.Context) {
	c.Status(http.StatusOK)
}
