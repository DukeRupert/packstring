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
				Description: "Cover miles of prime water on the Missouri River below Holter Dam in our heated jet boat. " +
					"This stretch is famous for producing trophy rainbow and brown trout year-round. " +
					"Perfect for anglers who want to hit multiple productive runs in a single day.",
				Waters: []string{"Missouri River (Craig to Cascade)"},
				Includes: []string{
					"All flies and terminal tackle",
					"Heated jet boat with casting platforms",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "From $500/person",
			},
			{
				Title:   "Drift Boat Trips",
				Slug:    "drift-boat",
				Tagline: "Classic Fly Fishing at Its Best",
				Description: "Float through scenic canyons and wade productive riffles from a traditional drift boat. " +
					"A quieter, more intimate experience that puts you right in the seams where big fish hold. " +
					"We fish the Missouri, Big Horn, and Blackfoot Rivers depending on season and conditions.",
				Waters: []string{"Missouri River", "Big Horn River", "Blackfoot River"},
				Includes: []string{
					"All flies and terminal tackle",
					"Drift boat with comfortable seating",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "From $500/person",
			},
			{
				Title:   "Lake Trips",
				Slug:    "lake",
				Tagline: "Big Water, Big Fish",
				Description: "Target walleye, perch, and trout on Montana's premier reservoirs. " +
					"We troll and jig aboard a fully equipped boat with electronics to find the fish. " +
					"Great for families and groups looking for a fun, productive day on the water.",
				Waters: []string{"Canyon Ferry Reservoir", "Fort Peck Lake", "Holter Lake"},
				Includes: []string{
					"All tackle and bait",
					"Fully equipped fishing boat with electronics",
					"Lunch and drinks (full day)",
					"Fish cleaning and bagging",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "From $450/person",
			},
			{
				Title:   "Wade Trips",
				Slug:    "wade",
				Tagline: "Boots in the Water, Rod in Hand",
				Description: "For the angler who wants to feel the river underfoot. " +
					"We hike into productive stretches of smaller rivers and creeks, targeting wild trout in pocket water and riffles. " +
					"An ideal choice for fly fishing purists and anyone who loves exploring on foot.",
				Waters: []string{"Gallatin River", "Shields River", "Various spring creeks"},
				Includes: []string{
					"All flies and terminal tackle",
					"Waders and boots (if needed)",
					"Streamside lunch (full day)",
					"Drinks and snacks",
				},
				Duration: "Full Day (8 hrs) or Half Day (4 hrs)",
				Price:    "From $400/person",
			},
			{
				Title:   "Specialty Trips",
				Slug:    "specialty",
				Tagline: "Beyond Trout — Something Different",
				Description: "Looking for something off the beaten path? We offer guided trips for pike, smallmouth bass, " +
					"Chinook salmon, lake trout, and winter ice fishing. " +
					"These trips are tailored to adventurous anglers who want a unique Montana experience.",
				Waters: []string{"Missouri River", "Fort Peck Lake", "Canyon Ferry Reservoir", "Various rivers"},
				Includes: []string{
					"All tackle and bait",
					"Specialized equipment for target species",
					"Lunch and drinks (full day)",
					"Ice fishing gear and shelter (winter trips)",
				},
				Duration: "Full Day (8 hrs)",
				Price:    "From $450/person",
			},
		},
	}
}
