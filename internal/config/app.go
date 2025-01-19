package config

import (
	"database/sql"
	"go-rest-api-with-postgres/internal/router"
	"net/http"
)

type RegisterConfig struct {
	App *http.ServeMux
	Db  *sql.DB
}

// register parameter yang dibutuhkan pada setiap app
func Register(cfg *RegisterConfig) {
	
	config := router.RouterConfig{
		Routing: cfg.App,
	}

	config.Route()
}
