package main

import (
	"context"
	"github.com/Munovv/broblogo/internal/blog-service/agent"
	"github.com/Munovv/broblogo/internal/blog-service/config"
	"github.com/Munovv/broblogo/internal/blog-service/handler"
	"github.com/Munovv/broblogo/internal/blog-service/repository"
	"github.com/Munovv/broblogo/internal/blog-service/server"
	"github.com/Munovv/broblogo/internal/blog-service/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация конфигурации проекта
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed while init configs: %s", err.Error())
		return
	}

	// Подключение к базе данных
	db, err := repository.NewDatabase(cfg.Mongo)
	if err != nil {
		log.Fatalf("failed connect to db: %s", err.Error())
		return
	}

	// Инициализация зависимостей
	srv := server.NewServer(
		cfg.Server,
		handler.NewHandler(
			service.NewService(
				repository.NewRepository(db, cfg.Mongo.Collection),
			),
			agent.NewAuthAgent(),
		).InitRoutes(),
	)

	// Запуск сервера
	go func() {
		if err = srv.Run(); err != nil {
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
