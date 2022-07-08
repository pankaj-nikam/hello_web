package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":3000"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, templatePath string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templatePath)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err.Error())
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("Listening on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
