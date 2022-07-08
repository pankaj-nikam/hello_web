package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("Listening on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
