package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, templatePath string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templatePath, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err.Error())
	}
}

var tc = make(map[string]*template.Template)

//RenderTemplate renders the template using html/template
func RenderTemplate(w http.ResponseWriter, templatePath string) {
	var tmpl *template.Template
	var err error

	//Check to see if we already have the template in cache
	_, inMap := tc[templatePath]
	if !inMap {
		log.Println("creating template cache")
		//need to create template and put in cache
		err = createTemplateCache(templatePath)
		if err != nil {
			log.Fatal("error", err.Error())
			return
		}
	} else {
		//We have template in cache
		log.Println("using cached template")
	}

	tmpl = tc[templatePath]

	tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	tmp := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(tmp...)
	if err != nil {
		return err
	} else {
		tc[t] = tmpl
	}
	return nil
}
