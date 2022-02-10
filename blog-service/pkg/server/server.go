package server

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/configs"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *configs.Server, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 10 << 20,
		ReadTimeout:    cfg.ReadTimeout * time.Second,
		WriteTimeout:   cfg.WriteTimeout * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
