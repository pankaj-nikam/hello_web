package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, templatePath string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templatePath)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err.Error())
	}
}
