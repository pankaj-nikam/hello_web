package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//RenderTemplate renders the template using html/template
func RenderTemplate(w http.ResponseWriter, templatePath string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templatePath, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err.Error())
	}
}
