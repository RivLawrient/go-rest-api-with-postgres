package config

import (
	"net/http"

	"github.com/spf13/viper"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Atur header CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func NewServer(viper *viper.Viper, app *http.ServeMux) *http.Server {
	cors := EnableCORS(app)
	server := http.Server{
		Addr:    viper.GetString("SERVER_ADDR"),
		Handler: cors,
	}

	return &server
}
