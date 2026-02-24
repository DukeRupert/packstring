# Packstring — UI Style Guide
**For AI agents and developers working on Packstring outfitter sites**
Version 1.0 · Working Ranch Structure + Golden Hour Atmosphere

---

## How to Use This Document

This is a machine-readable style guide for building outfitter and guide service websites under the Packstring brand. When modifying or creating UI for a Packstring site:

1. Import `packstring-design-system.css` — it contains all tokens and base component styles
2. Use only the token names defined in that file — never hardcode hex values or pixel sizes
3. Follow the rules in this document before writing any CSS or HTML
4. When in doubt: let the photography do the heavy lifting. The design is the frame, not the painting

---

## Brand Overview

**Packstring** is a standalone product by Firefly Software — a productized website package for Montana outfitters, hunting guides, and fishing operations.

Packstring is **not** a sub-brand of Firefly. It has its own identity, palette, typography, and voice. The two brands share technical infrastructure and quality standards, but their audiences and aesthetics are entirely different.

---

## Core Aesthetic

**"The Montana they came here to find."**

Packstring sites sell the *experience before the experience*. The audience is out-of-state sportsmen, families, and tourists who have an idealized, cinematic vision of Montana — think Yellowstone the TV show, Filson catalogs, Yeti commercials. The website must match the movie playing in their heads.

The site should make a dentist from Connecticut feel like booking this trip is the most important decision he'll make all year. Clean enough to feel trustworthy, atmospheric enough to make someone reach for their credit card.

**The positioning spectrum:**

```
Kitschy/Camo ←————————|—→ Packstring ——→ Luxury Resort
(Cabela's catalog)     (Filson / Yeti / Patagonia)  (Amangiri)
```

Premium-outdoor. Aspirational but not unapproachable. The guide knows what he's doing, the scenery is unreal, and this trip will be the story you tell at every dinner party for five years.

### Design Principles

1. **Photo-forward** — The landscape and the catch/harvest do the selling. Design exists to frame great photography, not compete with it.
2. **Two-click clarity** — A tourist on their phone in a hotel lobby can find trip info and pricing in two taps. No nav mazes, no buried content.
3. **Premium but approachable** — Like a well-built lodge, not a five-star hotel. Professional, not pretentious.
4. **Trustworthy** — License numbers, real testimonials, real pricing, no fluff. These people are spending $500–$5,000. They need to feel confident.

---

## What NOT to Do

These are hard rules. Do not violate them.

```
✗ No border-radius larger than 4px — not on buttons, cards, inputs, images, nothing
✗ No purple, violet, neon, or blue-toned gradients
✗ No Inter, Roboto, Arial, or system-ui fonts
✗ No decorative illustrations, icon packs, or emoji in UI
✗ No box-shadows larger than 0 4px 20px (and only on buttons/cards)
✗ No parallax scrolling
✗ No JS-driven motion libraries
✗ No color values hardcoded — always use CSS custom properties
✗ No content changes — preserve all client copy, links, prices, and testimonials exactly
✗ No background images from external CDNs — inline SVG only for decorative backgrounds
✗ No star fields or firefly animations — those belong to the Firefly brand
✗ No antler fonts, wood textures, camo patterns, or rustic clip art
✗ No stock photography — use client photos or high-quality royalty-free Montana imagery
✗ Do not touch Go template tags ({{ }}) in .templ or .html files
```

---

## Color System

### Light Surfaces
| Token | Hex | Use |
|---|---|---|
| `--color-sand` | `#F2EDE4` | Page background |
| `--color-sand-dk` | `#DED6C8` | Borders, dividers, grid gaps |
| `--color-sand-lt` | `#F8F5EF` | Card hover state, light wells |
| `--color-cream` | `#FAF8F3` | Card backgrounds, form inputs |

### Dark Surfaces — Timber
| Token | Hex | Use |
|---|---|---|
| `--color-timber` | `#1A1410` | Deepest dark — nav, footer, hero overlays |
| `--color-timber-mid` | `#2A2118` | Dark section backgrounds |
| `--color-timber-lt` | `#3A2E22` | Cards on dark bg, dark surface borders |
| `--color-timber-edge` | `#4A3D2E` | Subtle dividers on dark bg |

