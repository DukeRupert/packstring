package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/firefly/packstring/internal/handlers"
)

// mustParseTemplate builds a template set for a single page file,
// combining it with the base layout and all partials.
func mustParseTemplate(page string) *template.Template {
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "layouts", "*.html")))
	template.Must(tmpl.ParseGlob(filepath.Join("templates", "partials", "*.html")))
	template.Must(tmpl.ParseFiles(filepath.Join("templates", "pages", page)))
	return tmpl
}

func main() {
	// Build a separate template set per page to avoid "content" block collisions
	templates := map[string]*template.Template{
		"home":     mustParseTemplate("home.html"),
		"trips":    mustParseTemplate("trips.html"),
		"fishing":  mustParseTemplate("fishing.html"),
		"hunting":  mustParseTemplate("hunting.html"),
		"packages": mustParseTemplate("packages.html"),
	}

	pages := handlers.NewPages(templates)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages
	mux.HandleFunc("GET /{$}", pages.HomePage)
	mux.HandleFunc("GET /trips/{$}", pages.TripsHub)
	mux.HandleFunc("GET /trips/fishing/{$}", pages.FishingPage)
	mux.HandleFunc("GET /trips/hunting/{$}", pages.HuntingPage)
	mux.HandleFunc("GET /trips/packages/{$}", pages.PackagesPage)

	// Contact form
	contact := handlers.NewContact()
	mux.HandleFunc("POST /contact", contact.Submit)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
