package solvalutions_nl

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
	router.Get("/cases", cases)
	router.Get("/contact", contact)

	router.Get("/diensten", services)
	router.Get("/integraties", integrations)
	router.Get("/websites", websites)
	router.Get("/conversie", conversion)
	router.Get("/apps", apps)
	router.Get("/dashboards", dashboards)
	router.Get("/werkgebied", area)
	router.Get("/over", about)
	router.Get("/aanpak", approach)

	router.Get("/sitemap.xml", sitemap)
}

func home(c fiber.Ctx) error {
	title := "Website laten maken & software op maat Zutphen | Solvalutions"
	desc := "Website laten maken in Zutphen? Ook webshops, apps, koppelingen en software op maat voor MKB en logistiek. Direct met de engineer, vaste prijs per fase."
	url := "https://solvalutions.nl"

	jsonLd := template.HTML(`
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "ProfessionalService",
      "name": "Solvalutions",
      "url": "` + url + `",
      "logo": "` + url + `/assets/img/solvalutions-logo.png",
      "image": "` + url + `/assets/img/solvalutions-og-image.png",
      "description": "` + desc + `",
      "telephone": "+31649691374",
      "email": "info@solvalutions.nl",
      "taxID": "NL005319604B24",
      "priceRange": "€€€",
      "founder": {
        "@type": "Person",
        "name": "Rob Meijerink",
        "url": "https://robmeijerink.nl"
      },
      "address": {
        "@type": "PostalAddress",
        "addressLocality": "Zutphen",
        "addressRegion": "Gelderland",
        "addressCountry": "NL"
      },
      "areaServed": [
        { "@type": "AdministrativeArea", "name": "Gelderland" },
        { "@type": "AdministrativeArea", "name": "Overijssel" },
        { "@type": "AdministrativeArea", "name": "Utrecht" },
        { "@type": "AdministrativeArea", "name": "Drenthe" },
        { "@type": "AdministrativeArea", "name": "Flevoland" },
        { "@type": "City", "name": "Zutphen" },
        { "@type": "City", "name": "Apeldoorn" },
        { "@type": "City", "name": "Deventer" },
        { "@type": "City", "name": "Arnhem" },
        { "@type": "City", "name": "Nijmegen" },
        { "@type": "City", "name": "Enschede" },
        { "@type": "City", "name": "Zwolle" },
        { "@type": "City", "name": "Amersfoort" },
        { "@type": "City", "name": "Utrecht" },
        { "@type": "Country", "name": "Nederland" }
      ],
      "knowsAbout": [
        "Softwaremodernisering",
        "Systeem- en ERP-integraties",
        "API-ontwikkeling",
        "Automatisering met AI",
        "Backend-ontwikkeling",
        "Websites",
        "DevOps",
        "Cloud-infrastructuur"
      ],
      "sameAs": [
        "https://linkedin.com/in/robm89",
        "https://github.com/robmeijerink"
      ]
    }
    </script>
    `)

	return web.Render(c, "home", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  jsonLd + faqStructuredData(homeFAQ),
		"FAQ":             homeFAQ,
	})
}