### Ink (Text)
| Token | Hex | Use |
|---|---|---|
| `--color-ink` | `#1E1A14` | Primary text |
| `--color-ink-mid` | `#3D3428` | Secondary text |
| `--color-ink-faded` | `#6B5F50` | Captions, descriptions |
| `--color-stone` | `#8A7E6E` | Placeholders, metadata |

### Core Brand — Saddle
| Token | Hex | Use |
|---|---|---|
| `--color-saddle` | `#6B3A1F` | Brand identity, headings on feature sections, borders |
| `--color-saddle-lt` | `#7E4828` | Saddle hover |
| `--color-saddle-dk` | `#5A3018` | Saddle pressed/active |

### Primary Action — Forest Green
| Token | Hex | Use |
|---|---|---|
| `--color-forest` | `#2E6B45` | Primary CTA buttons, action links |
| `--color-forest-lt` | `#3A8458` | Forest hover |
| `--color-forest-dk` | `#245838` | Forest pressed/active |

### Wonder Accent — Copper
| Token | Hex | Use |
|---|---|---|
| `--color-copper` | `#B8652A` | Warm accent, eyebrows/kickers, golden hour glow |
| `--color-copper-lt` | `#D07A35` | Copper hover, gradient highlights |
| `--color-copper-dim` | `rgba(184,101,42,0.15)` | Ambient copper fills |
| `--color-copper-glow` | `rgba(184,101,42,0.08)` | Very subtle warmth on dark sections |

### Supporting Colors
| Token | Hex | Use |
|---|---|---|
| `--color-river` | `#3A6B6E` | Info states, water/fishing context, tech badges |
| `--color-river-lt` | `#4A8A8E` | River hover |
| `--color-sage` | `#5A7A52` | Success states, secondary actions |
| `--color-sage-lt` | `#6E9466` | Sage hover |
| `--color-golden` | `#C8900A` | Star ratings, premium callouts |
| `--color-golden-lt` | `#E0A820` | Golden hover |

### Semantic Mapping
| Purpose | Color |
|---|---|
| Primary CTA | `--color-forest` |
| Secondary CTA | `--color-saddle` |
| Error / Warning | `--color-copper` |
| Success | `--color-sage` |
| Info | `--color-river` |
| Eyebrow / Kicker | `--color-copper` |
| Link text (body) | `--color-forest` |
| Nav border accent | `--color-copper` |

---

## Typography

### Font Stack
```css
--font-display: 'Bitter', Georgia, serif;
--font-body:    'Lora', Georgia, serif;
--font-ui:      'Barlow Semi Condensed', 'Arial Narrow', sans-serif;
--font-mono:    'Source Code Pro', 'Courier New', monospace;
```

**Google Fonts import (add to every page `<head>`):**
```html
<link href="https://fonts.googleapis.com/css2?family=Bitter:ital,wght@0,400;0,500;0,600;0,700;0,800;1,400&family=Lora:ital,wght@0,400;0,500;0,600;0,700;1,400;1,500&family=Barlow+Semi+Condensed:wght@400;500;600;700&family=Source+Code+Pro:wght@400;600&display=swap" rel="stylesheet">
```

### When to Use Each Font

| Font | Use for | Never use for |
|---|---|---|
| Bitter | Hero titles, section headings, trip names, prices, stat numbers, testimonial quotes | Body text, labels, nav, buttons |
| Lora | All body copy, card descriptions, trip details, form helper text | Headings, UI labels |
| Barlow Semi Condensed | Nav links, buttons, section labels, card eyebrows, form labels, badges | Body text, long-form content |
| Source Code Pro | License numbers, coordinates, metadata, season dates, version numbers | Headings, body copy |

### Type Hierarchy

