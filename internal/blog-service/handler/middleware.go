package handler

import (
	"errors"
	response "github.com/Munovv/broblogo/internal/pkg/http"
	rest "github.com/Munovv/broblogo/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// corsMiddleware - CORS Middleware
func (h *handler) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// authMiddleware - Authorization Middleware
func (h *handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		response.NewError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		response.NewError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		response.NewError(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.getUserId(headerParts[1])
	if err != nil {
		response.NewError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *handler) getUserId(token string) (string, error) {
	reqBody := rest.AuthServiceRequest{Token: token}

	resp, err := h.agent.VerifyUser(reqBody)
	if err != nil {
		return "", err
	}

	return resp.UserId, nil
}

func (h *handler) getUserIdFromContext(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idStr, nil
}
