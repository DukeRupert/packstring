package data

// TripSection represents a single trip or package offering rendered by the trip-section partial.
type TripSection struct {
	Title         string
	Slug          string   // used for ?trip= query param on /contact/
	Tagline       string
	Description   string
	Image         string   // base path without size suffix, e.g. "/static/img/trips/jet-boat"
	LocationLabel string   // "Waters", "Hunting Areas", "Destinations" — defaults to "Waters" in template
	Locations     []string
	Season        string // optional: e.g. "Sept 15 – Nov 25" for hunting
	Includes      []string
	Duration      string
	Price         string
	Availability  []DateSlot // populated at render time from availability.yaml
}

// FishingPageData holds all data rendered on the /trips/fishing/ page.
type FishingPageData struct {
	Meta  PageMeta
	Trips []TripSection
}

// GetFishingPageData returns seed content for the fishing trips page.
func GetFishingPageData() FishingPageData {
	return FishingPageData{
		Meta: PageMeta{
			Title:        "Fishing Trips — MT Hunt & Fish Outfitters",
			Description:  "Guided fishing trips on the Missouri River, Canyon Ferry, Fort Peck, and more. Jet boat, drift boat, wade, and lake trips from Helena, Montana.",
			CanonicalURL: SiteURL + "/trips/fishing/",
			OGImage:      SiteURL + "/static/img/trips/fishing-card-800w.webp",
		},
		Trips: []TripSection{
			{
				Title:   "Jet Boat Trips",
				Slug:    "jet-boat",
				Tagline: "Missouri River — Land of Giants",
				Image:   "/static/img/trips/jet-boat",
				Description: "The Missouri below Holter Dam runs cold and clear through canyon water most anglers never see from a road. " +
					"Forrest covers miles of it in a heated jet boat, putting people on rainbow and brown trout that run 18 to 24 inches year-round. " +
					"Multiple productive runs in a single day.",
				Locations: []string{"Missouri River (Craig to Cascade)"},
				Includes: []string{
					"All flies and terminal tackle",
					"Heated jet boat with casting platforms",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "$500/person",
			},
			{
				Title:   "Drift Boat Trips",
				Slug:    "drift-boat",
				Tagline: "Missouri. Big Horn. Blackfoot.",
				Image:   "/static/img/trips/drift-boat",
				Description: "A drift boat puts you in the seams where big trout hold. " +
					"Float the Missouri, the Big Horn, or the Blackfoot depending on season and conditions. " +
					"Quieter than a jet boat. Closer to the water. The way fly fishing was meant to be done.",
				Locations: []string{"Missouri River", "Big Horn River", "Blackfoot River"},
				Includes: []string{
					"All flies and terminal tackle",
					"Drift boat with comfortable seating",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "$500/person",
			},
			{
				Title:   "Lake Trips",
				Slug:    "lake",
				Tagline: "Canyon Ferry. Fort Peck. Holter.",
				Image:   "/static/img/trips/lake",
				Description: "Walleye, perch, and trout on Canyon Ferry, Fort Peck, and Holter. " +
					"Forrest trolls and jigs aboard a boat rigged with sonar and downriggers. " +
					"Good water for families. Kids catch fish here.",
				Locations: []string{"Canyon Ferry Reservoir", "Fort Peck Lake", "Holter Lake"},
				Includes: []string{
					"All tackle and bait",
					"Fully equipped fishing boat with electronics",
					"Lunch and drinks (full day)",
					"Fish cleaning and bagging",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "$450/person",
			},
			{
				Title:   "Wade Trips",
				Slug:    "wade",
				Tagline: "Boots in the Water, Rod in Hand",
				Image:   "/static/img/trips/wade",
				Description: "No boat, no motor. Just the river underfoot and wild trout in pocket water and riffles. " +
					"Forrest hikes into productive stretches of the Gallatin, the Shields, and spring creeks that don't show up on most maps.",
				Locations: []string{"Gallatin River", "Shields River", "Various spring creeks"},
				Includes: []string{
					"All flies and terminal tackle",
					"Waders and boots (if needed)",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "$400/person",
			},
			{
				Title:   "Specialty Trips",
				Slug:    "specialty",
				Tagline: "Beyond Trout",
				Image:   "/static/img/trips/specialty",
				Description: "Pike. Smallmouth bass. Chinook salmon. Lake trout. Winter ice fishing on Canyon Ferry. " +
					"Forrest runs these trips when the trout water gets crowded. " +
					"Different species, different methods, same guide who knows where they hold.",
				Locations: []string{"Missouri River", "Fort Peck Lake", "Canyon Ferry Reservoir", "Various rivers"},
				Includes: []string{
					"All tackle and bait",
					"Specialized equipment for target species",
					"Lunch and drinks (full day)",
					"Ice fishing gear and shelter (winter trips)",
				},
				Duration: "Full Day (8 hrs)",
				Price:    "$450/person",
			},
		},
	}
}
