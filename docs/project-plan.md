# Packstring — Project Plan

## Project Overview

**Project name:** Packstring
**Repo name:** `packstring`
**Domain:** `packstring.dev` (grab this)
**Demo URL:** `demo.fireflysoftware.dev/outfitter/` (until packstring.dev is set up)
**Stack:** Go (http server + html/template) · Tailwind CSS · Alpine.js (minimal) · htmx (contact form)
**Based on:** MT Hunt & Fish Outfitters (Forrest Fawthrop, Helena MT)

Packstring is Firefly Software's outfitter/guide website product — a repeatable,
high-quality website template system built specifically for Montana outfitters and guides.
The name comes from the line of pack horses and mules that carry everything an outfitter
needs into the backcountry. That's what this does for their online presence.

This is a static-ish site served by a small Go binary. No database needed for MVP.
The contact form posts to a Go handler that would send email in production;
for the demo it just returns a success message.

---

## Project Scaffolding

```
packstring/
├── cmd/
│   └── server/
│       └── main.go                 # HTTP server, route registration
├── internal/
│   └── handlers/
│       ├── pages.go                # Page handlers (home, trips, gallery, contact)
│       └── contact.go              # Contact form handler (POST /contact)
├── templates/
│   ├── layouts/
│   │   └── base.html               # Base layout (head, nav, footer, scripts)
│   ├── partials/
│   │   ├── nav.html                # Navigation bar
│   │   ├── footer.html             # Site footer
│   │   ├── trip-card.html          # Reusable trip card component
│   │   ├── testimonial.html        # Testimonial block component
│   │   ├── gallery-item.html       # Gallery photo item
│   │   └── contact-form.html       # Inquiry form partial
│   └── pages/
│       ├── home.html               # Homepage
│       ├── trips.html              # Trips overview hub
│       ├── fishing.html            # Fishing trips detail
│       ├── hunting.html            # Hunting trips detail
│       ├── packages.html           # Multi-day packages
│       ├── gallery.html            # Photo gallery
│       └── contact.html            # Contact / Book a Trip
├── static/
│   ├── css/
│   │   └── output.css              # Compiled Tailwind output
│   ├── js/
│   │   └── app.js                  # Alpine.js init + lightbox + gallery filter
│   ├── img/
│   │   ├── hero/                   # Hero images (WebP, 1920x1080 max)
│   │   ├── trips/                  # Trip section photos
│   │   ├── gallery/                # Gallery photos (full + thumbs)
│   │   └── icons/                  # Favicon, logo, social icons
│   └── fonts/                      # Self-hosted fonts if needed
├── tailwind.config.js
├── package.json                    # Tailwind build dependency
├── Makefile                        # Build/run/dev commands
├── Dockerfile                      # Production container
├── .air.toml                       # Hot reload config (development)
├── go.mod
├── go.sum
└── README.md
```

### Key Architecture Decisions

1. **Go `html/template` over templ** — For this demo, standard library templates are fine.
   The site is simple enough that templ's type safety isn't needed, and it keeps
   the demo approachable if you ever show the code to another developer or hire help.
   Easy to migrate to templ later if this becomes a repeatable product.

2. **No database** — All content is in template files and Go structs. Trip data lives
   in a `trips.go` data file as Go structs that get passed to templates.
   This makes the demo trivially deployable and the content easy to swap for a new client.

3. **Alpine.js for interactivity** — Gallery filtering, lightbox, mobile nav toggle.
   No build step for JS, just a CDN include or vendored file.

4. **htmx for the contact form** — POST to `/contact`, swap in a success message.
   No full page reload. Clean, minimal.

5. **Tailwind via CLI** — `npx tailwindcss` watches template files, outputs to
   `static/css/output.css`. No webpack, no Vite, no complexity.

---

## Feature List: MVP → Complete

### Milestone 1: MVP (The "Show Forrest" Version)

**Goal:** A deployed, visually polished site that demonstrates the value proposition.
Contact form works client-side (shows success state) but doesn't actually send email.
All content uses real trip info from Forrest's current site + placeholder/royalty-free photos.

**Estimated time: 15-20 hours**

#### 1.1 — Project Setup (2 hours)

- [ ] Initialize Go module, directory structure
- [ ] Set up Tailwind CSS with config and build script
- [ ] Create base layout template (head, nav shell, footer shell, script tags)
- [ ] Create Makefile with `dev`, `build`, `css-watch` targets
- [ ] Set up `.air.toml` for hot reload
- [ ] Verify: `make dev` serves a "Hello World" page at localhost:8080

**Git checkpoint: "scaffold: project structure and build pipeline"**

#### 1.2 — Navigation & Footer (2 hours)

- [x] Build responsive nav: logo, 5 links (Home, Trips, Gallery, Contact, Call Now)
- [x] Mobile hamburger menu (Alpine.js toggle)
- [x] "Call Now" button styled as accent CTA in nav (always visible, phone icon on mobile)
- [x] Footer: contact info, license numbers, quick links, copyright
- [x] Verify: nav works on mobile (375px), tablet (768px), desktop (1280px)

**Git checkpoint: "layout: responsive nav and footer"**

