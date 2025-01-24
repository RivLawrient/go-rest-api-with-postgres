package config

import (
	"database/sql"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/router"
	"net/http"
)

type RegisterConfig struct {
	App *http.ServeMux
	Db  *sql.DB
}

// register parameter yang dibutuhkan pada setiap app
func Register(cfg *RegisterConfig) {
	walletRepository := wallet.NewWalletRepository(cfg.Db)
	walletUsecase := wallet.NewWalletUsecase(walletRepository)
	walletController := wallet.NewWalletController(walletUsecase)

	config := router.RouterConfig{
		Routing:          cfg.App,
		WalletController: walletController,
	}

	config.Route()
}
