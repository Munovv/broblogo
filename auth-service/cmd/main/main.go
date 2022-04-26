package main

import (
	"context"
	"github.com/Munovv/broblogo/auth-service/internal/config"
	"github.com/Munovv/broblogo/auth-service/internal/handler/rest"
	"github.com/Munovv/broblogo/auth-service/internal/repository"
	"github.com/Munovv/broblogo/auth-service/internal/server"
	"github.com/Munovv/broblogo/auth-service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация конфигурации проекта
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("an error occurred while init configs: %s", err.Error())
		return
	}

	// Подключение к mongo
	db, err := repository.InitDb(cfg.Mongo)
	if err != nil {
		return
	}

	// Инициализация зависимостей
	repo := repository.NewRepository(db, cfg.Mongo.Collection)
	service := service.NewService(repo)
	handler := rest.NewHandler(service)

	// Создание маршрутов
	router := handler.InitRoutes()

	srv := new(server.Server)
	go func() {
		if err = srv.Run(cfg.Server, router); err != nil {
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
