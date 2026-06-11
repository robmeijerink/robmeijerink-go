# CLAUDE CODE BRIEF — Commercial & SEO polish for three finished websites

## Context

Three live websites, one owner: **Rob Meijerink**, independent senior software engineer ("Solvalutions", KvK 98238175, Zutphen, Gelderland, NL). Stack: own multi-tenant **Go** backend on Google App Engine behind Cloudflare; templates in **Tailwind CSS v4** (tokens: `content-strong`, `content-base`, `content-muted`, `brand-base`, `brand-dark`, `accent-base`, `accent-dark`, `panel`, `panel-soft`).

- **solvalutions.nl** (Dutch) — THE business card. Must win assignments. Top priority.
- **solvalutions.com** (English) — tech showcase for international/remote work.
- **robmeijerink.nl** (Dutch) — personal authority/expertise site.

## PRIME DIRECTIVE — read twice

**The sites are FINISHED.** Design, structure, page set, tone of voice ("je"-form, approachable senior partner), and visual identity are final and approved.

- You may make **small tweaks** that sharpen what exists: titles, meta descriptions, headings, microcopy, alt texts, internal links, structured data, missing technical SEO plumbing.
- You may **NOT** redesign, restructure, add/remove pages, rewrite whole sections, change the tone, change the color/typography, or alter the image/identity in any way.
- Never break existing Tailwind classes, element IDs, or JS hooks (`blueprint-progress-bar`, `is-active`, `hero-card-zone`, etc.).
- When in doubt between "improve" and "leave alone": leave alone.

Every change you make is judged on ONE question: **does this make it more likely that a Dutch business owner contacts Rob for an assignment?**

## STRATEGIC REALITY — what these sites are actually for right now

Three fresh domains with no backlink authority will not rank for competitive commercial keywords for months. In Rob's first year, **the sites' #1 job is converting outreach-driven visitors**: a director or agency owner who received Rob's email or call and clicks through to verify he's real and credible. Rankings are the slow-burn bonus, not the engine.

Therefore the priority order for ALL work is:
1. **Conversion of a checking visitor** (trust, clarity, effortless contact — especially on mobile, since directors click through from email on their phones).
2. **Technical SEO hygiene** (cheap now, compounds later).
3. **Keyword targeting** (long-tail and local only; do not chase head terms).

Effort split across sites: **~80% solvalutions.nl, ~15% solvalutions.com, ~5% robmeijerink.nl.**

