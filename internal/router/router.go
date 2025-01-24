package router

import (
	"go-rest-api-with-postgres/internal/app/wallet"
	"net/http"
)

type RouterConfig struct {
	Routing          *http.ServeMux
	WalletController *wallet.WalletController
}

// membungkus semua endpoint yang dibuat.
// menghindari penggunaan root endpoint ("/") agar menghasilkan 404 page not found.
func (c *RouterConfig) Route() {

	c.Routing.HandleFunc("/wallet", c.WalletController.HandleNewWallet)
}
