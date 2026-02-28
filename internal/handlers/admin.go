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
	"strconv"
	"strings"
	"sync"
	"time"

	"os"

	"github.com/firefly/packstring/internal/data"
	"github.com/firefly/packstring/internal/db"
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
		"formatCents": func(cents int) string {
			dollars := cents / 100
			remainder := cents % 100
			if remainder == 0 {
				return fmt.Sprintf("$%d", dollars)
			}
			return fmt.Sprintf("$%d.%02d", dollars, remainder)
		},
		"formatCents64": func(cents int64) string {
			dollars := cents / 100
			remainder := cents % 100
			if remainder == 0 {
				return fmt.Sprintf("$%d", dollars)
			}
			return fmt.Sprintf("$%d.%02d", dollars, remainder)
		},
		"timeAgo": func(t time.Time) string {
			d := time.Since(t)
			switch {
			case d < time.Minute:
				return "just now"
			case d < time.Hour:
				m := int(d.Minutes())
				if m == 1 {
					return "1 minute ago"
				}
				return fmt.Sprintf("%d minutes ago", m)
			case d < 24*time.Hour:
				h := int(d.Hours())
				if h == 1 {
					return "1 hour ago"
				}
				return fmt.Sprintf("%d hours ago", h)
			case d < 48*time.Hour:
				return "yesterday"
			default:
				return t.Format("Jan 2, 2006")
			}
		},
		"divide": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"statusLabel": func(s string) string {
			labels := map[string]string{
				"new":       "New",
				"contacted": "Contacted",
				"booked":    "Booked",
				"archived":  "Archived",
			}
			if l, ok := labels[s]; ok {
				return l
			}
			return s
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
	store        *db.Store
}

