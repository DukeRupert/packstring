package data

// TripsHubData holds data for the /trips/ hub page.
type TripsHubData struct {
	Meta      PageMeta
	TripCards []TripCard
}

// GetTripsHubData returns the trip cards displayed on the trips hub page.
func GetTripsHubData() TripsHubData {
	return TripsHubData{
		Meta: PageMeta{
			Title:        "Guided Trips — MT Hunt & Fish Outfitters",
			Description:  "Fishing and hunting trips from Helena, Montana. Jet boat, drift boat, wade, and lake fishing. Elk, deer, bear, and antelope hunts. Multi-day packages.",
			CanonicalURL: SiteURL + "/trips/",
			OGImage:      SiteURL + "/static/img/hero/hero-montana-1600w.webp",
		},
		TripCards: []TripCard{
			{
				Title:       "Fishing Trips",
				Description: "The Missouri below Holter Dam produces more trout per mile than any river in the Lower 48. Jet boat, drift boat, wade, and lake trips. Half-day and full-day.",
				Image:       "/static/img/trips/fishing-card",
				URL:         "/trips/fishing/",
			},
			{
				Title:       "Hunting Trips",
				Description: "Elk come through the Elkhorn timber like they have for centuries. Elk, deer, bear, and antelope hunts on private ranches and public land near Helena.",
				Image:       "/static/img/trips/hunting-card",
				URL:         "/trips/hunting/",
			},
			{
				Title:       "Packages",
				Description: "Three days. Three waters. Three chances to tell a story nobody back home will believe. The Triple Header and the 6-Pack.",
				Image:       "/static/img/trips/packages-card",
				URL:         "/trips/packages/",
			},
		},
	}
}
