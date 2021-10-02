package handlers

import (
	"github.com/marstan/bookings/pkg/config"
	"github.com/marstan/bookings/pkg/models"
	"github.com/marstan/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func SetRepo(r *Repository)  {
	Repo = r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test": "test123",
	}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
