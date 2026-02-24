# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Packstring is Firefly Software's outfitter/guide website product — a repeatable, high-quality website template system for Montana outfitters and guides. The pilot client is MT Hunt & Fish Outfitters (Forrest Fawthrop, Helena MT).

**Module:** `github.com/firefly/packstring`

## Stack

- **Go 1.25.1** — HTTP server using `net/http` + `html/template` (standard library, no framework)
- **Tailwind CSS** — CLI-based (`npx tailwindcss`), outputs to `static/css/output.css`
- **Alpine.js** — CDN/vendored, used for gallery filtering, lightbox, mobile nav toggle
- **htmx** — contact form POST with partial swap (no full page reload)
- **No database** — all content lives in Go structs and template files

## Build Commands (Makefile)

```
make dev          # Start Go server with Air hot reload + Tailwind watch
make build        # Build Go binary + compile Tailwind for production
make css-watch    # Tailwind CLI in watch mode
make css-build    # Tailwind CLI single build (minified)
make images       # Image optimization (WebP conversion + resize)
make deploy       # Build Docker image and deploy
make lighthouse   # Run Lighthouse CI audit
```

## Architecture

**Entry point:** `cmd/server/main.go` — HTTP server and route registration.

**Handlers:** `internal/handlers/` — `pages.go` serves all page routes, `contact.go` handles the POST `/contact` form submission.

**Templates:** Three-tier structure in `templates/`:
- `layouts/base.html` — master template (head, nav, footer, scripts)
- `partials/` — reusable components (nav, footer, trip-card, testimonial, gallery-item, contact-form)
- `pages/` — individual page templates composed from partials

**Static assets:** `static/` — compiled CSS, JS, images (organized by hero/trips/gallery/icons), fonts.

**Data:** Trip/service content is defined as Go structs (e.g., in a `trips.go` data file) passed to templates. No database.

## Site Routes (7 pages)

```
/                   → Homepage (hero, trip cards, testimonials, seasonal callout)
/trips/             → Trips overview hub (3 cards: fishing, hunting, packages)
/trips/fishing/     → Fishing trips (sections: jet boat, drift boat, lake, wade, specialty)
/trips/hunting/     → Hunting trips (sections: elk, deer, bear, antelope)
/trips/packages/    → Multi-day packages (Triple Header, 6-Pack)
/gallery/           → Photo gallery with category filtering and lightbox
/contact/           → Inquiry form + contact info (accepts ?trip= query param to pre-fill)
```

## Key Design Decisions

1. **Go `html/template` over templ** — standard library keeps the demo approachable; can migrate to templ later if this becomes a multi-client product.
2. **No build tool for JS** — Alpine.js and htmx are CDN includes or vendored files. No Webpack/Vite.
3. **Tailwind via CLI only** — watches template files, no bundler integration.
4. **Hot reload with Air** — configured in `.air.toml` for development.
5. **Honeypot field** for contact form spam prevention (no CAPTCHA).

## Color Palette

```
Primary:    #1B4332  (deep forest green)
Secondary:  #B68D40  (warm gold)
Accent:     #D4A574  (warm tan)
Background: #FEFAE0  (warm cream)
Text:       #1A1A1A  (near-black)
White:      #FFFFFF  (cards, form backgrounds)
```

The visual identity should feel like leather and campfire, not Silicon Valley.

## Performance Targets

- PageSpeed: 95+ mobile, 98+ desktop
- First Contentful Paint: < 1.5s
- Page weight: < 500KB per page (excluding lazy-loaded gallery images)
- No image over 200KB; all images WebP with responsive srcset (400w, 800w, 1600w)

## Detailed Documentation

See `docs/project-plan.md` for the full feature checklist, milestones, and git checkpoint conventions. See `docs/outfitter-service-offering-and-demo-spec.md` for page-level content specs and the business context.
