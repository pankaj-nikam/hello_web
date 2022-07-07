package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes written:", n)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is about page")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes written:", n)
	}
}

func Sum(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 3)
	n, err := fmt.Fprintf(w, "The sum is %d", sum)
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
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/sum", Sum)
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", nil)
}
