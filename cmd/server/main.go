package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/firefly/packstring/internal/handlers"
)

func main() {
	// Parse templates: each page template is combined with the base layout
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "layouts", "*.html")))
	template.Must(tmpl.ParseGlob(filepath.Join("templates", "pages", "*.html")))

	pages := handlers.NewPages(tmpl)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages
	mux.HandleFunc("GET /{$}", pages.HomePage)

	// Contact form
	contact := handlers.NewContact()
	mux.HandleFunc("POST /contact", contact.Submit)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
