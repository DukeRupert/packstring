package data

// FishingTrip represents a single fishing trip offering.
type FishingTrip struct {
	Title       string
	Slug        string   // used for ?trip= query param on /contact/
	Tagline     string
	Description string
	Waters      []string
	Includes    []string
	Duration    string
	Price       string
}

// FishingPageData holds all data rendered on the /trips/fishing/ page.
type FishingPageData struct {
	Trips []FishingTrip
}

// GetFishingPageData returns seed content for the fishing trips page.
func GetFishingPageData() FishingPageData {
	return FishingPageData{
		Trips: []FishingTrip{
			{
				Title:   "Jet Boat Trips",
				Slug:    "jet-boat",
				Tagline: "Missouri River — Land of Giants",
				Description: "The Missouri below Holter Dam runs cold and clear through canyon water most anglers never see from a road. " +
					"Forrest covers miles of it in a heated jet boat, putting people on rainbow and brown trout that run 18 to 24 inches year-round. " +
					"Multiple productive runs in a single day.",
				Waters: []string{"Missouri River (Craig to Cascade)"},
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
				Description: "A drift boat puts you in the seams where big trout hold. " +
					"Float the Missouri, the Big Horn, or the Blackfoot depending on season and conditions. " +
					"Quieter than a jet boat. Closer to the water. The way fly fishing was meant to be done.",
				Waters: []string{"Missouri River", "Big Horn River", "Blackfoot River"},
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
				Description: "Walleye, perch, and trout on Canyon Ferry, Fort Peck, and Holter. " +
					"Forrest trolls and jigs aboard a boat rigged with sonar and downriggers. " +
					"Good water for families. Kids catch fish here.",
				Waters: []string{"Canyon Ferry Reservoir", "Fort Peck Lake", "Holter Lake"},
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
				Description: "No boat, no motor. Just the river underfoot and wild trout in pocket water and riffles. " +
					"Forrest hikes into productive stretches of the Gallatin, the Shields, and spring creeks that don't show up on most maps.",
				Waters: []string{"Gallatin River", "Shields River", "Various spring creeks"},
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
				Description: "Pike. Smallmouth bass. Chinook salmon. Lake trout. Winter ice fishing on Canyon Ferry. " +
					"Forrest runs these trips when the trout water gets crowded. " +
					"Different species, different methods, same guide who knows where they hold.",
				Waters: []string{"Missouri River", "Fort Peck Lake", "Canyon Ferry Reservoir", "Various rivers"},
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
