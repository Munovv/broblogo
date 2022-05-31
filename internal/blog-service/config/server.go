package config

import "github.com/spf13/viper"

type Server struct {
	Port         string
	ReadTimeout  int64
	WriteTimeout int64
}

func GetServerPort(v *viper.Viper) string {
	return v.GetString("server.port")
}

func GetServerReadTimeout(v *viper.Viper) int64 {
	return v.GetInt64("server.read_timeout")
}

func GetServerWriteTimeout(v *viper.Viper) int64 {
	return v.GetInt64("server.write_timeout")
}

func initServer(v *viper.Viper) *Server {
	return &Server{
		Port:         GetServerPort(v),
		ReadTimeout:  GetServerReadTimeout(v),
		WriteTimeout: GetServerWriteTimeout(v),
	}
}
