package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const portNumber = ":3000"

func home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes written:", n)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is about page")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes written:", n)
	}
}

func sum(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		fmt.Fprintf(w, "Error occurred while converting x")
	}
	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		fmt.Fprintf(w, "Error occurred while converting y")
	}
	sum := AddValues(x, y)
	n, err := fmt.Fprintf(w, "The sum of %d and %d is %d", x, y, sum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes written:", n)
	}
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/sum", sum)
	fmt.Println("Listening on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
