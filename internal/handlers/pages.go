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

func (p *Pages) HuntingPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetHuntingPageData()
	if err := p.templates["hunting"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) PackagesPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetPackagesPageData()
	if err := p.templates["packages"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) ContactPage(w http.ResponseWriter, r *http.Request) {
	if err := p.templates["contact"].ExecuteTemplate(w, "base.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
