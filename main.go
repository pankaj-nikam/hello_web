package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bytes written:", n)
		}
	})
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", nil)
}
