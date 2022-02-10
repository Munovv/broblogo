package main

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/configs"
	"github.com/Munovv/broblogo/blog-service/pkg/delivery/http"
	"github.com/Munovv/broblogo/blog-service/pkg/repository"
	"github.com/Munovv/broblogo/blog-service/pkg/server"
	"github.com/Munovv/broblogo/blog-service/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := configs.Init(viper.GetViper())
	if err != nil {
		log.Fatalf("messages occurred while init app configs with message :%s", err.Error())
	}

	db := repository.NewMongoDb(cfg.Mongo)
	repo := repository.NewMongoRepository(db, cfg.Mongo)
	services := service.NewService(repo)
	handler := http.NewHandler(services)

	s := new(server.Server)
	go func() {
		if err := s.Run(cfg.Server, handler.InitRoutes()); err != nil {
			log.Fatalf("messages occurred while running http server with message: %s", err.Error())
		}
	}()
	log.Println("Server has been running")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	log.Println("Server has been stopped")

	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatalf("messages occurred while stopping http server with message: %s", err.Error())
	}
	log.Print("Server has been exited properly")
}
