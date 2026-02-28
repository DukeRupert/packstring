package handlers

import (
	"html/template"
	"net/http"

	"github.com/firefly/packstring/internal/data"
)

type Pages struct {
	templates    map[string]*template.Template
	availability *data.AvailabilityStore
}

func NewPages(templates map[string]*template.Template, availability *data.AvailabilityStore) *Pages {
	return &Pages{templates: templates, availability: availability}
}

// attachAvailability populates the Availability field on each TripSection from the store.
func (p *Pages) attachAvailability(trips []data.TripSection) {
	for i := range trips {
		trips[i].Availability = p.availability.Get(trips[i].Slug)
	}
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
	p.attachAvailability(pageData.Trips)
	if err := p.templates["fishing"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) HuntingPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetHuntingPageData()
	p.attachAvailability(pageData.Trips)
	if err := p.templates["hunting"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) PackagesPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetPackagesPageData()
	p.attachAvailability(pageData.Packages)
	if err := p.templates["packages"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) GalleryPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetGalleryPageData()
	if err := p.templates["gallery"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (p *Pages) ContactPage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetContactPageData()
	if err := p.templates["contact"].ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