### Conversion checklist for solvalutions.nl (highest priority of the whole brief)
- [ ] `tel:` link works and is thumb-reachable on mobile on every page; mailto CTA pre-fills a subject.
- [ ] Trust strip visible without scrolling far: real portrait photo present (flag to Rob if missing — it is a conversion factor, not a detail), KvK, Zutphen/Oost-Nederland, "geen cookies, geen tracking".
- [ ] The knelpunt-analyse is explained in one sentence wherever it is the CTA (what you get, that it's fixed-price and small).
- [ ] **Flag for Rob's decision (do not implement unilaterally):** adding a concrete price indication to the knelpunt-analyse ("vanaf €X, vaste prijs") — a visible number lowers the contact threshold for directors. One line, his call.
- [ ] Mobile pass on every page: readable hero, no horizontal scroll, CTA above the fold.

### Measurement (correction to earlier "no analytics" stance)
Add **Cloudflare Web Analytics** (cookieless, aggregate, no cross-site tracking). It does not violate the "geen cookies, geen tracking van bezoekersgedrag" promise — it sets no cookies and profiles no one — and without it Rob is flying blind on whether the site converts. Do NOT add Google Analytics, pixels, or anything cookie-based. Keep the no-cookies trust line; it remains true.

## TECHNOLOGY NAMING — important correction

- **Lead with PHP and Laravel.** That is what Rob's buyers and their searches know ("Laravel developer", "PHP applicatie traag", "Laravel applicatie laten moderniseren"). 10+ years of professional PHP/Laravel is his strongest commercial asset.
- **Go is the performance differentiator**, mentioned as plain "Go" — Rob writes Go without frameworks. **Never mention "Fiber" or any Go framework anywhere on the sites.** If templates currently say "Fiber v3", change to "Go" (e.g. "Multi-tenant backend in Go").
- Positioning formula where both appear: *thuis in PHP/Laravel (where existing systems live), bouwt in Go (where speed and low running costs matter)*.

## COMMERCIAL COPY RULES (within existing text only — sharpen, don't rewrite)

1. Buyer language over tech language on solvalutions.nl: a director searches and reads in terms of problems ("applicatie traag", "systemen koppelen", "minder handwerk", "software laten bouwen"). Headings and meta must match that intent.
2. Every page keeps exactly one primary CTA: **"Plan een knelpunt-analyse"** (mailto) with "Of bel even" (tel:) as secondary. Make sure both are present, clickable, and above the fold reachable — but do not redesign CTA blocks.
3. Trust signals must be visible but subtle (footer level): KvK 98238175, "Gevestigd in Zutphen · heel Oost-Nederland en remote", phone, "geen cookies, geen tracking". Consistent NAP on every page and in schema.
4. Honesty is a commercial weapon here, not a compliance chore: directors smell inflated claims. Keep every number defensible; keep "referentie op aanvraag" labels; never invent clients, metrics or a team ("ik", never "wij"); never name the former employer. Fixed-price/outcome language only — no staffing words ("inhuur", "detachering", "join your team").

## SEO TASKS

### solvalutions.nl — do this first and most thoroughly

1. **Page titles & meta descriptions** (≤60 / ≤155 chars), written to sell the click, one intent per page. Target keyword themes:
   - Home: software moderniseren / vastgelopen software / MKB Oost-Nederland
   - /diensten: Laravel applicatie moderniseren · systemen koppelen / API koppeling laten maken · maatwerksoftware MKB · applicatie traag
   - /cases: ervaringen / resultaten softwaremodernisering
   - /over: freelance senior developer Zutphen / Gelderland
   - /contact: software developer Oost-Nederland contact
   The Go handler file already contains corrected titles/descriptions and JSON-LD (ProfessionalService with Zutphen/Gelderland + areaServed; Person on /over; Service ItemList on /diensten; real case items on /cases). Keep and refine — do not replace wholesale. Sweep all JSON-LD for "Fiber" and replace with "Go".
2. **Wire the sitemap**: `router.Get("/sitemap.xml", sitemap)` exists but is unregistered. Add `/robots.txt` (allow all, `Sitemap:` line).
3. **Open Graph / Twitter cards** on every page via an `OGMeta` key in each handler's `fiber.Map`, rendered in the master `<head>`: og:title, og:description, og:url, og:type, og:locale=nl_NL, og:image (one branded 1200×630 PNG per site, simple: logo + tagline in brand colors), twitter:card=summary_large_image.
4. **On-page hygiene**: exactly one H1 per page; descriptive Dutch alt text on images; explicit width/height + lazy-loading below the fold; self-referencing canonical per page; 404 page on-brand with links home (only if missing).
5. **Internal links**: ensure the path home → diensten → cases → contact is linked in copy, and aanpak ↔ diensten cross-link. Anchor text in buyer language ("bekijk wat ik voor je oplos"), no "klik hier".
6. **Suggest in a code comment** (do not fake in markup): create a Google Business Profile for local visibility in Zutphen/Oost-Nederland.

### solvalutions.com — light pass

- English metadata with the same discipline; og:locale=en_US; own sitemap/robots/canonicals.
- Keyword themes: PHP/Laravel modernization · legacy PHP to Go migration · Go backend engineer Netherlands/remote. No framework names for Go.
- Schema: same Organization entity (anchored to solvalutions.nl) with `sameAs`; areaServed remote/worldwide.
- Copy tweaks only to sharpen outcome framing for international clients; identical honesty rules.

### robmeijerink.nl — light pass

- Person schema is the centerpiece: `Person` (worksFor → Solvalutions, knowsAbout: PHP, Laravel, Go, legacy modernization, systems integration, DevOps; sameAs: LinkedIn `linkedin.com/in/robm89`, GitHub `github.com/robmeijerink`, solvalutions.nl, solvalutions.com).
- Articles (if/where present): Article/TechArticle schema with author → that Person; titles targeting informational long-tail ("waarom Go voor je backend", "PHP naar Go migreren"). Each article may end with one quiet link to solvalutions.nl — no selling on this site.
- This site must NOT carry service/pricing pages or compete with solvalutions.nl on commercial Dutch keywords.

### Cross-site (all three)

- Consistent entity wiring: robmeijerink.nl = canonical Person, solvalutions.nl = canonical Organization; all JSON-LD cross-references via sameAs/founder/worksFor.
- One-line footer cross-links between the sites (subtle, no link farm).
- No hreflang between .nl and .com (different content, not translations). Self-canonicals everywhere.
- No analytics other than cookieless Cloudflare Web Analytics; no cookies, chat widgets, or consent banners — ever.
- Email consistency: whichever public address is used (currently `rob.meijerink@gmail.com`), page content and JSON-LD must match exactly on all sites.

## FOR ROB — actions Claude Code cannot do (output these as a TODO list at the end of the run)

These have higher short-term ROI than any code change. Claude Code: print this list verbatim when done.

1. **Create a Google Business Profile** (category: software company / IT-dienstverlener, location Zutphen, service area Oost-Nederland, link to solvalutions.nl, add the portrait photo). For local queries this outranks anything on-page for months. Free, ~30 minutes.
2. **Upload a real portrait photo** wherever the templates have a placeholder — it is a measurable conversion factor for a trust-driven local market.
3. **Align LinkedIn headline + email signature** word-for-word with the solvalutions.nl hero proposition. Consistency across the three touchpoints is what flips perception from "candidate" to "supplier".
4. Decide on the **price indication for the knelpunt-analyse** (see conversion checklist) and the **public email address** (gmail vs info@solvalutions.nl) — then tell Claude Code to apply it everywhere at once.
5. Submit each domain in **Google Search Console** and submit the sitemaps after this run.

## DEFINITION OF DONE

- [ ] Conversion checklist (mobile tel/CTA, trust strip, knelpunt-analyse one-liner) verified on every solvalutions.nl page.
- [ ] Cookieless Cloudflare Web Analytics snippet present on all three sites; nothing cookie-based anywhere.
- [ ] "FOR ROB" TODO list printed at the end of the run.
- [ ] Zero occurrences of "Fiber" (or any Go framework name) across all three sites, templates, and JSON-LD.
- [ ] PHP/Laravel visibly present in diensten/services copy and metadata as the primary professional competence.
- [ ] Unique, buyer-intent title + meta description on every page of all three sites.
- [ ] sitemap.xml route registered + robots.txt live on each domain.
- [ ] OG/Twitter meta + og-image rendering on every page.
- [ ] Valid, page-matching JSON-LD (nothing in schema that isn't visible on the page).
- [ ] One H1 per page; alt texts; canonicals; 404 present.
- [ ] CTA (knelpunt-analyse + bel) intact and clickable on every solvalutions.nl page.
- [ ] Claims audit: report any superlative/number found in copy for Rob's final judgment — change nothing substantive without flagging.
- [ ] Lighthouse ≥95 on performance/SEO/accessibility/best-practices per site.
- [ ] Diff review: confirm no layout, structure, tone, or identity changes anywhere.

## WHAT NOT TO DO

- No redesigns, no new pages, no removed pages, no rewritten sections, no tone changes.
- No "Fiber" or Go framework names anywhere.
- No invented clients, metrics, testimonials, or team; no former-employer name.
- No staffing/secondment language.
- No cookie-based analytics, pixels, chat widgets, or consent banners (cookieless Cloudflare Web Analytics is the only allowed measurement).
- No translating solvalutions.nl into English for the .com.
- No thin doorway pages per city/keyword.