#### 1.3 — Homepage (4 hours)

- [x] Hero section: full-width background image, headline, subheadline, 2 CTA buttons
- [x] Quick intro: 2-3 sentence summary, license credentials
- [x] Trip cards section: 3 cards (Fishing, Hunting, Packages) with placeholder images
- [x] Testimonials section: 2-3 testimonials in a simple grid or stacked layout
- [x] Seasonal callout banner: "Now Booking: Summer 2026" with CTA
- [x] Verify: page looks polished on mobile and desktop, hero image loads fast

**Git checkpoint: "pages: homepage complete"**

#### 1.4 — Trips Hub + Fishing Page (3 hours)

- [x] Trips overview (`/trips/`): 3 cards linking to fishing, hunting, packages
- [x] Fishing page (`/trips/fishing/`): sectioned layout with:
  - Jet Boat section
  - Drift Boat section
  - Lake Trips section
  - Wade Trips section
  - Specialty section (pike, ice fishing, bass, salmon)
- [x] Each section: placeholder photo, description, includes, duration, price, "Book This Trip" CTA
- [x] Sticky or inline "Questions? Call Forrest" element
- [x] Verify: sections are scannable, CTAs link to /contact/?trip=xxx

**Git checkpoint: "pages: trips hub and fishing detail"**

#### 1.5 — Hunting + Packages Pages (2 hours)

- [x] Hunting page (`/trips/hunting/`): sectioned by game type (elk, deer, bear, antelope)
- [x] Each section: placeholder photo, season dates, pricing, what's included, CTA
- [x] Packages page (`/trips/packages/`): Triple Header and 6-Pack packages
- [x] Each package: full description, itinerary summary, pricing, CTA
- [x] Verify: consistent styling with fishing page

**Git checkpoint: "pages: hunting and packages"**

#### 1.6 — Contact Page + Form (3 hours)

- [x] Build inquiry form (name, email, phone, trip interest dropdown, dates, party size, experience level, message)
- [x] Trip interest dropdown pre-selects based on `?trip=` query param (Alpine.js reads URL)
- [x] htmx POST to `/contact` → returns success message partial (swaps form content)
- [x] Go handler: validates required fields, returns success HTML or error HTML
- [x] Below form: phone number (large), email, address, response time note
- [x] Honeypot field for spam prevention
- [x] Verify: form submits cleanly, validation works, success state displays

**Git checkpoint: "pages: contact form with htmx submission"**

#### 1.7 — MVP Polish & Deploy (2 hours)

- [x] Responsive pass: check every page at 375px, 768px, 1280px
- [x] Typography pass: consistent heading sizes, line heights, spacing
- [x] Color pass: establish palette (dark green / warm gold / cream — outdoorsy, premium feel)
- [x] Add favicon and Open Graph meta tags
- [x] Dockerfile for deployment
- [x] Deploy to demo.fireflysoftware.dev/outfitter/
- [x] Verify: site loads fast, looks professional, all links work

**Git checkpoint: "release: mvp deployed"**

---

### Milestone 2: Complete Demo (The "Pitch to Anyone" Version)

**Goal:** Add gallery, real image optimization, SEO, and the auto-responder stub.
This is the version you'd show at a MOGA meeting or send in a cold email.

**Estimated time: 10-15 hours on top of MVP**

#### 2.1 — Photo Gallery (4 hours)

- [ ] Gallery page (`/gallery/`): responsive masonry or CSS grid layout
- [ ] Category filter tabs: All | Fishing | Hunting | Scenery | Camp (Alpine.js)
- [ ] Lightbox on click (Alpine.js component or lightweight library like GLightbox)
- [ ] Lazy loading with `loading="lazy"` + Intersection Observer fallback
- [ ] Image optimization pipeline: script or Makefile target to convert to WebP, generate thumbs
- [ ] Verify: gallery loads fast even with 20+ images, filter is instant, lightbox works on mobile

**Git checkpoint: "feature: photo gallery with filtering and lightbox"**

#### 2.2 — Image Optimization Pass (2 hours)

- [ ] Source royalty-free Montana/outdoor photos for demo (Unsplash, Pexels)
- [ ] Process all images: WebP conversion, responsive srcset (400w, 800w, 1600w)
- [ ] Replace all placeholder images across all pages
- [ ] Hero images get special treatment: pre-load hint, optimized dimensions
- [ ] Verify: no image over 200KB, total homepage weight under 500KB

**Git checkpoint: "perf: optimized images across all pages"**

#### 2.3 — SEO & Structured Data (2 hours)

- [ ] Unique meta title and description for every page
- [ ] LocalBusiness structured data (JSON-LD) in base layout
- [ ] TouristTrip schema on fishing and hunting pages
- [ ] Open Graph tags (title, description, image) for each page
- [ ] Generate sitemap.xml from Go (or static file)
- [ ] robots.txt
- [ ] Canonical URLs
- [ ] Verify: Google Rich Results Test passes, Lighthouse SEO score 95+

**Git checkpoint: "seo: structured data and meta tags"**

#### 2.4 — Contact Form Auto-Responder Stub (2 hours)

