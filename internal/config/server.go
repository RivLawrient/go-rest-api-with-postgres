package config

import (
	"net/http"

	"github.com/spf13/viper"
)

func NewServer(viper *viper.Viper, app *http.ServeMux) *http.Server {
	server := http.Server{
		Addr:    viper.GetString("SERVER_ADDR"),
		Handler: app,
	}

	return &server
}
