package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pankaj-nikam/hello_web/pkg/config"
	"github.com/pankaj-nikam/hello_web/pkg/handlers"
	"github.com/pankaj-nikam/hello_web/pkg/render"
)

const portNumber = ":3000"

func main() {
	app := config.AppConfig{}
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Listening on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
