package handlers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/firefly/packstring/internal/data"
)

// AdminFuncMap returns template functions needed by admin templates.
func AdminFuncMap() template.FuncMap {
	return template.FuncMap{
		"jsonSlots": func(slots []data.DateSlot) template.JS {
			type jsSlot struct {
				Dates  string `json:"dates"`
				Status string `json:"status"`
				Note   string `json:"note"`
			}
			out := make([]jsSlot, len(slots))
			for i, s := range slots {
				out[i] = jsSlot{Dates: s.Dates, Status: s.Status, Note: s.Note}
			}
			b, _ := json.Marshal(out)
			return template.JS(b)
		},
	}
}

// tripMeta holds display info for each trip slug used in the admin editor.
type tripMeta struct {
	Slug     string
	Name     string
	Category string // "Fishing", "Hunting", "Packages"
}

// allTrips defines the canonical order and display names for the admin editor.
var allTrips = []tripMeta{
	{"jet-boat", "Jet Boat Trips", "Fishing"},
	{"drift-boat", "Drift Boat Trips", "Fishing"},
	{"lake", "Lake Trips", "Fishing"},
	{"wade", "Wade Trips", "Fishing"},
	{"specialty", "Specialty Trips", "Fishing"},
	{"elk-hunting", "Elk Hunts", "Hunting"},
	{"deer-hunting", "Deer Hunts", "Hunting"},
	{"bear-hunting", "Bear Hunts", "Hunting"},
	{"antelope-hunting", "Antelope Hunts", "Hunting"},
	{"triple-header", "Montana Triple Header", "Packages"},
	{"six-pack", "Montana 6-Pack", "Packages"},
}

// adminTripGroup is passed to the template for rendering grouped trips.
type adminTripGroup struct {
	Category string
	Trips    []adminTrip
}

type adminTrip struct {
	Slug  string
	Name  string
	Slots []data.DateSlot
}

type Admin struct {
	templates    map[string]*template.Template
	availability *data.AvailabilityStore
	password     string
	sessions     map[string]time.Time // token → expiry
	mu           sync.Mutex
}

func NewAdmin(templates map[string]*template.Template, availability *data.AvailabilityStore, password string) *Admin {
	return &Admin{
		templates:    templates,
		availability: availability,
		password:     password,
		sessions:     make(map[string]time.Time),
	}
}

// RequireAuth wraps a handler, redirecting to login if the session is invalid.
func (a *Admin) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("admin_session")
		if err != nil {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		a.mu.Lock()
		expiry, ok := a.sessions[cookie.Value]
		if ok && time.Now().After(expiry) {
			delete(a.sessions, cookie.Value)
			ok = false
		}
		a.mu.Unlock()

		if !ok {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func (a *Admin) LoginPage(w http.ResponseWriter, r *http.Request) {
	d := map[string]any{
		"Meta":  data.PageMeta{Title: "Admin Login — MT Hunt & Fish Outfitters"},
		"Error": "",
	}
	if err := a.templates["admin-login"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering admin login: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *Admin) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")
	if subtle.ConstantTimeCompare([]byte(password), []byte(a.password)) != 1 {
		d := map[string]any{
			"Meta":  data.PageMeta{Title: "Admin Login — MT Hunt & Fish Outfitters"},
			"Error": "Wrong password. Try again.",
		}
		w.WriteHeader(http.StatusUnauthorized)
		if err := a.templates["admin-login"].ExecuteTemplate(w, "base.html", d); err != nil {
			log.Printf("Error rendering admin login: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Generate session token
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	token := hex.EncodeToString(tokenBytes)

	expiry := time.Now().Add(24 * time.Hour)
	a.mu.Lock()
	a.sessions[token] = expiry
	a.mu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "admin_session",
		Value:    token,
		Path:     "/admin/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expiry,
	})

	http.Redirect(w, r, "/admin/availability", http.StatusSeeOther)
}

func (a *Admin) Logout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("admin_session"); err == nil {
		a.mu.Lock()
		delete(a.sessions, cookie.Value)
		a.mu.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "admin_session",
		Value:    "",
		Path:     "/admin/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func (a *Admin) EditPage(w http.ResponseWriter, r *http.Request) {
	trips := a.availability.GetAll()
	groups := buildTripGroups(trips)

	d := map[string]any{
		"Meta":    data.PageMeta{Title: "Availability Editor — MT Hunt & Fish Outfitters"},
		"Groups":  groups,
		"Message": "",
	}
	if err := a.templates["admin"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering admin editor: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *Admin) SaveAvailability(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	trips := make(map[string][]data.DateSlot)

	for _, tm := range allTrips {
		slug := tm.Slug
		// Find how many slots were submitted for this trip
		// Form fields: slots[slug][0][dates], slots[slug][0][status], slots[slug][0][note]
		var slots []data.DateSlot
		for i := 0; ; i++ {
			prefix := fmt.Sprintf("slots[%s][%d]", slug, i)
			dates := strings.TrimSpace(r.FormValue(prefix + "[dates]"))
			status := strings.TrimSpace(r.FormValue(prefix + "[status]"))

			if dates == "" && status == "" {
				break
			}
			if dates == "" {
				continue // skip empty rows
			}
			if status == "" {
				status = "open"
			}

			note := strings.TrimSpace(r.FormValue(prefix + "[note]"))
			slots = append(slots, data.DateSlot{
				Dates:  dates,
				Status: status,
				Note:   note,
			})
		}
		if len(slots) > 0 {
			trips[slug] = slots
		}
	}

	if err := a.availability.Save(trips); err != nil {
		log.Printf("[admin] save error: %v", err)
		a.renderResult(w, trips, "Error saving: "+err.Error())
		return
	}

	a.renderResult(w, trips, "Availability saved successfully.")
}

func (a *Admin) renderResult(w http.ResponseWriter, trips map[string][]data.DateSlot, message string) {
	groups := buildTripGroups(trips)
	d := map[string]any{
		"Groups":  groups,
		"Message": message,
	}
	if err := a.templates["admin"].ExecuteTemplate(w, "admin-form", d); err != nil {
		log.Printf("Error rendering admin form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func buildTripGroups(trips map[string][]data.DateSlot) []adminTripGroup {
	groupMap := make(map[string][]adminTrip)
	var categories []string
	seen := make(map[string]bool)

	for _, tm := range allTrips {
		if !seen[tm.Category] {
			seen[tm.Category] = true
			categories = append(categories, tm.Category)
		}
		slots := trips[tm.Slug]
		if slots == nil {
			slots = []data.DateSlot{}
		}
		groupMap[tm.Category] = append(groupMap[tm.Category], adminTrip{
			Slug:  tm.Slug,
			Name:  tm.Name,
			Slots: slots,
		})
	}

	// Sort categories in the order they first appear (Fishing, Hunting, Packages)
	var groups []adminTripGroup
	for _, cat := range categories {
		groups = append(groups, adminTripGroup{
			Category: cat,
			Trips:    groupMap[cat],
		})
	}
	// Ensure stable order
	sort.SliceStable(groups, func(i, j int) bool {
		order := map[string]int{"Fishing": 0, "Hunting": 1, "Packages": 2}
		return order[groups[i].Category] < order[groups[j].Category]
	})
	return groups
}
