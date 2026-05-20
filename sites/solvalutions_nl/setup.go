package solvalutions_nl

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
	router.Get("/cases", cases)
	router.Get("/contact", contact)

	router.Get("/diensten", services)
	router.Get("/over", about)
	router.Get("/aanpak", approach)

	router.Get("/sitemap.xml", sitemap)
}

func home(c fiber.Ctx) error {
	title := "Software moderniseren & Development | Solvalutions"
	desc := "Modernisering van trage of vastgelopen PHP- en Go-software voor MKB en maakindustrie in Oost-Nederland. Minder handwerk, minder fouten, vaste prijs per fase."
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
      "email": "rob.meijerink@gmail.com",
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
        { "@type": "AdministrativeArea", "name": "Oost-Nederland" },
        { "@type": "Country", "name": "Nederland" }
      ],
      "knowsAbout": [
        "Legacy modernisering",
        "Go",
        "Fiber",
        "Laravel",
        "PHP",
        "Systeemintegratie",
        "API-ontwikkeling",
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
		"StructuredData":  jsonLd,
	})
}

func contact(c fiber.Ctx) error {
	title := "Contact — Even sparren over je ICT vraagstuk | Solvalutions"
	desc := "Loopt er iets vast in je software? Stuur een mail of bel even — rechtstreeks met Rob, geen accountmanager. Geen verplichtingen."
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
	      "email": "rob@solvalutions.nl",
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
		"StructuredData":  jsonLd,
	})
}

func about(c fiber.Ctx) error {
	title := "Over Rob Meijerink — engineer achter Solvalutions"
	desc := "Tien jaar full-stack developer, nu zelfstandig vanuit Zutphen. Eén engineer, geen tussenlagen — je werkt rechtstreeks met de persoon die het bouwt."
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
        "jobTitle": "Software Engineer",
        "description": "Senior full-stack en backend engineer met ruim 10 jaar ervaring in legacy-modernisering, systeemintegratie en maatwerksoftware.",
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
          "Go",
          "Fiber",
          "Laravel",
          "PHP",
          "Legacy modernisering",
          "Systeemintegratie",
          "API-ontwikkeling",
          "Cloud-infrastructuur",
          "DevOps"
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
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  jsonLd,
	})
}

func approach(c fiber.Ctx) error {
	title := "Aanpak — Eerst begrijpen, dan software ontwikkelen | Solvalutions"
	desc := "Klein beginnen, vaste prijs per fase, naast je systeem bouwen. Geen open eindjes — zo werk ik aan modernisering, koppelingen en maatwerk."
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
		"StructuredData":  jsonLd,
	})
}

func cases(c fiber.Ctx) error {
	title := "Cases — Zakelijke ICT Oplossingen | Solvalutions"
	desc := "Een selectie van trajecten: legacy-modernisering, maatwerk dat handwerk wegnam, security-aanscherping en systeemkoppelingen. Referentie op aanvraag."
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
	        "name": "Vastgelopen kernsysteem stap voor stap herbouwd, zonder dat het bedrijf stillag",
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
		"Title":           title,
		"MetaDescription": desc,
		"StructuredData":  structuredData,
	})
}

func services(c fiber.Ctx) error {
	title := "Diensten — Wat ik voor je oplos | Solvalutions"
	desc := "Legacy-modernisering, systeemkoppelingen en API's, maatwerkapplicaties en performance. Concrete uitkomsten met vaste prijs per fase."
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
	          "name": "Legacy-modernisering",
	          "description": "Verouderde systemen stap voor stap herbouwen naar gestructureerde, geteste techniek — draaiend naast de bestaande omgeving, zodat de operatie niet stillegt.",
	          "serviceType": "Softwaremodernisering",
	          "provider": {
	            "@type": "Organization",
	            "name": "Solvalutions",
	            "url": "` + url + `"
	          },
	          "areaServed": [
	            { "@type": "AdministrativeArea", "name": "Oost-Nederland" },
	            { "@type": "Country", "name": "Nederland" }
	          ]
	        }
	      },
	      {
	        "@type": "ListItem",
	        "position": 2,
	        "item": {
	          "@type": "Service",
	          "name": "Systeemkoppelingen en API's",
	          "description": "Bestaande applicaties en third-party diensten via betrouwbare koppelingen en API's met elkaar laten samenwerken, zodat informatie automatisch doorstroomt.",
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
	          "name": "Procesoptimalisatie en maatwerksoftware",
	          "description": "Handmatig terugkerend werk omzetten naar één maatwerkapplicatie die het werk wegneemt in plaats van verplaatst.",
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
	          "description": "Performance-kritische onderdelen herbouwen (onder andere in Go) en zorgen dat software blijft draaien met geautomatiseerde uitrol, monitoring en doordachte infrastructuur.",
	          "serviceType": "Performance engineering en DevOps",
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
		"StructuredData":  structuredData,
	})
}

func sitemap(c fiber.Ctx) error {
	host := c.Hostname()
	now := time.Now().Format("2006-01-02")

	domainNL := "https://www.solvalutions.nl"
	domainCOM := "https://www.solvalutions.com"

	type translation struct {
		lang string
		url  string
	}

	pages := []struct {
		loc          string
		priority     string
		translations []translation
	}{
		{
			loc:      "/",
			priority: "1.0",
			translations: []translation{
				{lang: "nl-NL", url: domainNL + "/"},
				{lang: "en", url: domainCOM + "/"},
				{lang: "x-default", url: domainCOM + "/"},
			},
		},
		{
			loc:      "/aanpak",
			priority: "0.8",
			translations: []translation{
				{lang: "nl-NL", url: domainNL + "/aanpak"},
				{lang: "en", url: domainCOM + "/approach"},
				{lang: "x-default", url: domainCOM + "/approach"},
			},
		},
		{
			loc:          "/diensten",
			priority:     "0.8",
			translations: nil,
		},
		{
			loc:      "/cases",
			priority: "0.8",
			translations: []translation{
				{lang: "nl-NL", url: domainNL + "/cases"},
				{lang: "en", url: domainCOM + "/cases"},
				{lang: "x-default", url: domainCOM + "/cases"},
			},
		},
		{
			loc:      "/over",
			priority: "0.8",
			translations: []translation{
				{lang: "nl-NL", url: domainNL + "/over"},
				{lang: "en", url: domainCOM + "/about"},
				{lang: "x-default", url: domainCOM + "/about"},
			},
		},
		{
			loc:      "/contact",
			priority: "0.8",
			translations: []translation{
				{lang: "nl-NL", url: domainNL + "/contact"},
				{lang: "en", url: domainCOM + "/contact"},
				{lang: "x-default", url: domainCOM + "/contact"},
			},
		},
	}

	var xml strings.Builder

	xml.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	xml.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">` + "\n")

	for _, page := range pages {
		xml.WriteString(fmt.Sprintf("\t<url>\n\t\t<loc>https://%s%s</loc>\n\t\t<lastmod>%s</lastmod>\n\t\t<changefreq>monthly</changefreq>\n\t\t<priority>%s</priority>\n", host, page.loc, now, page.priority))

		if page.translations != nil {
			for _, t := range page.translations {
				xml.WriteString(fmt.Sprintf("\t\t<xhtml:link rel=\"alternate\" hreflang=\"%s\" href=\"%s\" />\n", t.lang, t.url))
			}
		}

		xml.WriteString("\t</url>\n")
	}

	xml.WriteString(`</urlset>`)

	c.Type("xml")
	c.Set("X-Content-Type-Options", "nosniff")

	return c.SendString(xml.String())
}
