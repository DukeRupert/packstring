package data

// SiteURL is the canonical base URL for the demo site.
const SiteURL = "https://mthuntfish.com"

// PageMeta holds SEO metadata rendered in the <head> of every page.
type PageMeta struct {
	Title        string // feeds <title> and og:title
	Description  string // feeds <meta description> and og:description
	CanonicalURL string // absolute URL
	OGImage      string // absolute URL to OG image
}
