# Firefly Software — Packstring: Outfitter & Guide Service Offering

## Part 1: Service Page Brainstorm

### Page Concept: `/outfitters/` or `/guides/`

**Headline:** "Your clients pay $500 a day. Your website should look like it."

**Opening hook:** Montana outfitters and guides run premium experiences — $400 fishing trips, $5,000 elk hunts, $3,500 multi-day packages. But most of them are booking through websites that haven't been updated since flip phones. The coastal families and out-of-state sportsmen paying these rates expect a professional online presence. When they're comparing three guides, they book with the one whose site looks trustworthy and makes it easy.

**Core pitch (3 value props):**

1. **Look like a $5,000 experience** — Your website is the first impression for out-of-state clients who've never set foot in Montana. A modern, fast site with great photos and clear trip info converts browsers into booked clients.

2. **Stop losing bookings to phone tag** — An online inquiry form with trip selection, party size, and preferred dates means you wake up to real leads instead of voicemails you can't return from the river.

3. **Own your business, not a platform** — No Airbnb commissions, no FishingBooker fees, no WordPress plugin nightmares. Your site, your data, your bookings.

**What's included (the standard outfitter package):**

- Custom-built website (not a template)
- Trip/service pages with pricing and photos
- Online trip inquiry form (trip type, dates, party size, message)
- Photo gallery optimized for fast loading
- Mobile-first design (tourists browse on phones)
- Google Business Profile optimization
- Hosting, SSL, DNS included
- Unlimited content updates
- Direct developer access

**Optional add-ons (Phase 2 upsells):**

- Stripe deposit collection (take $500 deposits online)
- Seasonal availability calendar (simple open/booked/limited display, NOT a full booking engine)
- Email capture + basic automated follow-up ("Thanks for your inquiry, here's our gear list")
- Fishing/hunting report blog (guide can text or email updates, developer posts them)
- Google Reviews integration

**Pricing suggestion:**

- Same as standard Firefly pricing: $500 build + $35/mo
- Stripe deposit integration add-on: +$250 build, no additional monthly
- Availability calendar: +$250 build, no additional monthly

**Social proof section:**

- Before/after screenshots from a real outfitter site rebuild (this is what the demo is for)
- "One extra booking per season pays for your website for the entire year"

**CTA:** "Let's talk about your operation. No pressure, no sales pitch — just an honest conversation about whether a new site would bring you more bookings."

**Distribution/outreach ideas:**

- MOGA (500+ members) — inquire about business membership or newsletter sponsorship
- FOAM (1,000+ members) — similar approach
- Montana Governor's Conference on Tourism — attend or sponsor
- Direct outreach to outfitters with bad websites (you can literally show them their site vs. the demo)
- Fishing/hunting forums and Facebook groups

---

## Part 2: Pilot Site Audit — Montana Hunt & Fish Outfitters

### About the Business

- **Owner:** Forrest Fawthrop
- **Location:** 610 Jeanne Rd, Helena, MT 59602
- **License:** Outfitter #20194, US Coast Guard Captain #3427176
- **Phone:** (406) 459-5352
- **Current site:** mthuntandfishoutfitters.com
- **Current platform:** WordPress (built by "Bozeman Interactive")
- **Services:** Guided fishing (walleye, pike, bass, lake trout, chinook salmon, rainbow/brown trout) and hunting (elk, deer, bear, antelope)
- **Trip types:** Jet boat trips, drift boat trips, wade trips, ice fishing, lake fishing, multi-day packages
- **Key waters:** Canyon Ferry Lake, Fort Peck Reservoir, Holter Lake, Missouri River, Big Horn River, Blackfoot River, Gallatin River, Shields River
- **Packages:** "Montana Triple Header" and "Montana 6-Pack" multi-day packages

### Current Site Problems

**Navigation disaster:**

- 30+ menu items across deeply nested dropdowns
- Duplicate entries ("Trips" and "Trips we offer" are separate menus)
- Gallery sub-menu has 17 items
- "Trips we offer" has entries 4 levels deep
- Mobile menu is overwhelming and unusable
- A tourist looking for "how much does a fishing trip cost" would have to click through 3-4 levels

**Homepage issues:**

- Hero is a Vimeo embed that auto-loads (slow, bandwidth-heavy)
- Wall of text with no visual hierarchy
- Photos are raw JPGs, not optimized (slow loading)
- "COMING SOON!" sections still present
- No clear CTA — no "Book a Trip" or "Check Availability" button
- No pricing visible on homepage
- Logo shows "Walleye Hunter Outfitters" while the site is "Montana Hunting & Fishing Outfitters" (brand confusion)

**Missing critical features:**

