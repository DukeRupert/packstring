package data

// GalleryCategory represents a filterable image category.
type GalleryCategory struct {
	Name string // display label, e.g. "Fishing"
	Slug string // filter key, e.g. "fishing"
}

// GalleryImage represents a single gallery photo (or placeholder).
type GalleryImage struct {
	ID       int
	Src      string // base path without size suffix, e.g. "/static/img/gallery/gallery-01"
	Thumb    string // thumbnail path, e.g. "/static/img/gallery/gallery-01-thumb.webp"
	Alt      string
	Category string // matches GalleryCategory.Slug
}

// GalleryPageData holds all data rendered on the gallery page.
type GalleryPageData struct {
	Meta       PageMeta
	Categories []GalleryCategory
	Images     []GalleryImage
}

// GetGalleryPageData returns seed content for the gallery page.
func GetGalleryPageData() GalleryPageData {
	return GalleryPageData{
		Meta: PageMeta{
			Title:        "Photo Gallery — MT Hunt & Fish Outfitters",
			Description:  "Photos from guided fishing and hunting trips with MT Hunt & Fish Outfitters. Missouri River, Canyon Ferry, Elkhorn Mountains, and more.",
			CanonicalURL: SiteURL + "/gallery/",
			OGImage:      SiteURL + "/static/img/hero/hero-montana-1600w.webp",
		},
		Categories: []GalleryCategory{
			{Name: "Fishing", Slug: "fishing"},
			{Name: "Hunting", Slug: "hunting"},
			{Name: "Scenery", Slug: "scenery"},
			{Name: "Camp", Slug: "camp"},
		},
		Images: []GalleryImage{
			// Fishing (5)
			{ID: 1, Src: "/static/img/gallery/gallery-01", Thumb: "/static/img/gallery/gallery-01-thumb.webp", Alt: "Jet boat on the Missouri at sunrise", Category: "fishing"},
			{ID: 2, Src: "/static/img/gallery/gallery-02", Thumb: "/static/img/gallery/gallery-02-thumb.webp", Alt: "Drift boat below Holter Dam with mountains behind", Category: "fishing"},
			{ID: 3, Src: "/static/img/gallery/gallery-03", Thumb: "/static/img/gallery/gallery-03-thumb.webp", Alt: "Angler landing a brown trout on the Missouri", Category: "fishing"},
			{ID: 4, Src: "/static/img/gallery/gallery-04", Thumb: "/static/img/gallery/gallery-04-thumb.webp", Alt: "Canyon Ferry Lake at golden hour with calm water", Category: "fishing"},
			{ID: 5, Src: "/static/img/gallery/gallery-05", Thumb: "/static/img/gallery/gallery-05-thumb.webp", Alt: "Wade fishing in a side channel of the Missouri", Category: "fishing"},

			// Hunting (4)
			{ID: 6, Src: "/static/img/gallery/gallery-06", Thumb: "/static/img/gallery/gallery-06-thumb.webp", Alt: "Elk on a ridge in the Elkhorn Mountains at dawn", Category: "hunting"},
			{ID: 7, Src: "/static/img/gallery/gallery-07", Thumb: "/static/img/gallery/gallery-07-thumb.webp", Alt: "Mule deer buck in the Big Belt foothills", Category: "hunting"},
			{ID: 8, Src: "/static/img/gallery/gallery-08", Thumb: "/static/img/gallery/gallery-08-thumb.webp", Alt: "Hunter glassing a valley from a rocky overlook", Category: "hunting"},
			{ID: 9, Src: "/static/img/gallery/gallery-09", Thumb: "/static/img/gallery/gallery-09-thumb.webp", Alt: "Antelope on open prairie south of Helena", Category: "hunting"},

			// Scenery (5)
			{ID: 10, Src: "/static/img/gallery/gallery-10", Thumb: "/static/img/gallery/gallery-10-thumb.webp", Alt: "Missouri River canyon in fall color", Category: "scenery"},
			{ID: 11, Src: "/static/img/gallery/gallery-11", Thumb: "/static/img/gallery/gallery-11-thumb.webp", Alt: "Snow-capped peaks above the Gates of the Mountains", Category: "scenery"},
			{ID: 12, Src: "/static/img/gallery/gallery-12", Thumb: "/static/img/gallery/gallery-12-thumb.webp", Alt: "Sunrise over the Helena Valley looking west", Category: "scenery"},
			{ID: 13, Src: "/static/img/gallery/gallery-13", Thumb: "/static/img/gallery/gallery-13-thumb.webp", Alt: "Elkhorn Mountains with wildflower meadow in foreground", Category: "scenery"},
			{ID: 14, Src: "/static/img/gallery/gallery-14", Thumb: "/static/img/gallery/gallery-14-thumb.webp", Alt: "Storm clouds building over Canyon Ferry Lake", Category: "scenery"},

			// Camp (4)
			{ID: 15, Src: "/static/img/gallery/gallery-15", Thumb: "/static/img/gallery/gallery-15-thumb.webp", Alt: "Campfire on the riverbank after a full day on the water", Category: "camp"},
			{ID: 16, Src: "/static/img/gallery/gallery-16", Thumb: "/static/img/gallery/gallery-16-thumb.webp", Alt: "Camp kitchen setup with the Missouri in the background", Category: "camp"},
			{ID: 17, Src: "/static/img/gallery/gallery-17", Thumb: "/static/img/gallery/gallery-17-thumb.webp", Alt: "Gear laid out before a morning hunt", Category: "camp"},
			{ID: 18, Src: "/static/img/gallery/gallery-18", Thumb: "/static/img/gallery/gallery-18-thumb.webp", Alt: "Tailgate lunch with a view of the valley", Category: "camp"},
		},
	}
}
