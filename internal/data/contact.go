package data

// ContactPageData holds data rendered on the /contact/ page.
type ContactPageData struct {
	Meta PageMeta
}

// GetContactPageData returns metadata for the contact page.
func GetContactPageData() ContactPageData {
	return ContactPageData{
		Meta: PageMeta{
			Title:        "Book a Trip — MT Hunt & Fish Outfitters",
			Description:  "Contact Forrest Fawthrop to book a guided fishing or hunting trip out of Helena, Montana. Call (406) 459-5352 or send an inquiry.",
			CanonicalURL: SiteURL + "/contact/",
			OGImage:      SiteURL + "/static/img/hero/hero-montana-1600w.webp",
		},
	}
}

// ContactSuccessData holds form submission data displayed in the auto-reply preview.
type ContactSuccessData struct {
	Name      string
	Email     string
	Trip      string // human-readable display name
	Dates     string
	PartySize string
}

// TripDisplayName maps a trip slug to a human-readable name.
func TripDisplayName(slug string) string {
	names := map[string]string{
		"jet-boat":          "Jet Boat Trips",
		"drift-boat":        "Drift Boat Trips",
		"lake":              "Lake Trips",
		"wade":              "Wade Trips",
		"specialty":         "Specialty Trips",
		"elk-hunting":       "Elk Hunts",
		"deer-hunting":      "Deer Hunts",
		"bear-hunting":      "Bear Hunts",
		"antelope-hunting":  "Antelope Hunts",
		"triple-header":     "Montana Triple Header",
		"six-pack":          "Montana 6-Pack",
	}
	if name, ok := names[slug]; ok {
		return name
	}
	return slug
}
