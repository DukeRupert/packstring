package handlers

import (
	"html/template"
	"net/http"

	"github.com/firefly/packstring/internal/data"
)

type Pages struct {
	tmpl *template.Template
}

func NewPages(tmpl *template.Template) *Pages {
	return &Pages{tmpl: tmpl}
}

func (p *Pages) HomePage(w http.ResponseWriter, r *http.Request) {
	pageData := data.GetHomePageData()
	if err := p.tmpl.ExecuteTemplate(w, "base.html", pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