```
Hero title:      Bitter 800, clamp(36px, 5vw, 64px), line-height 1.05
Section heading: Bitter 800, clamp(26px, 3.5vw, 42px), line-height 1.1
Trip name:       Bitter 700, 20-24px, line-height 1.15
Card title:      Bitter 700, 16-18px, line-height 1.2
Price:           Bitter 600-700, 22-28px
Section label:   Barlow Semi Condensed 600, 10-11px, uppercase, letter-spacing 0.35em
Nav / button:    Barlow Semi Condensed 600-700, 12-13px, uppercase, letter-spacing 0.15-0.2em
Body:            Lora 400, 15-16px, line-height 1.7
Italic body:     Lora 400 italic (subtitles, subheads, pull quotes)
Metadata:        Source Code Pro 400, 10-11px, letter-spacing 0.08em
Kicker/eyebrow:  Barlow Semi Condensed 600, 11px, uppercase, letter-spacing 0.35em, color --color-copper
```

---

## Layout

- **Max content width:** `1100px`, centered with `margin-inline: auto`
- **Section padding:** `80px 40px` desktop, `56px 20px` mobile
- **Card grid gap:** `2px` — the container background color becomes the visible divider line
- **Border radius:** `4px` on all interactive elements (buttons, cards, inputs, images)
- **Photo aspect ratios:** Hero 16:9 or 3:2, cards 4:3, gallery thumbs 1:1

### Grid Divider Pattern
```html
<!-- Light grid (trip cards): -->
<div style="display:grid; grid-template-columns:repeat(3,1fr); gap:2px; background:var(--color-sand-dk);">
  <div style="background:var(--color-cream); border-radius:4px;">Card 1</div>
  <div style="background:var(--color-cream); border-radius:4px;">Card 2</div>
  <div style="background:var(--color-cream); border-radius:4px;">Card 3</div>
</div>

<!-- Dark grid (features/stats): -->
<div style="display:grid; grid-template-columns:repeat(3,1fr); gap:2px; background:var(--color-timber);">
  <div style="background:var(--color-timber-mid); border-radius:4px;">Stat 1</div>
  <div style="background:var(--color-timber-mid); border-radius:4px;">Stat 2</div>
  <div style="background:var(--color-timber-mid); border-radius:4px;">Stat 3</div>
</div>
```

---

## Components

### Navigation
```html
<nav class="nav">
  <a class="nav-brand" href="/">
    <!-- Client logo or text brand -->
    Montana Hunt & Fish
  </a>
  <ul class="nav-links">
    <li><a href="/trips">Trips</a></li>
    <li><a href="/gallery">Gallery</a></li>
    <li><a href="/about">About</a></li>
  </ul>
  <a class="nav-cta" href="/contact">Book a Trip</a>
</nav>
```
- Background: `--color-timber`
- Bottom border: `2px solid --color-copper`
- Active link: copper underline, `scaleX` 0→1 from left
- CTA button: `--color-forest` background, hover `--color-forest-lt`
- Brand text: Bitter 700, `--color-cream`

### Buttons
```html
<button class="btn btn-primary">Book This Trip</button>
<button class="btn btn-secondary">Learn More</button>
<button class="btn btn-outline">View All Trips</button>
<button class="btn btn-outline-light">View All Trips (dark bg)</button>
<button class="btn btn-ghost">Dismiss</button>
<!-- Sizes: btn-sm, btn-lg -->
```

| Variant | Background | Text | Border | Use |
|---|---|---|---|---|
| `btn-primary` | `--color-forest` | white | none | Primary CTA: Book, Inquire, Submit |
| `btn-secondary` | `--color-saddle` | white | none | Secondary: Learn More, View Details |
| `btn-outline` | transparent | `--color-forest` | `--color-forest` | Tertiary on light bg |
| `btn-outline-light` | transparent | `--color-cream` | `--color-cream` | Tertiary on dark bg |
| `btn-ghost` | transparent | `--color-ink-faded` | none | Dismiss, Cancel |

All buttons: `border-radius: 4px`, `font-family: var(--font-ui)`, `font-weight: 600`, `font-size: 13px`, `letter-spacing: 0.15em`, `text-transform: uppercase`.

