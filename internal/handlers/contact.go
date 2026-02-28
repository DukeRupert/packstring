package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/firefly/packstring/internal/data"
	"github.com/firefly/packstring/internal/db"
)

type Contact struct {
	templates map[string]*template.Template
	store     *db.Store // nil if no database configured
}

func NewContact(templates map[string]*template.Template, store *db.Store) *Contact {
	return &Contact{templates: templates, store: store}
}

func (c *Contact) Submit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Honeypot check — if the hidden "website" field has a value, it's a bot
	if r.FormValue("website") != "" {
		// Silently return success to avoid tipping off the bot
		c.renderSuccess(w, data.ContactSuccessData{})
		return
	}

	// Validate required fields
	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))

	var errors []string
	if name == "" {
		errors = append(errors, "Name is required.")
	}
	if email == "" {
		errors = append(errors, "Email is required.")
	} else if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		errors = append(errors, "Please enter a valid email address.")
	}

	if len(errors) > 0 {
		c.renderErrors(w, errors)
		return
	}

	tripSlug := strings.TrimSpace(r.FormValue("trip"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	dates := strings.TrimSpace(r.FormValue("dates"))
	partySize := strings.TrimSpace(r.FormValue("party_size"))
	experience := strings.TrimSpace(r.FormValue("experience"))
	message := strings.TrimSpace(r.FormValue("message"))
	tripName := data.TripDisplayName(tripSlug)

	// Store in database if available
	if c.store != nil {
		inq := &db.Inquiry{
			Name:       name,
			Email:      email,
			Phone:      phone,
			TripSlug:   tripSlug,
			TripName:   tripName,
			Dates:      dates,
			PartySize:  partySize,
			Experience: experience,
			Message:    message,
		}
		id, err := c.store.CreateInquiry(inq)
		if err != nil {
			log.Printf("[contact] DB error: %v", err)
		} else {
			log.Printf("[contact] inquiry #%d from %s <%s> — trip: %s", id, name, email, tripSlug)
		}
	} else {
		log.Printf("Contact inquiry from %s <%s> — trip: %s", name, email, tripSlug)
	}

	c.renderSuccess(w, data.ContactSuccessData{
		Name:      name,
		Email:     email,
		Trip:      tripName,
		Dates:     dates,
		PartySize: partySize,
	})
}

func (c *Contact) renderSuccess(w http.ResponseWriter, successData data.ContactSuccessData) {
	if err := c.templates["contact"].ExecuteTemplate(w, "contact-success", successData); err != nil {
		log.Printf("Error rendering contact success: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *Contact) renderErrors(w http.ResponseWriter, errors []string) {
	// Return an OOB swap targeting #form-errors so the form itself is preserved.
	// htmx swaps this into the existing #form-errors div without touching the rest of the form.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var b strings.Builder
	b.WriteString(`<div id="form-errors" hx-swap-oob="true" class="bg-cream border border-copper/30 rounded-[4px] p-4 mb-6">`)
	b.WriteString(`<p class="font-ui text-[11px] uppercase tracking-[0.35em] text-copper mb-2">Please fix the following</p>`)
	b.WriteString(`<ul class="font-body text-ink-mid text-sm space-y-1">`)
	for _, e := range errors {
		b.WriteString(`<li>`)
		b.WriteString(template.HTMLEscapeString(e))
		b.WriteString(`</li>`)
	}
	b.WriteString(`</ul></div>`)

	w.Write([]byte(b.String()))
}
