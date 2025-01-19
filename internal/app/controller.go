package app

import (
	"fmt"
	"net/http"
)

type RootController struct {
}

func NewRootController() *RootController {
	return &RootController{}
}

func (c *RootController) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the root!")
	fmt.Println("Root handler hello")
}

func (c *RootController) Bru(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the root!")
	fmt.Println("apalah")
}