### Trip Cards
```html
<div class="card-grid card-grid--3">
  <div class="trip-card">
    <img class="trip-card-photo" src="/images/trips/canyon-ferry.webp" alt="Canyon Ferry walleye fishing" loading="lazy">
    <div class="trip-card-body">
      <div class="trip-card-eyebrow">Starting at</div>
      <div class="trip-card-price">$450 / day</div>
      <div class="trip-card-title">Missouri River Jet Boat</div>
      <p class="trip-card-desc">Full-day guided trip through the Land of Giants. Trophy walleye, smallmouth, and pike.</p>
      <a href="/contact/?trip=jet-boat" class="btn btn-primary">Book This Trip</a>
    </div>
  </div>
</div>
```

### Forms
```html
<div class="form-grid">
  <div class="form-group">
    <label class="form-label">Name</label>
    <input class="form-input" type="text" placeholder="First and last">
  </div>
  <div class="form-group">
    <label class="form-label">Trip Interest</label>
    <select class="form-select">
      <option>Fishing Trip</option>
      <option>Hunting Trip</option>
      <option>Multi-Day Package</option>
      <option>Not Sure Yet</option>
    </select>
  </div>
  <div class="form-group form-group--full">
    <label class="form-label">Message</label>
    <textarea class="form-textarea" placeholder="Tell us about your trip — dates, group size, experience level, anything."></textarea>
  </div>
</div>
```

Form inputs: `background: var(--color-cream)`, `border: 1px solid var(--color-sand-dk)`, `border-radius: 4px`. Focus state: `border-color: var(--color-forest)`, `box-shadow: 0 0 0 2px var(--color-forest-lt)` at 20% opacity.

### Badges
```html
<!-- Outline variants -->
<span class="badge badge--saddle">Jet Boat</span>
<span class="badge badge--forest">Full Day</span>
<span class="badge badge--copper">Premium</span>
<span class="badge badge--river">Walleye</span>
<span class="badge badge--stone">Spring–Fall</span>

<!-- Solid variants -->
<span class="badge badge--solid-forest">Available</span>
<span class="badge badge--solid-saddle">Limited</span>
<span class="badge badge--solid-copper">Featured</span>
```

### Testimonials
```html
<section class="section--testimonial">
  <blockquote class="testimonial">
    <p class="testimonial-text">"Best fishing trip I've ever been on. Forrest put us on fish all day and the scenery was unbelievable."</p>
    <cite class="testimonial-cite">
      <span class="testimonial-name">Mark & Sarah Thompson</span>
      <span class="testimonial-origin">Dallas, TX</span>
    </cite>
  </blockquote>
</section>
```

Testimonial section background: `--color-timber-mid`. Quote text: Bitter 700 italic, `--color-cream`. Cite: Barlow Semi Condensed 600, `--color-stone`.

### Alerts
```html
<div class="alert alert--success">
  <div class="alert-icon"><!-- 16px SVG --></div>
  <div>
    <div class="alert-title">Inquiry Received</div>
    <div class="alert-body">Thanks! Forrest will get back to you within 24 hours.</div>
  </div>
</div>
<!-- Variants: alert--info (river), alert--success (sage), alert--warning (copper), alert--error (copper dark) -->
```

### Progress / Availability Bars
```html
<div class="progress-track">
  <div class="progress-fill progress-fill--forest" style="width:75%"></div>
</div>
<!-- Variants: progress-fill--forest, progress-fill--copper, progress-fill--river -->
```

Use for simple availability displays: "Summer 2026: 75% booked"

---

## Section Patterns

### Dark Section (Golden Hour)
```html
<section class="section--dark">
  <div class="golden-glow"></div>
  <div class="container">
    <!-- content here -->
  </div>
</section>
```

Unlike Firefly's star fields and firefly animations, Packstring dark sections use a subtle warm radial gradient — a golden-hour glow effect suggesting sunset or campfire light. No animated particles.

```css
.golden-glow {
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 120%;
  height: 60%;
  background: radial-gradient(
    ellipse at 50% 100%,
    var(--color-copper-glow) 0%,
    transparent 70%
  );
  pointer-events: none;
}
```

