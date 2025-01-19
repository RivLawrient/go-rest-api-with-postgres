package router

import (
	"go-rest-api-with-postgres/internal/app"
	"net/http"
)

type RouterConfig struct {
	H    *http.ServeMux
	Root *app.RootController
}

// membungkus semua endpoint yang dibuat.
// menghindari penggunaan root endpoint ("/") agar menghasilkan 404 page not found.
func (c *RouterConfig) Router() {
	c.H.HandleFunc("/hello", c.Root.RootHandler)
	c.H.HandleFunc("/lol", c.Root.Bru)
}
