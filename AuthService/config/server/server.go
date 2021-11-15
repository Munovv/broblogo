package server

import "github.com/spf13/viper"

type ServerConf struct {
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