- [ ] Expand contact handler: on successful submission, show what the auto-reply would say
- [ ] Display mock auto-reply email preview: "Here's what your client would receive..."
- [ ] Include gear list link and response time promise in the preview
- [ ] In production this would send via SMTP; for demo, just display the preview
- [ ] Verify: form submit → success + auto-reply preview renders cleanly

**Git checkpoint: "feature: auto-responder preview"**

#### 2.5 — Performance & Final QA (2 hours)

- [ ] Run Lighthouse audit: target 95+ on Performance, Accessibility, SEO, Best Practices
- [ ] Test on real mobile devices (not just devtools)
- [ ] Check all internal links
- [ ] Check print stylesheet (outfitters print things — gear lists, trip details)
- [ ] Add subtle page transitions or scroll animations if time allows (CSS only, no heavy JS)
- [ ] Final deploy to demo.fireflysoftware.dev/outfitter/
- [ ] Take before/after screenshots for portfolio

**Git checkpoint: "release: complete demo"**

---

### Milestone 3: Production-Ready Template (Future — After First Client)

These features turn the demo into a repeatable product for multiple outfitter clients.
Only build these once you have a paying client.

#### 3.1 — Email Integration

- [ ] SMTP integration for contact form (Postmark, SES, or self-hosted)
- [ ] Auto-responder sends real email to inquiring client
- [ ] Notification email to outfitter with all form data formatted cleanly

#### 3.2 — Stripe Deposit Collection

- [ ] "Pay Deposit" button on trip pages or as a step after form submission
- [ ] Stripe Checkout session created server-side
- [ ] Configurable deposit amount per trip type
- [ ] Success/cancel redirect pages
- [ ] Webhook handler for payment confirmation
- [ ] Email receipt to client and notification to outfitter

#### 3.3 — Seasonal Availability Calendar

- [ ] Simple data structure: trip type → date range → status (open / limited / booked)
- [ ] Visual calendar or date grid on trip pages
- [ ] Outfitter updates via a simple admin form or you update it for them (fits $35/mo model)

#### 3.4 — Blog / Fishing Report

- [ ] Markdown-based blog posts (no CMS, files in repo)
- [ ] RSS feed
- [ ] Latest 2-3 posts shown on homepage
- [ ] Good for SEO and showing the operation is active

#### 3.5 — Content Management Strategy

- [ ] Define how you onboard a new outfitter client:
  - 30-minute intake call → you extract all content
  - Client provides 10-20 photos
  - You create a content branch, client reviews on staging URL
  - Go live within 1-2 weeks
- [ ] Create a client intake form/questionnaire template
- [ ] Create a "swap the content" checklist for reusing the template

---

## Development Workflow

### Daily Process

1. Start a feature branch from `main`
2. Work on one checkbox group from the milestone
3. Test at mobile + desktop breakpoints
4. Commit with the checkpoint message listed above
5. Merge to `main` when the feature group is complete
6. Deploy `main` to demo URL after each milestone

### Makefile Targets

```makefile
dev:          # Start Go server with Air hot reload + Tailwind watch
build:        # Build Go binary + compile Tailwind for production
css-watch:    # Tailwind CLI in watch mode
css-build:    # Tailwind CLI single build (minified)
images:       # Run image optimization script (WebP conversion + resize)
deploy:       # Build Docker image and deploy
lighthouse:   # Run Lighthouse CI audit
```

### Color Palette (Starting Point)

```
Primary:    #1B4332  (deep forest green — trust, outdoors)
Secondary:  #B68D40  (warm gold — premium, Montana sunlight)
Accent:     #D4A574  (warm tan — leather, earth)
Background: #FEFAE0  (warm cream — clean but not sterile)
Text:       #1A1A1A  (near-black — easy to read)
White:      #FFFFFF  (cards, form backgrounds)
```

These are starting points — adjust once you see them with real photos.
The key insight: avoid the bright blue/white "tech startup" palette. This needs
to feel like leather and campfire, not Silicon Valley.

### Typography (Starting Point)

```
Headings:   Inter or similar clean sans-serif (not too techy)
Body:       System font stack (fast loading, familiar)
Accent:     Consider a subtle serif for testimonial quotes
```

---

## What to Do First

Literally the next 5 things to do, in order:

1. `mkdir packstring && cd packstring && go mod init github.com/firefly/packstring`
2. Set up the directory structure from the scaffolding section above
3. Get Tailwind CLI installed and generating output.css from your templates
4. Build the base layout + nav + footer (Milestone 1.1 + 1.2)
5. Build the homepage (Milestone 1.3)

After step 5, you have something you can look at and feel. Everything after that
is filling in pages and polishing.

---

## Timeline Summary

| Milestone | Hours | Calendar (part-time) | Deliverable |
|-----------|-------|---------------------|-------------|
| MVP | 15-20 | ~1 week | Deployed demo you can show Forrest |
| Complete | 10-15 | ~1 more week | Portfolio-ready demo for any outfitter |
| Production | Per client | After first sale | Real site with email, payments |

The MVP is the priority. Get it deployed, show it to Forrest over the fence,
and see if he bites. Everything else follows from that conversation.
