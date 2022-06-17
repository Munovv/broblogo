package main

import (
	"context"
	"github.com/Munovv/broblogo/internal/auth-service/handler"
	"github.com/Munovv/broblogo/internal/auth-service/repository"
	"github.com/Munovv/broblogo/internal/auth-service/service"
	"github.com/Munovv/broblogo/internal/pkg/config"
	"github.com/Munovv/broblogo/internal/pkg/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig("configs", "config")
	if err != nil {
		log.Fatalf("failed config init: %s", err.Error())
		return
	}

	db, err := repository.InitDb(&cfg.Mongo)
	if err != nil {
		return
	}

	srv := http.NewServer(
		cfg.Server,
		handler.NewHandler(
			service.NewService(
				repository.NewRepository(db, cfg.Mongo.Collection),
			),
		).InitRoutes(),
	)

	go func() {
		if err = srv.Start(); err != nil {
			log.Fatalf("an error occurred while running http server: %s", err.Error())
			return
		}
	}()
	log.Println("Server has been started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	log.Println("Server has been stopped")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("an error occurred on server shutdown: %s", err.Error())
		return
	}
	log.Print("Server has been exited properly")
}
