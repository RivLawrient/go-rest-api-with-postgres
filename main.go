package main

import (
	"fmt"
	"net/http"
)

// Handler untuk root
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the root!")
	fmt.Println("Root handler called")
}

// Handler untuk /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
	fmt.Println("Hello handler called")
}

// Handler untuk /about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
	fmt.Println("About handler called")
}

// membungus semua endpoint
func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	return mux
}

func main() {
	// create config server like port and ip
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: Router(),
	}

	fmt.Printf("server running at http://%s \n\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
