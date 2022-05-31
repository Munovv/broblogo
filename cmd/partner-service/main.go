package main

import (
	"context"
	"fmt"
	"github.com/Munovv/broblogo/internal/partner-service/config"
	"github.com/Munovv/broblogo/internal/partner-service/handler"
	"github.com/Munovv/broblogo/internal/partner-service/repository"
	"github.com/Munovv/broblogo/internal/partner-service/server"
	"github.com/Munovv/broblogo/internal/partner-service/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.NewConfig("configs", "config")
	if err != nil {
		log.Fatalf("failed config init: %s", err.Error())
		return
	}

	fmt.Printf("%v", cfg)

	// Инициализация базы данных
	db, err := repository.NewMongoDatabase(cfg.Mongo)
	if err != nil {
		log.Fatalf("failed database connection: %s", err.Error())

		return
	}

	// Инициализация зависимостей
	srv := server.NewServer(
		cfg.Server,
		handler.NewHandler(
			service.NewService(
				repository.NewRepository(db, cfg.Mongo.Collection),
			),
		).InitRoutes(),
	)

	// Запуск REST сервера
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("an error occurred while running http server: %s", err.Error())
			return
		}
	}()
	log.Println("Server has been started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	log.Println("Server has been stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("an error occurred on server shutdown: %s", err.Error())
		return
	}
	log.Print("Server has been exited properly")
}
