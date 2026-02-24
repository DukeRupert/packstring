# Packstring Copywriter Agent Memory

## Client Details: MT Hunt & Fish Outfitters
- Guide: Forrest Fawthrop, Helena MT
- Montana Outfitter License #20194
- USCG Captain License #3427176
- Phone: (406) 459-5352
- Address: 610 Jeanne Rd, Helena, MT 59602
- Waters: Missouri River, Canyon Ferry, Fort Peck, Holter, Gallatin, Shields, Big Horn, Blackfoot
- Hunting areas: Elkhorn Mountains, Big Belt Mountains
- Trips: Jet boat, drift boat, wade, lake, specialty (pike/smallmouth/salmon/ice), hunting (elk/deer/bear/antelope)
- Packages: Triple Header, 6-Pack

## Key Voice Findings (Feb 2026 Audit)
- See [audit-findings.md](audit-findings.md) for full details
- Most pervasive issue: first-person plural ("we"/"our"/"us") used throughout -- must be third-person narrator or Forrest's name
- Banned words found: "world-class" (trips hub), "unforgettable" (packages card), "experience" (multiple)
- Trip cards missing Showman's Rule hooks -- all three homepage cards are factual but flat
- "From $X/person" pricing pattern = same as banned "starting at" -- state prices plainly
- Question headings used in CTA sections -- voice guide says stay declarative
- Testimonial attribution missing city/state and date

## Copy Data Locations
- Homepage trip cards + testimonials: `/home/dukerupert/Repos/packstring/internal/data/homepage.go`
- Trips hub card descriptions: `/home/dukerupert/Repos/packstring/internal/data/tripshub.go`
- Fishing trip descriptions, taglines, prices: `/home/dukerupert/Repos/packstring/internal/data/trips.go`
- Hunting trip data: needs new file (e.g., `hunting.go`) -- copy drafted Feb 2026
- Packages data: needs new file (e.g., `packages.go`) -- copy drafted Feb 2026
- Template copy (headings, CTAs, hero text): in `templates/pages/*.html` and `templates/partials/*.html`
- Brand voice guide: `/home/dukerupert/Repos/packstring/docs/packstring-brand-voice-guide.md`

## Approved Copy Patterns
- Hero subhead pattern: list places (period-separated), state tenure, end with short declarative
  - Fishing: "The Missouri. Canyon Ferry. Fort Peck. The Gallatin. Forrest has fished them all for 25 years. He knows where they hold."
  - Hunting: "The Elkhorns. The Big Belts. The Missouri breaks. Forrest has guided these drainages for over 25 years. He knows where they move."
  - Packages: "More days on the water and in the field. One guide who knows all of it. Forrest builds these trips around you."
- Bottom CTA pattern: differentiate by reader mental state (fishing=logistics, hunting=planning, packages=customization)
- Hunting trips use "Contact for pricing" (not "From $X" or "Starting at")
- Package prices stated plainly: "$3,500/person", "$5,500/person"

## Struct/Template Architecture Notes
- trip-section.html partial hardcodes "Waters" label -- needs LocationLabel field for hunting ("Hunting Areas") and packages ("Destinations")
- Option: generic Trip struct with LocationLabel string field to reuse partial across all trip types
