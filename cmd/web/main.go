package main

import (
	"fmt"
	"net/http"

	"github.com/pankaj-nikam/hello_web/cmd/pkg/handlers"
)

const portNumber = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Listening on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
