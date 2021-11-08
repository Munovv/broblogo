package main

import (
	"context"
	"github.com/Munovv/broblogo/GoMainService/pkg/config"
	"github.com/Munovv/broblogo/GoMainService/pkg/server"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := initConfig(viper.GetViper()); err != nil {
		log.Fatalf("an error occurred on init main config: %s", err.Error())
	}
	srvCnf := initServerConfig(viper.GetViper())

	srv := new(server.Server)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(srvCnf.Port); err != nil {
			log.Fatalf("an error occurred while running http server: %s", err.Error())
		}
	}()
	log.Println("Server has been started")

	<-done
	log.Println("Server has been stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("an error occurred on server shutdown: %s", err.Error())
	}
	log.Print("Server has been exited properly")

}

func initConfig(v *viper.Viper) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func initServerConfig(v *viper.Viper) config.ServerConf {
	return config.ServerConf{
		Port:         config.GetServerPort(v),
		ReadTimeout:  config.GetServerReadTimeout(v),
		WriteTimeout: config.GetServerWriteTimeout(v),
	}
}