func contact(c fiber.Ctx) error {
	title := "Contact: even sparren over je ICT-vraagstuk | Solvalutions"
	desc := "Loopt er iets vast in je software? Stuur een mail of bel even. Je spreekt rechtstreeks met Rob, zonder accountmanager. Vrijblijvend."
	url := "https://solvalutions.nl"

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "ContactPage",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "url": "` + url + `/contact",
	  "isPartOf": {
	    "@type": "WebSite",
	    "name": "Solvalutions",
	    "url": "` + url + `"
	  },
	  "mainEntity": {
	    "@type": "Organization",
	    "name": "Solvalutions",
	    "url": "` + url + `",
	    "logo": "` + url + `/assets/img/solvalutions-logo.png",
	    "founder": {
	      "@type": "Person",
	      "name": "Rob Meijerink",
	      "url": "https://robmeijerink.nl"
	    },
	    "address": {
	      "@type": "PostalAddress",
	      "addressLocality": "Zutphen",
	      "addressRegion": "Gelderland",
	      "addressCountry": "NL"
	    },
	    "contactPoint": {
	      "@type": "ContactPoint",
	      "telephone": "+31649691374",
	      "email": "info@solvalutions.nl",
	      "contactType": "customer service",
	      "areaServed": "NL",
	      "availableLanguage": ["nl", "en"]
	    }
	  }
	}
	</script>
	`)

	return web.Render(c, "contact", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  jsonLd + breadcrumbs("/contact", "Contact"),
	})
}

func about(c fiber.Ctx) error {
	title := "Over Solvalutions: Rob Meijerink, software architect in Zutphen"
	desc := "Sinds 2011 actief als developer en architect, nu zelfstandig vanuit Zutphen. Eén engineer, zonder tussenlagen: je werkt rechtstreeks met de persoon die het bouwt."
	url := "https://solvalutions.nl"

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
      "@context": "https://schema.org",
      "@type": "AboutPage",
      "name": "` + title + `",
      "description": "` + desc + `",
      "url": "` + url + `/over",
      "isPartOf": {
        "@type": "WebSite",
        "name": "Solvalutions",
        "url": "` + url + `"
      },
      "mainEntity": {
        "@type": "Person",
        "name": "Rob Meijerink",
        "jobTitle": "Software architect",
        "description": "Software architect en developer, actief sinds 2011, gespecialiseerd in softwaremodernisering, systeemkoppelingen, automatisering en maatwerksoftware.",
        "url": "https://robmeijerink.nl",
        "image": "` + url + `/assets/img/foto-rob-meijerink-solvalutions-avatar-320.jpg",
        "address": {
          "@type": "PostalAddress",
          "addressLocality": "Zutphen",
          "addressRegion": "Gelderland",
          "addressCountry": "NL"
        },
        "knowsLanguage": ["nl", "en", "de", "es"],
        "knowsAbout": [
          "Softwaremodernisering",
          "Systeem- en ERP-integraties",
          "API-ontwikkeling",
          "Automatisering met AI",
          "Backend-ontwikkeling",
          "Websites",
          "DevOps",
          "Cloud-infrastructuur"
        ],
        "alumniOf": {
          "@type": "EducationalOrganization",
          "name": "Saxion Hogescholen"
        },
        "worksFor": {
          "@type": "Organization",
          "name": "Solvalutions",
          "url": "` + url + `"
        },
        "sameAs": [
          "https://linkedin.com/in/robm89",
          "https://github.com/robmeijerink"
        ]
      }
    }
	</script>
	`)

	// Lekker flexibel via de fiber.Map
	return web.Render(c, "about", fiber.Map{
		"CTATitle":        "Zullen we kennismaken?",
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  jsonLd + breadcrumbs("/over", "Over"),
	})
}

