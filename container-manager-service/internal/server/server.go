package server

import (
	"context"
	"github.com/Munovv/broblogo/container-manager-service/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// server - структура HTTP сервера
type server struct {
	httpServer *http.Server
}

// Start - запуск HTTP сервера
func (s *server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown - остановка HTTP сервера
func (s *server) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(shutdownCtx)
}

func NewServer(cfg config.Server, handler *gin.Engine) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			MaxHeaderBytes: cfg.MaxWriteMegabytes << 20,
			ReadTimeout:    cfg.ReadTimeout * time.Second,
			WriteTimeout:   cfg.WriteTimeout * time.Second,
		},
	}
}
