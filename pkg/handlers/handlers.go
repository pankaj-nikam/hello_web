package handlers

import (
	"net/http"

	"github.com/pankaj-nikam/hello_web/pkg/config"
	"github.com/pankaj-nikam/hello_web/pkg/models"
	"github.com/pankaj-nikam/hello_web/pkg/render"
)

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//Repo the repository used by handlers
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"
	//Send the data received to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
