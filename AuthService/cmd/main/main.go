package main

import (
	"context"
	"github.com/Munovv/broblogo/AuthService/config"
	"github.com/Munovv/broblogo/AuthService/pkg"
	"github.com/Munovv/broblogo/AuthService/pkg/user/http"
	"github.com/Munovv/broblogo/AuthService/pkg/user/repository/mongo"
	"github.com/Munovv/broblogo/AuthService/pkg/user/service"
	"github.com/Munovv/broblogo/AuthService/server"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init(viper.GetViper())
	cfgDb := config.InitMongoDb(viper.GetViper())
	cfgMw := config.InitMiddleware(viper.GetViper())

	db := pkg.InitDb(cfgDb)

	userRepo := mongo.NewUserRepository(db, cfgDb.Collection)
	userService := service.NewUserService(userRepo, cfgMw.HashSalt)
	userHandler := http.NewUserHandler(userService)
	mainHandler := pkg.NewHandler(userHandler)

	router := mainHandler.InitRoutes()

	if err := config.Init(viper.GetViper()); err != nil {
		log.Fatalf("an error occurred on init main config: %s", err.Error())
	}
	srvCnf := config.InitServer(viper.GetViper())

	srv := new(server.Server)
	go func() {
		if err := srv.Run(srvCnf.Port, router); err != nil {
			log.Fatalf("an error occurred while running http server: %s", err.Error())
		}
	}()
	log.Println("Server has been started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	log.Println("Server has been stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("an error occurred on server shutdown: %s", err.Error())
	}
	log.Print("Server has been exited properly")
}