- No online inquiry form with trip selection
- No pricing on the rates page (or if present, buried in the nav chaos)
- No seasonal availability information
- No Google reviews integration
- No structured data for local SEO
- No mobile optimization (WordPress theme is barely responsive)
- No clear "what to expect" or trip preparation info on trip pages

**What IS working (preserve these):**

- Forrest has good photo content (fish, scenery, boats, camps)
- The business offers genuine variety (fishing AND hunting, multiple water types)
- He has a "Fishing Report" blog section (good for SEO, shows he's active)
- Testimonials page exists
- FAQ page exists
- Gear list page exists (practical, useful content)

---

## Part 3: Demo Site Spec Sheet

### Project: Packstring Demo — MT Hunt & Fish Outfitters Rebuild

**Purpose:** Create a functional demo site that demonstrates the value of a modern outfitter website. This will be used as:

1. A sales tool to show Forrest what his site *could* look like
2. A portfolio piece for pitching other outfitters
3. A template foundation for future outfitter builds

**Note:** This is a demo/pilot using publicly available information. It will NOT go live without Forrest's involvement and approval.

---

### Site Architecture (7 pages)

```
/                       → Homepage
/trips/                 → All trips overview
/trips/fishing/         → Fishing trips (all water types)
/trips/hunting/         → Hunting trips (all game types)
/trips/packages/        → Multi-day packages
/gallery/               → Photo gallery
/contact/               → Inquiry form + contact info
```

Compared to the current 30+ page sprawl, this structure lets a visitor find anything in 2 clicks max. Individual trip details (Canyon Ferry, Fort Peck, etc.) are sections within the fishing page, not separate pages buried in sub-menus.

---

### Page Specs

#### Homepage (`/`)

**Hero section:**

- Full-width photo (Montana river/mountain scene — use a placeholder or royalty-free for demo)
- Headline: "Montana Hunting & Fishing Outfitters"
- Subhead: "Guided fishing and hunting trips from Helena, Montana. Over 25 years on the water."
- Two CTA buttons: "View Trips" → /trips/ | "Book a Trip" → /contact/

**Quick intro (2-3 sentences max):**

- Who Forrest is, what the operation offers, and why it's different
- License numbers for credibility (Outfitter #20194, USCG Captain)

**Trip cards (3 cards):**

- Fishing Trips → /trips/fishing/
- Hunting Trips → /trips/hunting/
- Packages → /trips/packages/
- Each card: photo, trip type, starting price, brief description, "Learn more" link

**Testimonials section:**

- 2-3 client testimonials with names
- Link to Google Reviews

**Seasonal availability callout:**

- Simple visual: "Now booking: Summer 2026 fishing trips"
- CTA to contact form

**Footer:**

- Contact info (phone, email, address)
- License numbers
- Quick links
- Social media links

#### Trips Overview (`/trips/`)

Simple hub page with cards linking to:

- Fishing Trips
- Hunting Trips
- Multi-Day Packages

Each card has a hero photo, short description, and starting price.

#### Fishing Trips (`/trips/fishing/`)

**Structure:** Single scrollable page with sections for each water/trip type:

- **Jet Boat Trips** — Missouri River "Land of Giants"
- **Drift Boat Trips** — Missouri River, Big Horn River, Blackfoot River
- **Lake Trips** — Canyon Ferry, Fort Peck, Holter Lake
- **Wade Trips** — Gallatin River, Shields River
- **Specialty** — Pike, Ice Fishing, Smallmouth Bass, Chinook Salmon, Lake Trout

Each section includes:

- 1-2 photos
- Brief description of the experience
- What's included
- Duration and pricing
- "Book This Trip" button → /contact/?trip=jet-boat (pre-fills inquiry form)

**Sidebar or sticky element:**

- Quick pricing summary table
- "Questions? Call Forrest: (406) 459-5352"

#### Hunting Trips (`/trips/hunting/`)

Same structure as fishing, sectioned by game type:

- Elk
- Deer (mule deer + whitetail)
- Bear
- Antelope

Each section: photo, season dates, pricing, what's included, "Book This Trip" CTA.

#### Packages (`/trips/packages/`)

Dedicated page for the multi-day offerings:

- Montana Triple Header Package
- Montana 6-Pack Package

Each package: full description, what's included, itinerary overview, pricing, photos, CTA.

#### Gallery (`/gallery/`)

- Responsive photo grid (masonry or similar)
- Filterable by category: Fishing | Hunting | Scenery | Camp
- Lazy-loaded for performance
- Photos optimized to WebP format
- Click to enlarge with lightbox

#### Contact / Book a Trip (`/contact/`)

**Trip Inquiry Form:**

```
- Name (required)
- Email (required)
- Phone (required)
- Trip Interest (dropdown: Fishing Trip, Hunting Trip, Multi-Day Package, Not Sure)
- Preferred Dates (text field or date range picker)
- Party Size (dropdown: 1, 2, 3, 4, 5+)
- Experience Level (dropdown: Beginner, Intermediate, Experienced)
- Message (textarea)
- Submit button
```

Form submissions:

- Email notification to Forrest immediately
- Auto-reply to the client: "Thanks for your inquiry! Forrest will get back to you within 24 hours. In the meantime, here's our gear list: [link]"

**Below the form:**

- Phone number (prominent)
- Email address
- Physical address
- Business hours / response time expectation
- Map embed (optional)

---

### Technical Spec

**Stack:**

- HTML/CSS/JS static site (Go templating for build, same as Firefly's standard approach)
- Tailwind CSS for styling
- No WordPress, no page builders
- Hosted on Firefly infrastructure

**Performance targets:**

- Google PageSpeed: 95+ mobile, 98+ desktop
- First Contentful Paint: < 1.5s
- Total page weight: < 500KB per page (excluding lazy-loaded gallery images)

**Image handling:**

- All images converted to WebP with fallback
- Responsive srcset for mobile/tablet/desktop
- Lazy loading on gallery and below-fold images
- Photo dimensions: hero images 1920x1080 max, card images 800x600 max, gallery thumbs 400x400

**Form handling:**

- Server-side form processing (no third-party form services)
- Email delivery via SMTP
- Basic spam protection (honeypot field + rate limiting)
- Auto-responder with gear list attachment or link

**SEO:**

- Structured data (LocalBusiness, TouristTrip schema)
- Meta descriptions per page
- Open Graph tags for social sharing
- Sitemap.xml
- Robots.txt

**Analytics:**

- Plausible or privacy-friendly alternative (same as Firefly standard)

---

### Content Needed from Client (for real build)

For the demo, we'll use placeholder content and publicly available info. For a real build, Forrest would need to provide:

1. **10-20 high-quality photos** — fish catches, scenery, boats, camp, happy clients
2. **Current pricing** for all trip types
3. **Trip descriptions** — what's included, duration, what to expect (can be gathered in a 30-min phone call)
4. **3-5 client testimonials** — name and where they're from
5. **Season/availability info** — when do bookings open, what months are which trips available
6. **Gear list** — already exists on current site
7. **FAQ content** — already exists on current site
8. **Bio/about info** — years of experience, background, licenses, personal story

---

### Demo Development Plan

**Phase 1 — Structure & Layout (Day 1-2)**

- Set up project scaffolding (Go templates, Tailwind, build pipeline)
- Build responsive layout shell with navigation
- Create homepage with hero, trip cards, testimonial section
- Create contact page with inquiry form

**Phase 2 — Trip Pages (Day 3-4)**

- Build fishing trips page with sectioned layout
- Build hunting trips page
- Build packages page
- Wire up "Book This Trip" buttons to pre-fill contact form

**Phase 3 — Gallery & Polish (Day 5)**

- Build filterable photo gallery with lightbox
- Optimize all images (WebP, srcset, lazy loading)
- Mobile testing and refinement
- PageSpeed optimization pass

**Phase 4 — Form Backend & Final (Day 6)**

- Implement form processing and auto-responder
- Add structured data / SEO meta
- Final QA across devices
- Deploy to demo URL (e.g., demo.fireflysoftware.dev/outfitter/ or packstring.dev)

**Estimated total dev time:** ~30-40 hours across 1-2 weeks

---

### Success Metrics (for pitching Forrest and other outfitters)

Once live, track:

- **Inquiry form submissions per month** vs. whatever Forrest currently gets via phone/email
- **Google PageSpeed scores** (before vs. after)
- **Google Search Console impressions** for "Helena fishing guide," "Montana walleye guide," etc.
- **Time on site / bounce rate** compared to industry benchmarks
- **Phone calls from website** (if using a tracked number)

The pitch to Forrest: "If your new site generates even one extra booking per season that you wouldn't have gotten otherwise, it's paid for itself 10x over."

---

### What This Demo Proves to Other Outfitters

When you show this demo to the next outfitter prospect, the conversation becomes:

> "Here's what I built for a Helena fishing and hunting outfitter. Notice how fast it loads, how easy it is to find trip info, and how the inquiry form works. Compare that to [pull up their current site on your phone]. Which one would you book a $4,000 hunt from? I can build you one just like this — $500 to build, $35/mo to run. That's less than one day's guide fee."

That's the entire sales pitch. The demo does the selling.
