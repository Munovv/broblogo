package http

import (
	"github.com/Munovv/broblogo/blog-service/pkg/messages"
	"github.com/Munovv/broblogo/blog-service/pkg/models"
	"github.com/Munovv/broblogo/blog-service/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostHandler struct {
	service service.Post
}

type createPostInput struct {
	UserId  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type getPostInput struct {
	UserId string `json:"userId""`
	Guid   string `json:"guid"`
}

type getPostsInput struct {
	UserId string `json:"userId"`
}

func NewPostHandler(service service.Post) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	inp := new(createPostInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var post models.Post
	post.UserId = inp.UserId
	post.Title = inp.Title
	post.Content = inp.Content

	if err := h.service.CreatePost(c, &post); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	inp := new(getPostInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	post, err := h.service.GetPost(c, inp.Guid, inp.UserId)
	if err != nil {
		if err == messages.ErrAccessDenied {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	userGuid, err := getUserGuid(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	posts, err := h.service.GetPosts(c, userGuid)
	if err != nil {
		if err == messages.ErrPostNotFound {
			c.JSON(http.StatusOK, posts)
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) EditPosts(c *gin.Context) {
	var inp models.Post

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.EditPost(c, &inp); err != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.AbortWithStatus(http.StatusOK)
	return
}

func (h *PostHandler) RemovePost(c *gin.Context) {
	var inp models.Post

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.RemovePost(c, inp.GUID, inp.UserId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
	return
}
