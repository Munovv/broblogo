package server

import (
	"context"
	"github.com/Munovv/broblogo/auth-service/internal/config/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *server.ServerConf, handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
