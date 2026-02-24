package data

// GalleryCategory represents a filterable image category.
type GalleryCategory struct {
	Name string // display label, e.g. "Fishing"
	Slug string // filter key, e.g. "fishing"
}

// GalleryImage represents a single gallery photo (or placeholder).
type GalleryImage struct {
	ID       int
	Src      string // empty until real images are added
	Thumb    string // empty until real images are added
	Alt      string
	Category string // matches GalleryCategory.Slug
}

// GalleryPageData holds all data rendered on the gallery page.
type GalleryPageData struct {
	Categories []GalleryCategory
	Images     []GalleryImage
}

// GetGalleryPageData returns seed content for the gallery page.
func GetGalleryPageData() GalleryPageData {
	return GalleryPageData{
		Categories: []GalleryCategory{
			{Name: "Fishing", Slug: "fishing"},
			{Name: "Hunting", Slug: "hunting"},
			{Name: "Scenery", Slug: "scenery"},
			{Name: "Camp", Slug: "camp"},
		},
		Images: []GalleryImage{
			// Fishing (5)
			{ID: 1, Alt: "Jet boat on the Missouri at sunrise", Category: "fishing"},
			{ID: 2, Alt: "Drift boat below Holter Dam with mountains behind", Category: "fishing"},
			{ID: 3, Alt: "Angler landing a brown trout on the Missouri", Category: "fishing"},
			{ID: 4, Alt: "Canyon Ferry Lake at golden hour with calm water", Category: "fishing"},
			{ID: 5, Alt: "Wade fishing in a side channel of the Missouri", Category: "fishing"},

			// Hunting (4)
			{ID: 6, Alt: "Elk on a ridge in the Elkhorn Mountains at dawn", Category: "hunting"},
			{ID: 7, Alt: "Mule deer buck in the Big Belt foothills", Category: "hunting"},
			{ID: 8, Alt: "Hunter glassing a valley from a rocky overlook", Category: "hunting"},
			{ID: 9, Alt: "Antelope on open prairie south of Helena", Category: "hunting"},

			// Scenery (5)
			{ID: 10, Alt: "Missouri River canyon in fall color", Category: "scenery"},
			{ID: 11, Alt: "Snow-capped peaks above the Gates of the Mountains", Category: "scenery"},
			{ID: 12, Alt: "Sunrise over the Helena Valley looking west", Category: "scenery"},
			{ID: 13, Alt: "Elkhorn Mountains with wildflower meadow in foreground", Category: "scenery"},
			{ID: 14, Alt: "Storm clouds building over Canyon Ferry Lake", Category: "scenery"},

			// Camp (4)
			{ID: 15, Alt: "Campfire on the riverbank after a full day on the water", Category: "camp"},
			{ID: 16, Alt: "Camp kitchen setup with the Missouri in the background", Category: "camp"},
			{ID: 17, Alt: "Gear laid out before a morning hunt", Category: "camp"},
			{ID: 18, Alt: "Tailgate lunch with a view of the valley", Category: "camp"},
		},
	}
}