func approach(c fiber.Ctx) error {
	title := "Aanpak: vaste prijs per fase | Solvalutions"
	desc := "Klein beginnen, vaste prijs per fase, naast je systeem bouwen. Zo werk ik aan modernisering, koppelingen en maatwerk."
	url := "https://solvalutions.nl"

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "WebPage",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "url": "` + url + `/aanpak",
	  "isPartOf": {
	    "@type": "WebSite",
	    "name": "Solvalutions",
	    "url": "` + url + `"
	  },
	  "about": {
	    "@type": "Thing",
	    "name": "Werkwijze bij softwaremodernisering en systeemintegratie"
	  },
	  "publisher": {
	    "@type": "Organization",
	    "name": "Solvalutions",
	    "url": "` + url + `",
	    "logo": "` + url + `/assets/img/solvalutions-logo.png"
	  }
	}
	</script>
	`)

	return web.Render(c, "approach", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  jsonLd + breadcrumbs("/aanpak", "Aanpak"),
	})
}

func cases(c fiber.Ctx) error {
	title := "Cases: opgeloste software- en systeemvraagstukken | Solvalutions"
	desc := "Een selectie van trajecten: gemoderniseerde kernsystemen, snellere dataverwerking, maatwerk dat handwerk wegnam en koppelingen tussen systemen."
	url := "https://solvalutions.nl"

	structuredData := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "CollectionPage",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "url": "` + url + `/cases",
	  "isPartOf": {
	    "@type": "WebSite",
	    "name": "Solvalutions",
	    "url": "` + url + `"
	  },
	  "mainEntity": {
	    "@type": "ItemList",
	    "itemListElement": [
	      {
	        "@type": "ListItem",
	        "position": 1,
	        "name": "Kernsysteem dat niet meer efficiënt werkte stap voor stap herbouwd, zonder dat het bedrijf stillag",
	        "description": "Legacy-modernisering: herbouw naar gestructureerde, geteste architectuur draaiend naast de bestaande omgeving."
	      },
	      {
	        "@type": "ListItem",
	        "position": 2,
	        "name": "Handmatig terugkerend werk vervangen door een maatwerkapplicatie",
	        "description": "Procesoptimalisatie: een proces dat op losse documenten en overtypen draaide, omgezet naar één applicatie."
	      },
	      {
	        "@type": "ListItem",
	        "position": 3,
	        "name": "Datasecurity aangescherpt en interne infrastructuur gemoderniseerd",
	        "description": "Verouderde, kwetsbare onderdelen vervangen en gevoelige data beter afgeschermd zonder verstoring van de werkprocessen."
	      },
	      {
	        "@type": "ListItem",
	        "position": 4,
	        "name": "Losse systemen aan elkaar gekoppeld zodat data niet meer overgetypt hoeft",
	        "description": "Systeemkoppelingen en API's tussen bestaande applicaties en third-party diensten."
	      }
	    ]
	  }
	}
	</script>
	`)

	return web.Render(c, "cases", fiber.Map{
		"CTATitle":        "Herken je dit vraagstuk?",
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  structuredData + breadcrumbs("/cases", "Cases"),
	})
}

