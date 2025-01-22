package main

import (
	"fmt"
	"go-rest-api-with-postgres/internal/config"
	"net/http"
)

func main() {
	app := http.NewServeMux()
	viper := config.NewViper()
	db := config.GetConnection(viper)
	server := config.NewServer(viper, app)

	config.Register(&config.RegisterConfig{
		App: app,
		Db:  db,
	})

	fmt.Printf("server running at http://%s \n\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
