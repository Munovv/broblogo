package middleware

import "github.com/spf13/viper"

type Middleware struct {
	AccessToken string
	HashSalt    string
}

func GetAccessToken(viper *viper.Viper) string {
	return viper.GetString("middleware.access_token")
}

func GetHashSalt(viper *viper.Viper) string {
	return viper.GetString("middleware.hash_salt")
}