### Hero Section
```html
<section class="hero">
  <img class="hero-photo" src="/images/hero.webp" alt="Missouri River at sunrise">
  <div class="hero-overlay"></div>
  <div class="hero-content">
    <h1 class="hero-title">Montana Hunting &<br>Fishing Outfitters</h1>
    <p class="hero-sub">Guided fishing and hunting trips from Helena, Montana. Over 25 years on the water.</p>
    <div class="hero-actions">
      <a class="btn btn-primary btn-lg" href="/trips">View Trips</a>
      <a class="btn btn-outline-light btn-lg" href="/contact">Book a Trip</a>
    </div>
  </div>
</section>
```

Hero uses a full-width photo with a gradient overlay (`--color-timber` at 40-60% opacity, heavier at bottom). Content sits on top with `position: relative; z-index: 2`.

### Pricing Strip
```html
<div class="pricing-strip">
  <div class="pricing-strip-text">Summer 2026 fishing trips now booking. Starting at $400/day.</div>
  <a class="btn btn-primary btn-sm" href="/contact">Reserve Your Dates</a>
</div>
```

Background: `--color-saddle`. Text: `--color-cream`, Lora 400 italic. CTA: forest green button.

### Tagline Strip
```html
<div class="tagline-strip">
  <div class="tagline-text">"One extra booking pays for your website for the entire year."</div>
  <div class="tagline-meta">— Packstring by Firefly Software</div>
</div>
```

Background: `--color-timber`. Text: Bitter 700 italic, `--color-cream`. Meta: Source Code Pro 400, `--color-stone`.

---

## Photography Guidelines

Photography is the single most important element on a Packstring site. The design exists to frame it.

### Requirements
- **All photos converted to WebP** with JPEG fallback
- **Responsive `srcset`** for mobile/tablet/desktop breakpoints
- **Lazy loading** on all images below the fold
- **Maximum dimensions:** Hero 1920×1080, trip cards 800×600, gallery thumbs 400×400

### Photo Style
- **Warm color grading** — lean into golden hour tones. Slightly warm white balance.
- **Show people** — clients holding fish, guides pointing at game, groups around camp. Not just empty landscapes.
- **Action and aftermath** — the cast, the catch, the campfire. Not posed portraits.
- **Cinematic framing** — wide shots with depth, not flat snapshots.

### What to Avoid
- Grip-and-grin photos with dead eyes and fluorescent lighting
- Blurry phone photos (if client photos are low quality, crop tight and apply warm grade)
- Stock photography of any kind
- Photos with visible brand logos from competitors

---

## Animation Rules

Packstring allows slightly more motion than Firefly, focused on revealing photography.

| Element | Animation | Trigger | Duration |
|---|---|---|---|
| Nav underline | `scaleX` 0→1 from left | `:hover`, `.active` | `0.22s ease` |
| Button lift | `translateY(-1px)` | `:hover` | `0.15s ease` |
| Card top bar | `scaleX` 0→1 from left | `:hover` | `0.25s ease` |
| Color/border transitions | — | `:hover`, `:focus` | `0.12s–0.18s` |
| Photo reveal | `opacity 0→1, translateY(12px→0)` | Scroll into view | `0.5s ease-out` |
| Section content | `opacity 0→1, translateY(8px→0)` | Scroll into view | `0.4s ease-out` |

### Animation Constraints
- **Photo/section reveals:** CSS-only, triggered by `IntersectionObserver` adding a `.visible` class. No JS animation libraries.
- **No parallax** — ever.
- **No entrance animations on above-fold content** — hero loads immediately, no fade-in.
- **No infinite/looping animations** — no spinning, pulsing, or blinking elements.
- **Stagger reveals** when multiple cards enter view simultaneously: `animation-delay: 0.08s` per card.
- **`prefers-reduced-motion`** — respect the media query. All scroll-triggered animations should be disabled if the user prefers reduced motion.

```css
/* Base state for reveal elements */
.reveal {
  opacity: 0;
  transform: translateY(12px);
  transition: opacity 0.5s ease-out, transform 0.5s ease-out;
}

.reveal.visible {
  opacity: 1;
  transform: translateY(0);
}

@media (prefers-reduced-motion: reduce) {
  .reveal {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
```