func NewAdmin(templates map[string]*template.Template, availability *data.AvailabilityStore, password string, store *db.Store) *Admin {
	return &Admin{
		templates:    templates,
		availability: availability,
		password:     password,
		sessions:     make(map[string]time.Time),
		store:        store,
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

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
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

// Dashboard renders the admin home page with stat cards and recent inquiries.
func (a *Admin) Dashboard(w http.ResponseWriter, r *http.Request) {
	newCount, _ := a.store.CountInquiries("new")
	totalCount, _ := a.store.CountInquiries("")
	bookedCount, _ := a.store.CountInquiries("booked")
	totalDeposits, _ := a.store.TotalDepositsCents()
	recent, _ := a.store.RecentInquiries(5)

	d := map[string]any{
		"Meta":          data.PageMeta{Title: "Admin Dashboard — MT Hunt & Fish Outfitters"},
		"NewCount":      newCount,
		"TotalCount":    totalCount,
		"BookedCount":   bookedCount,
		"TotalDeposits": totalDeposits,
		"RecentInquiries": recent,
		"ActiveNav":     "dashboard",
	}
	if err := a.templates["admin-dashboard"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering dashboard: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// InquiriesList renders the inquiry list page with optional status filtering.
func (a *Admin) InquiriesList(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	inquiries, err := a.store.ListInquiries(status)
	if err != nil {
		log.Printf("Error loading inquiries: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get counts for filter tabs
	allCount, _ := a.store.CountInquiries("")
	newCount, _ := a.store.CountInquiries("new")
	contactedCount, _ := a.store.CountInquiries("contacted")
	bookedCount, _ := a.store.CountInquiries("booked")
	archivedCount, _ := a.store.CountInquiries("archived")

	d := map[string]any{
		"Meta":           data.PageMeta{Title: "Inquiries — MT Hunt & Fish Outfitters"},
		"Inquiries":      inquiries,
		"CurrentStatus":  status,
		"AllCount":       allCount,
		"NewCount":       newCount,
		"ContactedCount": contactedCount,
		"BookedCount":    bookedCount,
		"ArchivedCount":  archivedCount,
		"ActiveNav":      "inquiries",
	}
	if err := a.templates["admin-inquiries"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering inquiries: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// InquiryDetail renders the detail page for a single inquiry.
func (a *Admin) InquiryDetail(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid inquiry ID", http.StatusBadRequest)
		return
	}

	inq, err := a.store.GetInquiry(id)
	if err != nil {
		log.Printf("Error loading inquiry %d: %v", id, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if inq == nil {
		http.Error(w, "Inquiry not found", http.StatusNotFound)
		return
	}

	payments, _ := a.store.GetPaymentsByInquiry(id)
	depositConfig, _ := a.store.GetDepositConfig(inq.TripSlug)

	d := map[string]any{
		"Meta":          data.PageMeta{Title: fmt.Sprintf("Inquiry #%d — MT Hunt & Fish Outfitters", id)},
		"Inquiry":       inq,
		"Payments":      payments,
		"DepositConfig": depositConfig,
		"ActiveNav":     "inquiries",
	}
	if err := a.templates["admin-inquiry-detail"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering inquiry detail: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// UpdateInquiryStatus handles status changes via htmx POST.
func (a *Admin) UpdateInquiryStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid inquiry ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newStatus := r.FormValue("status")
	if err := a.store.UpdateInquiryStatus(id, newStatus); err != nil {
		log.Printf("Error updating inquiry status: %v", err)
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	inq, _ := a.store.GetInquiry(id)
	if inq == nil {
		http.Error(w, "Inquiry not found", http.StatusNotFound)
		return
	}

	// Return the updated status section via htmx
	w.Header().Set("HX-Trigger", `{"showToast": "Status updated to `+newStatus+`"}`)
	a.renderInquiryStatus(w, inq)
}

// UpdateInquiryNotes handles notes updates via htmx POST.
func (a *Admin) UpdateInquiryNotes(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid inquiry ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	notes := r.FormValue("notes")
	if err := a.store.UpdateInquiryNotes(id, notes); err != nil {
		log.Printf("Error updating inquiry notes: %v", err)
		http.Error(w, "Failed to save notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", `{"showToast": "Notes saved"}`)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<button type="submit" class="btn btn-primary">Save Notes</button>`)
}

// renderInquiryStatus writes the inquiry-status partial HTML.
func (a *Admin) renderInquiryStatus(w http.ResponseWriter, inq *db.Inquiry) {
	if err := a.templates["admin-inquiry-detail"].ExecuteTemplate(w, "inquiry-status", inq); err != nil {
		log.Printf("Error rendering inquiry status: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// DepositsPage renders the deposit configuration page.
func (a *Admin) DepositsPage(w http.ResponseWriter, r *http.Request) {
	configs, _ := a.store.ListDepositConfigs()

	// Build a map for quick lookup
	configMap := make(map[string]db.DepositConfig)
	for _, c := range configs {
		configMap[c.TripSlug] = c
	}

	// Build trip list with current config
	type tripDeposit struct {
		Slug        string
		Name        string
		AmountCents int
		Enabled     bool
	}
	var trips []tripDeposit
	for _, t := range allTrips {
		td := tripDeposit{Slug: t.Slug, Name: t.Name}
		if c, ok := configMap[t.Slug]; ok {
			td.AmountCents = c.AmountCents
			td.Enabled = c.Enabled
		}
		trips = append(trips, td)
	}

	d := map[string]any{
		"Meta":      data.PageMeta{Title: "Deposit Settings — MT Hunt & Fish Outfitters"},
		"Trips":     trips,
		"ActiveNav": "deposits",
	}
	if err := a.templates["admin-deposits"].ExecuteTemplate(w, "base.html", d); err != nil {
		log.Printf("Error rendering deposits page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// SaveDeposits saves deposit config for all trips.
func (a *Admin) SaveDeposits(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	for _, t := range allTrips {
		amountStr := r.FormValue("amount_" + t.Slug)
		amount, _ := strconv.Atoi(amountStr)
		enabled := r.FormValue("enabled_"+t.Slug) == "on"

		dc := &db.DepositConfig{
			TripSlug:    t.Slug,
			TripName:    t.Name,
			AmountCents: amount * 100, // convert dollars to cents
			Enabled:     enabled,
		}
		if err := a.store.SaveDepositConfig(dc); err != nil {
			log.Printf("Error saving deposit config for %s: %v", t.Slug, err)
		}
	}

	w.Header().Set("HX-Trigger", `{"showToast": "Deposit settings saved"}`)
	// Redirect back (htmx will handle the toast)
	http.Redirect(w, r, "/admin/deposits/", http.StatusSeeOther)
}

// GenerateDepositLink creates a Stripe Checkout session for an inquiry's deposit.
func (a *Admin) GenerateDepositLink(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid inquiry ID", http.StatusBadRequest)
		return
	}

	inq, err := a.store.GetInquiry(id)
	if err != nil || inq == nil {
		http.Error(w, "Inquiry not found", http.StatusNotFound)
		return
	}

	depositConfig, _ := a.store.GetDepositConfig(inq.TripSlug)
	if depositConfig == nil || !depositConfig.Enabled || depositConfig.AmountCents == 0 {
		w.Header().Set("HX-Trigger", `{"showToast": "No deposit configured for this trip type"}`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	siteURL := os.Getenv("SITE_URL")
	if siteURL == "" {
		siteURL = "http://localhost:8080"
	}

	checkoutURL, sessionID, err := CreateCheckoutSession(
		inq.Email,
		depositConfig.AmountCents,
		inq.TripName,
		siteURL+"/payments/success",
		siteURL+"/payments/cancel",
		id,
	)
	if err != nil {
		log.Printf("Error creating Stripe session: %v", err)
		w.Header().Set("HX-Trigger", `{"showToast": "Failed to create payment link. Check Stripe configuration."}`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save payment record
	payment := &db.Payment{
		InquiryID:       id,
		StripeSessionID: sessionID,
		AmountCents:     depositConfig.AmountCents,
		Currency:        "usd",
		Status:          "pending",
		CustomerEmail:   inq.Email,
	}
	if _, err := a.store.CreatePayment(payment); err != nil {
		log.Printf("Error saving payment record: %v", err)
	}

	// Return the checkout URL for the admin to copy
	w.Header().Set("HX-Trigger", `{"showToast": "Deposit link generated"}`)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<div id="deposit-link-result" class="mt-4 p-4 bg-cream border border-copper/30 rounded-[4px]">
		<p class="font-ui text-[11px] uppercase tracking-[0.35em] text-ink-faded mb-2">Payment Link</p>
		<div class="flex items-center gap-2">
			<input type="text" value="%s" readonly
				class="flex-1 bg-white border border-sand-dk rounded-[4px] px-3 py-2 font-mono text-sm text-ink"
				id="deposit-url" onclick="this.select()">
			<button type="button" onclick="navigator.clipboard.writeText(document.getElementById('deposit-url').value); this.textContent='Copied!'; setTimeout(() => this.textContent='Copy', 2000)"
				class="btn btn-primary btn-sm whitespace-nowrap">Copy</button>
		</div>
		<p class="font-body text-ink-faded text-xs mt-2">Send this link to the client via email or text message.</p>
	</div>`, template.HTMLEscapeString(checkoutURL))
}

// --- Availability Editor (existing) ---

func (a *Admin) EditPage(w http.ResponseWriter, r *http.Request) {
	trips := a.availability.GetAll()
	groups := buildTripGroups(trips)

	d := map[string]any{
		"Meta":      data.PageMeta{Title: "Availability Editor — MT Hunt & Fish Outfitters"},
		"Groups":    groups,
		"Message":   "",
		"ActiveNav": "availability",
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
		w.Header().Set("HX-Trigger", `{"showToast": "Error saving availability"}`)
		a.renderResult(w, trips, "Error saving: "+err.Error())
		return
	}

	w.Header().Set("HX-Trigger", `{"showToast": "Availability saved"}`)
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
