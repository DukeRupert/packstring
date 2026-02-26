package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/firefly/packstring/internal/data"
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

// mustParseAdminTemplate builds a template set with the admin FuncMap.
func mustParseAdminTemplate(page string) *template.Template {
	tmpl := template.Must(
		template.New("").Funcs(handlers.AdminFuncMap()).ParseGlob(filepath.Join("templates", "layouts", "*.html")),
	)
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
		"gallery":  mustParseTemplate("gallery.html"),
		"contact":  mustParseTemplate("contact.html"),
	}

	devMode := os.Getenv("PACKSTRING_DEV") == "1"
	availability := data.NewAvailabilityStore("data/availability.yaml", devMode)

	pages := handlers.NewPages(templates, availability)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages
	mux.HandleFunc("GET /{$}", pages.HomePage)
	mux.HandleFunc("GET /trips/{$}", pages.TripsHub)
	mux.HandleFunc("GET /trips/fishing/{$}", pages.FishingPage)
	mux.HandleFunc("GET /trips/hunting/{$}", pages.HuntingPage)
	mux.HandleFunc("GET /trips/packages/{$}", pages.PackagesPage)
	mux.HandleFunc("GET /gallery/{$}", pages.GalleryPage)
	mux.HandleFunc("GET /contact/{$}", pages.ContactPage)

	// Contact form
	contact := handlers.NewContact(templates)
	mux.HandleFunc("POST /contact", contact.Submit)

	// Admin routes (only if ADMIN_PASSWORD is set)
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword != "" {
		adminTemplates := map[string]*template.Template{
			"admin-login": mustParseTemplate("admin-login.html"),
			"admin":       mustParseAdminTemplate("admin.html"),
		}
		admin := handlers.NewAdmin(adminTemplates, availability, adminPassword)
		mux.HandleFunc("GET /admin/login", admin.LoginPage)
		mux.HandleFunc("POST /admin/login", admin.LoginSubmit)
		mux.HandleFunc("POST /admin/logout", admin.Logout)
		mux.HandleFunc("GET /admin/availability/{$}", admin.RequireAuth(admin.EditPage))
		mux.HandleFunc("POST /admin/availability", admin.RequireAuth(admin.SaveAvailability))
		log.Println("Admin routes registered at /admin/")
	} else {
		log.Println("ADMIN_PASSWORD not set — admin routes disabled")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
