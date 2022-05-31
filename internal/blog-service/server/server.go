package server

import (
	"context"
	cfg "github.com/Munovv/broblogo/blog-service/blog-service/config/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(cfg *cfg.ServerConf, handler *gin.Engine) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}
