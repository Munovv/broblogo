package config

import "github.com/spf13/viper"

type Middleware struct {
	AccessToken string
	HashSalt    string
}

func (cfg *Middleware) getConfig(v *viper.Viper) *Middleware {
	return &Middleware{
		AccessToken: getAccessToken(v),
		HashSalt:    getHashSalt(v),
	}
}

func getAccessToken(v *viper.Viper) string {
	return v.GetString("middleware.access_token")
}

func getHashSalt(v *viper.Viper) string {
	return v.GetString("middleware.hash_salt")
}
