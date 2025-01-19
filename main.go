package main

import (
	"fmt"
	"go-rest-api-with-postgres/internal/app"
	"go-rest-api-with-postgres/internal/router"
	"net/http"
)

func main() {
	routes := http.NewServeMux()
	root := app.NewRootController()
	config := router.RouterConfig{
		H:    routes,
		Root: root,
	}

	config.Router()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routes,
	}

	fmt.Printf("server running at http://%s \n\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
