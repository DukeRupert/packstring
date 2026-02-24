package data

// TripsHubData holds data for the /trips/ hub page.
type TripsHubData struct {
	TripCards []TripCard
}

// GetTripsHubData returns the trip cards displayed on the trips hub page.
func GetTripsHubData() TripsHubData {
	return TripsHubData{
		TripCards: []TripCard{
			{
				Title:       "Fishing Trips",
				Description: "From jet boat runs on the Missouri to quiet wade sessions on the Gallatin — half-day and full-day guided fishing across Montana's best waters.",
				Image:       "/static/img/trips/fishing.webp",
				URL:         "/trips/fishing/",
			},
			{
				Title:       "Hunting Trips",
				Description: "Elk, deer, bear, and antelope hunts on private ranches and public land near Helena. Fall and winter seasons available.",
				Image:       "/static/img/trips/hunting.webp",
				URL:         "/trips/hunting/",
			},
			{
				Title:       "Packages",
				Description: "Multi-day combo trips that pair fishing, hunting, and backcountry access into one unforgettable Montana experience.",
				Image:       "/static/img/trips/packages.webp",
				URL:         "/trips/packages/",
			},
		},
	}
}
