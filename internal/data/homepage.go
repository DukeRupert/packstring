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
				Description: "Jet boat, drift boat, wade, and lake fishing on the Missouri River and nearby waters. Half-day and full-day options.",
				Image:       "/static/img/trips/fishing.webp",
				URL:         "/trips/fishing/",
			},
			{
				Title:       "Hunting Trips",
				Description: "Guided elk, deer, bear, and antelope hunts across private ranches and public land near Helena, Montana.",
				Image:       "/static/img/trips/hunting.webp",
				URL:         "/trips/hunting/",
			},
			{
				Title:       "Packages",
				Description: "Multi-day combo trips — the Triple Header and 6-Pack — for the complete Montana outdoor experience.",
				Image:       "/static/img/trips/packages.webp",
				URL:         "/trips/packages/",
			},
		},
		Testimonials: []Testimonial{
			{
				Quote:  "Forrest put us on fish all day long. Best guide experience we've had in 20 years of fishing out West.",
				Name:   "Mike R.",
				Detail: "Missouri River Fishing Trip",
			},
			{
				Quote:  "Professional, knowledgeable, and genuinely fun to spend a day with. We'll be back every fall.",
				Name:   "Dan & Sarah K.",
				Detail: "Elk Hunting Trip",
			},
			{
				Quote:  "The Triple Header package was the highlight of our year. Fishing, hunting, and scenery you can't beat anywhere else.",
				Name:   "Tom W.",
				Detail: "Triple Header Package",
			},
		},
	}
}
