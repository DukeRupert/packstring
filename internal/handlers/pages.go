package handlers

import (
	"html/template"
	"net/http"
)

type Pages struct {
	tmpl *template.Template
}

func NewPages(tmpl *template.Template) *Pages {
	return &Pages{tmpl: tmpl}
}

func (p *Pages) HomePage(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "base.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
