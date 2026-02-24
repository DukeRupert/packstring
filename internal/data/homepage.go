package data

// TripCard represents a summary card for a trip category.
type TripCard struct {
	Title       string
	Description string
	Image       string // path under /static/img/trips/
	URL         string
}

// Testimonial represents a client review quote.
type Testimonial struct {
	Quote  string
	Name   string
	Detail string // e.g. "Helena, MT" or trip type
}

// HomePageData holds all data rendered on the homepage.
type HomePageData struct {
	TripCards    []TripCard
	Testimonials []Testimonial
}

// GetHomePageData returns seed content for the homepage.
func GetHomePageData() HomePageData {
	return HomePageData{
		TripCards: []TripCard{
			{
				Title:       "Fishing Trips",
				Description: "The Missouri below Holter Dam produces more trout per mile than any river in the Lower 48. Jet boat, drift boat, wade, and lake trips. Half-day and full-day.",
				Image:       "/static/img/trips/fishing.webp",
				URL:         "/trips/fishing/",
			},
			{
				Title:       "Hunting Trips",
				Description: "Elk come through the Elkhorn timber like they have for centuries. Guided hunts for elk, deer, bear, and antelope on private ranches and public land near Helena.",
				Image:       "/static/img/trips/hunting.webp",
				URL:         "/trips/hunting/",
			},
			{
				Title:       "Packages",
				Description: "Three days. Three waters. Three chances to tell a story nobody back home will believe. The Triple Header and the 6-Pack.",
				Image:       "/static/img/trips/packages.webp",
				URL:         "/trips/packages/",
			},
		},
		Testimonials: []Testimonial{
			{
				Quote:  "Forrest put us on fish all day long. Best guide experience we've had in 20 years of fishing out West.",
				Name:   "Mike R.",
				Detail: "Missouri River Jet Boat, June 2025",
			},
			{
				Quote:  "Professional, knowledgeable, and genuinely fun to spend a day with. We'll be back every fall.",
				Name:   "Dan & Sarah K.",
				Detail: "Elk Hunt, October 2024",
			},
			{
				Quote:  "The Triple Header package was the highlight of our year. Fishing, hunting, and scenery you can't beat anywhere else.",
				Name:   "Tom W.",
				Detail: "Triple Header Package, August 2025",
			},
		},
	}
}
