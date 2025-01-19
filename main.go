package main

import (
	"fmt"
	"go-rest-api-with-postgres/internal/config"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	config.Register(&config.RegisterConfig{
		App: app,
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: app,
	}

	fmt.Printf("server running at http://%s \n\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
