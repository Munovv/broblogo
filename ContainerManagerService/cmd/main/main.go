package main

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/config"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/delivery/http"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/repository"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/server"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Init(viper.GetViper())
	if err != nil {
		log.Fatalf("an error occurred init configs: %s", err.Error())
	}

	db := repository.NewMongoDb(cfg.Mongo)
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := http.NewHandler(service)

	router := handler.InitRoutes()

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Server, router); err != nil {
			log.Fatalf("an error occurred while running http server: %s", err.Error())
		}
	}()
	log.Println("Server has been started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	log.Println("Server has been stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("an error occurred on server shutting down %s", err.Error())
	}
	log.Print("Server has been exited properly")
}
