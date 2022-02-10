package http

import (
	"github.com/Munovv/broblogo/auth-service/pkg/user"
	"github.com/Munovv/broblogo/auth-service/pkg/user/model"
	"github.com/Munovv/broblogo/auth-service/pkg/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *service.Service
}

type getUserInput struct {
	UserGuid string `json:"user_guid"`
}

type getUserResponse struct {
	User *model.User `json:"user"`
}

type createUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createUserResponse struct {
	UserGuid string `json:"user_guid"`
}

func NewUserHandler(s *service.Service) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	inp := new(getUserInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := h.service.GetUser(c.Request.Context(), inp.UserGuid)
	if err != nil {
		if err == user.ErrUserNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUserResponse{
		User: u,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	inp := new(createUserInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	guid, err := h.service.CreateUser(c.Request.Context(), inp.Username, inp.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createUserResponse{
		UserGuid: guid,
	})
}
