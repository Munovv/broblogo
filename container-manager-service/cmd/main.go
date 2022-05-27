package main

import (
	"context"
	"github.com/Munovv/broblogo/container-manager-service/internal/composer"
	"github.com/Munovv/broblogo/container-manager-service/internal/config"
	"github.com/Munovv/broblogo/container-manager-service/internal/handler"
	"github.com/Munovv/broblogo/container-manager-service/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig("configs", "config")
	if err != nil {
		log.Fatalf("failed init service configs: %s", err.Error())
		return
	}

	srv := server.NewServer(
		cfg.Server,
		handler.NewHandler(
			composer.NewComposer(),
		).InitRoutes(),
	)

	// Запуск сервера
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
