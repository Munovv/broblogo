package config

import (
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (cfg *Server) getConfig(v *viper.Viper) *Server {
	return &Server{
		Port:         getServerPort(v),
		ReadTimeout:  getServerReadTimeout(v),
		WriteTimeout: getServerWriteTimeout(v),
	}
}

func getServerPort(v *viper.Viper) string {
	return v.GetString("server.port")
}

func getServerReadTimeout(v *viper.Viper) time.Duration {
	return v.GetDuration("server.read_timeout")
}

func getServerWriteTimeout(v *viper.Viper) time.Duration {
	return v.GetDuration("server.write_timeout")
}
