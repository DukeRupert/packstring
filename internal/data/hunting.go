package data

// HuntingPageData holds all data rendered on the /trips/hunting/ page.
type HuntingPageData struct {
	Trips []TripSection
}

// GetHuntingPageData returns seed content for the hunting trips page.
func GetHuntingPageData() HuntingPageData {
	return HuntingPageData{
		Trips: []TripSection{
			{
				Title:         "Elk Hunts",
				Slug:          "elk-hunting",
				Tagline:       "Elkhorn Timber, Big Belt Country",
				Image:         "/static/img/trips/elk-hunting",
				Description: "Elk push through the Elkhorn and Big Belt mountains every fall on the same trails they've used for generations. " +
					"Forrest hunts them on a mix of private ranch land and public ground, glassing ridgelines at first light and working the timber as the day warms. " +
					"Five to seven days in steep country. Come prepared to hike.",
				LocationLabel: "Hunting Areas",
				Locations:     []string{"Elkhorn Mountains", "Big Belt Mountains"},
				Season:        "Sept 15 – Nov 25",
				Includes: []string{
					"Licensed guide service for the duration of the hunt",
					"Field dressing and caping",
					"Pack-out assistance (stock or ATV depending on terrain)",
					"Game care and cooling",
					"Camp setup and breakdown",
					"Spotting scopes and optics",
				},
				Duration: "5–7 Days",
				Price:    "Contact for pricing",
			},
			{
				Title:         "Deer Hunts",
				Slug:          "deer-hunting",
				Tagline:       "Ranch Land and River Breaks",
				Image:         "/static/img/trips/deer-hunting",
				Description: "Whitetail and mule deer on private ranches outside Helena and in the coulees along the Missouri River breaks. " +
					"Forrest scouts these properties through the summer, running trail cameras and tracking patterns before the season opens. " +
					"Spot-and-stalk or stand hunting depending on terrain and conditions. Three to five days.",
				LocationLabel: "Hunting Areas",
				Locations:     []string{"Helena area private ranches", "Missouri River breaks"},
				Season:        "Oct 20 – Nov 25",
				Includes: []string{
					"Licensed guide service for the duration of the hunt",
					"Field dressing and caping",
					"Game care and cooling",
					"Trail camera scouting data",
					"Stand or blind setup where applicable",
					"Transport to and from hunting areas",
				},
				Duration: "3–5 Days",
				Price:    "Contact for pricing",
			},
			{
				Title:         "Bear Hunts",
				Slug:          "bear-hunting",
				Tagline:       "Spring and Fall in the Elkhorns",
				Image:         "/static/img/trips/bear-hunting",
				Description: "Black bear in the Elkhorn Mountains, the Big Belts, and Helena National Forest. " +
					"Spring hunts run bait stations set weeks in advance. Fall hunts work spot-and-stalk through berry patches and creek bottoms. " +
					"Forrest knows the drainages where bears den and feed. Five to seven days. Two seasons to hunt them.",
				LocationLabel: "Hunting Areas",
				Locations:     []string{"Elkhorn Mountains", "Big Belt Mountains", "Helena National Forest"},
				Season:        "Apr 15 – May 31, Sept 15 – Nov 25",
				Includes: []string{
					"Licensed guide service for the duration of the hunt",
					"Bait station setup and maintenance (spring)",
					"Field dressing and skinning",
					"Game care and cooling",
					"Pack-out assistance",
					"Spotting scopes and optics",
				},
				Duration: "5–7 Days",
				Price:    "Contact for pricing",
			},
			{
				Title:         "Antelope Hunts",
				Slug:          "antelope-hunting",
				Tagline:       "Open Prairie, Long Glass",
				Image:         "/static/img/trips/antelope-hunting",
				Description: "Pronghorn on central Montana prairie and Broadwater County grassland. " +
					"Flat country where you can see for miles and so can they. Forrest runs spot-and-stalk and blind hunts over water sources. " +
					"Two to three days. The fastest game animal in North America does not give you many chances.",
				LocationLabel: "Hunting Areas",
				Locations:     []string{"Central Montana prairie", "Broadwater County"},
				Season:        "Sept 1 – Oct 15",
				Includes: []string{
					"Licensed guide service for the duration of the hunt",
					"Ground blind setup and placement",
					"Spotting and range estimation",
					"Field dressing",
					"Transport to and from hunting areas",
					"Game care and cooling",
				},
				Duration: "2–3 Days",
				Price:    "Contact for pricing",
			},
		},
	}
}