func services(c fiber.Ctx) error {
	title := "Diensten: software op maat, koppelingen en websites | Solvalutions"
	desc := "Software moderniseren, systemen koppelen, werk automatiseren en websites bouwen voor MKB en logistiek. Elke opdracht is afgebakend, met een vaste prijs per fase."
	url := "https://solvalutions.nl"

	structuredData := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "WebPage",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "url": "` + url + `/diensten",
	  "isPartOf": {
	    "@type": "WebSite",
	    "name": "Solvalutions",
	    "url": "` + url + `"
	  },
	  "mainEntity": {
	    "@type": "ItemList",
	    "itemListElement": [
	      {
	        "@type": "ListItem",
	        "position": 1,
	        "item": {
	          "@type": "Service",
	          "name": "Softwaremodernisering",
	          "description": "Verouderde systemen stap voor stap herbouwen naar gestructureerde, geteste techniek. Het draait naast de bestaande omgeving, zodat de operatie niet stilvalt.",
	          "serviceType": "Softwaremodernisering",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          },
	          "areaServed": [
	            { "@type": "AdministrativeArea", "name": "Gelderland" },
	            { "@type": "AdministrativeArea", "name": "Overijssel" },
	            { "@type": "AdministrativeArea", "name": "Utrecht" },
	            { "@type": "AdministrativeArea", "name": "Drenthe" },
	            { "@type": "AdministrativeArea", "name": "Flevoland" },
	            { "@type": "City", "name": "Zutphen" },
	            { "@type": "City", "name": "Apeldoorn" },
	            { "@type": "City", "name": "Deventer" },
	            { "@type": "City", "name": "Arnhem" },
	            { "@type": "City", "name": "Nijmegen" },
	            { "@type": "City", "name": "Enschede" },
	            { "@type": "City", "name": "Zwolle" },
	            { "@type": "City", "name": "Amersfoort" },
	            { "@type": "City", "name": "Utrecht" },
	            { "@type": "Country", "name": "Nederland" }
	          ]
	        }
	      },
	      {
	        "@type": "ListItem",
	        "position": 2,
	        "item": {
	          "@type": "Service",
	          "name": "API-, systeem- en ERP-koppelingen",
	          "description": "Bestaande applicaties, ERP-systemen en diensten van derden aan elkaar koppelen, zodat gegevens automatisch doorstromen. Ook voor nieuwe API's en koppelingen op maat.",
	          "serviceType": "Systeemintegratie",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          }
	        }
	      },
	      {
	        "@type": "ListItem",
	        "position": 3,
	        "item": {
	          "@type": "Service",
	          "name": "Automatisering en maatwerksoftware",
	          "description": "Handmatig terugkerend werk automatiseren met maatwerksoftware en, waar het iets oplevert, AI, zodat het werk wordt weggenomen in plaats van verplaatst.",
	          "serviceType": "Maatwerkontwikkeling",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          }
	        }
	      },
	      {
	        "@type": "ListItem",
	        "position": 4,
	        "item": {
	          "@type": "Service",
	          "name": "Performance en betrouwbaar draaien",
	          "description": "Onderdelen die knellen sneller en voorspelbaarder maken en zorgen dat software blijft draaien met geautomatiseerde uitrol, monitoring en doordachte infrastructuur.",
	          "serviceType": "Performance engineering en DevOps",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          }
	        }
	      },
	      {
	        "@type": "ListItem",
	        "position": 5,
	        "item": {
	          "@type": "Service",
	          "name": "Moderne websites",
	          "description": "Snelle, onderhoudsarme websites die goed gevonden worden en toegankelijk zijn, met een vaste prijs per fase.",
	          "serviceType": "Webontwikkeling",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          }
	        }
	      }
	    ]
	  }
	}
	</script>
	`)

	return web.Render(c, "services", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  structuredData + breadcrumbs("/diensten", "Diensten") + faqStructuredData(servicesFAQ),
		"FAQ":             servicesFAQ,
	})
}

// lastModified holds the date each page's content last changed. Update the
// date for a page when its content changes, so lastmod in the sitemap stays honest.
var lastModified = map[string]string{
	"/":            "2026-11-05",
	"/aanpak":      "2026-11-05",
	"/diensten":    "2026-11-05",
	"/integraties": "2026-11-05",
	"/websites":    "2026-11-05",
	"/conversie":   "2026-11-05",
	"/apps":        "2026-11-05",
	"/dashboards":  "2026-11-05",
	"/werkgebied":  "2026-11-05",
	"/cases":       "2026-11-05",
	"/over":        "2026-11-05",
	"/contact":     "2026-11-05",
}

var sitemapPages = []struct {
	loc      string
	priority string
}{
	{loc: "/", priority: "1.0"},
	{loc: "/aanpak", priority: "0.8"},
	{loc: "/diensten", priority: "0.8"},
	{loc: "/integraties", priority: "0.9"},
	{loc: "/websites", priority: "0.9"},
	{loc: "/conversie", priority: "0.8"},
	{loc: "/apps", priority: "0.8"},
	{loc: "/dashboards", priority: "0.8"},
	{loc: "/werkgebied", priority: "0.7"},
	{loc: "/cases", priority: "0.8"},
	{loc: "/over", priority: "0.8"},
	{loc: "/contact", priority: "0.8"},
}

func sitemap(c fiber.Ctx) error {
	host := c.Hostname()

	// No hreflang alternates: .nl and .com are distinct content (NL commercial vs
	// EN showcase), not translations of each other. Each domain self-canonicals.

	var xml strings.Builder

	xml.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	xml.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")

	for _, page := range sitemapPages {
		xml.WriteString(fmt.Sprintf("\t<url>\n\t\t<loc>https://%s%s</loc>\n\t\t<lastmod>%s</lastmod>\n\t\t<changefreq>monthly</changefreq>\n\t\t<priority>%s</priority>\n", host, page.loc, lastModified[page.loc], page.priority))
		xml.WriteString("\t</url>\n")
	}

	xml.WriteString(`</urlset>`)

	c.Type("xml")
	c.Set("X-Content-Type-Options", "nosniff")

	return c.SendString(xml.String())
}
