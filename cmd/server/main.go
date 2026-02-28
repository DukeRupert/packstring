package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/firefly/packstring/internal/data"
	"github.com/firefly/packstring/internal/db"
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

	// Initialize SQLite database
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/packstring.db"
	}
	store, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer store.Close()
	log.Printf("Database opened at %s", dbPath)

	pages := handlers.NewPages(templates, availability)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// SEO files
	mux.HandleFunc("GET /sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/sitemap.xml")
	})
	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})

	// Pages
	mux.HandleFunc("GET /{$}", pages.HomePage)
	mux.HandleFunc("GET /trips/{$}", pages.TripsHub)
	mux.HandleFunc("GET /trips/fishing/{$}", pages.FishingPage)
	mux.HandleFunc("GET /trips/hunting/{$}", pages.HuntingPage)
	mux.HandleFunc("GET /trips/packages/{$}", pages.PackagesPage)
	mux.HandleFunc("GET /gallery/{$}", pages.GalleryPage)
	mux.HandleFunc("GET /contact/{$}", pages.ContactPage)

	// Contact form
	contact := handlers.NewContact(templates, store)
	mux.HandleFunc("POST /contact", contact.Submit)

	// Admin routes (only if ADMIN_PASSWORD is set)
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword != "" {
		adminTemplates := map[string]*template.Template{
			"admin-login":          mustParseTemplate("admin-login.html"),
			"admin":                mustParseAdminTemplate("admin.html"),
			"admin-dashboard":      mustParseAdminTemplate("admin-dashboard.html"),
			"admin-inquiries":      mustParseAdminTemplate("admin-inquiries.html"),
			"admin-inquiry-detail": mustParseAdminTemplate("admin-inquiry-detail.html"),
			"admin-deposits":       mustParseAdminTemplate("admin-deposits.html"),
		}
		admin := handlers.NewAdmin(adminTemplates, availability, adminPassword, store)

		// Auth
		mux.HandleFunc("GET /admin/login", admin.LoginPage)
		mux.HandleFunc("POST /admin/login", admin.LoginSubmit)
		mux.HandleFunc("POST /admin/logout", admin.Logout)

		// Dashboard
		mux.HandleFunc("GET /admin/{$}", admin.RequireAuth(admin.Dashboard))

		// Availability
		mux.HandleFunc("GET /admin/availability/{$}", admin.RequireAuth(admin.EditPage))
		mux.HandleFunc("POST /admin/availability", admin.RequireAuth(admin.SaveAvailability))

		// Inquiries
		mux.HandleFunc("GET /admin/inquiries/{$}", admin.RequireAuth(admin.InquiriesList))
		mux.HandleFunc("GET /admin/inquiries/{id}", admin.RequireAuth(admin.InquiryDetail))
		mux.HandleFunc("POST /admin/inquiries/{id}/status", admin.RequireAuth(admin.UpdateInquiryStatus))
		mux.HandleFunc("POST /admin/inquiries/{id}/notes", admin.RequireAuth(admin.UpdateInquiryNotes))

		// Deposits
		mux.HandleFunc("GET /admin/deposits/{$}", admin.RequireAuth(admin.DepositsPage))
		mux.HandleFunc("POST /admin/deposits", admin.RequireAuth(admin.SaveDeposits))
		mux.HandleFunc("POST /admin/inquiries/{id}/deposit", admin.RequireAuth(admin.GenerateDepositLink))

		// Stripe webhook (no auth — verified by signature)
		stripe := handlers.NewStripeHandler(store)
		mux.HandleFunc("POST /stripe/webhook", stripe.HandleWebhook)

		// Public payment pages
		paymentTemplates := map[string]*template.Template{
			"payment-success": mustParseTemplate("payment-success.html"),
			"payment-cancel":  mustParseTemplate("payment-cancel.html"),
		}
		mux.HandleFunc("GET /payments/success", handlers.PaymentSuccess(paymentTemplates))
		mux.HandleFunc("GET /payments/cancel", handlers.PaymentCancel(paymentTemplates))

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
