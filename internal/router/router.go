package router

import (
	"fmt"
	"net/http"
)

type RouterConfig struct {
	Routing *http.ServeMux
}

// membungkus semua endpoint yang dibuat.
// menghindari penggunaan root endpoint ("/") agar menghasilkan 404 page not found.
func (c *RouterConfig) Route() {

	c.Routing.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/hello")
	})
}
