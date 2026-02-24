package data

// PackagesPageData holds all data rendered on the /trips/packages/ page.
type PackagesPageData struct {
	Packages []TripSection
}

// GetPackagesPageData returns seed content for the multi-day packages page.
func GetPackagesPageData() PackagesPageData {
	return PackagesPageData{
		Packages: []TripSection{
			{
				Title:   "Montana Triple Header",
				Slug:    "triple-header",
				Tagline: "Fish. Hunt. Do It All in Five Days.",
				Description: "Two days on the Missouri chasing trout. One day in the Elkhorns after elk or deer. " +
					"Two flex days to fish Canyon Ferry, wade a spring creek, or just sit on the porch and do nothing. " +
					"Forrest builds the itinerary around the season, the conditions, and what you came to do. " +
					"Five days, four nights. Lodging, meals on guided days, and all gear included.",
				LocationLabel: "Destinations",
				Locations:     []string{"Missouri River", "Canyon Ferry", "Elkhorn Mountains"},
				Includes: []string{
					"Lodging coordination (4 nights)",
					"All meals on guided days",
					"All fishing tackle and gear",
					"All hunting gear (rifle/bow not included)",
					"Game processing coordination",
					"Airport pickup from Helena Regional",
					"Custom itinerary planning",
				},
				Duration: "5 Days / 4 Nights",
				Price:    "$3,500/person",
			},
			{
				Title:   "Montana 6-Pack",
				Slug:    "six-pack",
				Tagline: "Seven Days. Five Waters. One State That Has It All.",
				Description: "Three days fishing the Missouri, Fort Peck, and Canyon Ferry. " +
					"Two days hunting the Elkhorns and Big Belts. " +
					"Two days to pick your own — a second run at the river, a wade trip on the Gallatin, or a morning in a duck blind. " +
					"Forrest handles the logistics. Lodging, meals, gear, transport, game processing. All of it. " +
					"Seven days, six nights. The full Montana trip.",
				LocationLabel: "Destinations",
				Locations:     []string{"Missouri River", "Fort Peck", "Canyon Ferry", "Elkhorn Mountains", "Big Belt Mountains"},
				Includes: []string{
					"Lodging coordination (6 nights)",
					"All meals on guided days",
					"All fishing tackle and gear",
					"All hunting gear (rifle/bow not included)",
					"Game processing coordination",
					"Airport pickup from Helena Regional",
					"Custom itinerary planning",
					"Flex day activity options",
				},
				Duration: "7 Days / 6 Nights",
				Price:    "$5,500/person",
			},
		},
	}
}
