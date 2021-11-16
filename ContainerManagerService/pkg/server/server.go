package server

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *config.Server, handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 10 << 20,
		ReadTimeout:    cfg.ReadTimeout * time.Second,
		WriteTimeout:   cfg.ReadTimeout * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
