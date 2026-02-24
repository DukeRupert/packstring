package handlers

import (
	"html/template"
	"net/http"

	"github.com/firefly/packstring/internal/data"
)

type Pages struct {
	templates map[string]*template.Template
}

func NewPages(templates map[string]*template.Template) *Pages {
	return &Pages{templates: templates}
}

func (p *Pages) HomePage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetHomePageData()
	if err := p.templates["home"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) TripsHub(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetTripsHubData()
	if err := p.templates["trips"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) FishingPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetFishingPageData()
	if err := p.templates["fishing"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
