package http

import (
	"github.com/Munovv/broblogo/blog-service/pkg/messages"
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userGuid"
)

func getUserGuid(c *gin.Context) (string, error) {
	guid, ok := c.Get(userCtx)
	if !ok {
		return "", messages.ErrUserNotFound
	}

	guidStr, ok := guid.(string)
	if !ok {
		return "", messages.ErrUserIdInvalidType
	}

	return guidStr, nil
}