---

## Responsive Breakpoints

```css
/* Mobile first */
--bp-sm:  640px;   /* Large phones */
--bp-md:  768px;   /* Tablets */
--bp-lg:  1024px;  /* Small desktop */
--bp-xl:  1280px;  /* Full desktop */
```

### Mobile Priorities
Outfitter site visitors are frequently browsing on phones — in hotels, airports, or around the campfire. Mobile is not secondary.

1. **Trip info and pricing visible without scrolling** on the trips page
2. **Phone number tap-to-call** on every page
3. **Inquiry form usable with one thumb** — large tap targets, smart defaults
4. **Gallery loads fast** — thumbnails first, full images on tap
5. **Navigation collapses to hamburger** at `--bp-md` with full-screen overlay

---

## Tech Stack Context

- **Go** + `html/template` or `templ` — preserve all `{{ }}` tags
- **HTMX** — preserve all `hx-*` attributes exactly
- **Alpine.js** — preserve all `x-*` attributes exactly
- **Tailwind CSS** — utility classes for layout, custom properties for brand tokens
- Do not convert server-rendered components to client-side JavaScript

### Performance Targets
- Google PageSpeed: **95+ mobile**, 98+ desktop
- First Contentful Paint: **< 1.5s**
- Total page weight: **< 500KB** per page (excluding lazy-loaded gallery images)
- All images in **WebP** with responsive `srcset`

---

## SEO Requirements

Every Packstring site must include:

- **Structured data:** `LocalBusiness` + `TouristTrip` schema on relevant pages
- **Meta descriptions:** Unique per page, include location + service type
- **Open Graph tags:** For social sharing — photo, title, description per page
- **Sitemap.xml** and **robots.txt**
- **Alt text on every image** — descriptive, include species/location where relevant
- **Tap-to-call link** on phone numbers: `<a href="tel:+14064595352">`

---

## Quick Reference

**Background for this section?**
- Default page → `--color-sand`
- Cards/surfaces → `--color-cream`
- Nav/footer → `--color-timber`
- Dark features, testimonials → `--color-timber-mid` + golden glow
- Pricing/seasonal strip → `--color-saddle`
- CTA strip → `--color-forest`

**Font for this element?**
- Big title, trip name, price → Bitter
- Label, button, badge, nav → Barlow Semi Condensed
- Body paragraph, trip description → Lora
- License number, season dates, metadata → Source Code Pro

**Accent color?**
- Primary CTA → `--color-forest`, hover `--color-forest-lt`
- Secondary CTA → `--color-saddle`, hover `--color-saddle-lt`
- Eyebrow/kicker → `--color-copper`
- Warm atmospheric glow → `--color-copper-glow`
- Success → `--color-sage`
- Info/water context → `--color-river`
- Star ratings, premium → `--color-golden`

**Dark section checklist:**
- [ ] `section--dark` class
- [ ] `.golden-glow` radial gradient element
- [ ] Content in `.container` with `position: relative; z-index: 2`
- [ ] Text colors: headings `--color-cream`, body `--color-sand-dk`, meta `--color-stone`

---

## Differences from Firefly

For developers familiar with the Firefly style guide, here are the key differences:

| Aspect | Firefly | Packstring |
|---|---|---|
| Audience | Business owners buying web dev | Tourists booking outdoor trips |
| Aesthetic | Montana Utility + Wilderness Wonder | Working Ranch + Golden Hour |
| Border radius | 0px everywhere | 4px everywhere |
| Display font | Playfair Display | Bitter |
| Body font | Libre Baskerville | Lora |
| UI font | Oswald | Barlow Semi Condensed |
| Dark bg | Midnight navy (#0A0E18) | Dark timber (#1A1410) |
| CTA color | Rust (#8B3A1A) | Forest green (#2E6B45) |
| Accent | Amber (#E8922A) | Copper (#B8652A) |
| Atmosphere | Star fields + firefly animations | Golden hour radial gradient |
| Scroll animation | None allowed | Subtle photo/section fade-in |
| Content tone | Technical, direct, developer-speak | Aspirational, cinematic, story-driven |
