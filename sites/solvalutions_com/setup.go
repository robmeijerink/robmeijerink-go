package solvalutions_com

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
	router.Get("/about", about)
	router.Get("/approach", approach)

	router.Get("/sitemap.xml", sitemap)
}

func home(c fiber.Ctx) error {
	title := "Solvalutions | Digital Solutions for Business Problems"
	desc := "Solvalutions solves business problems with digital solutions: custom software, APIs, AI and cloud infrastructure. Backends in Rust and Go."
	url := "https://solvalutions.com"

	jsonLd := template.HTML(`
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "Organization",
      "name": "Solvalutions",
      "url": "` + url + `",
      "logo": "` + url + `/assets/img/solvalutions-logo.png",
      "image": "` + url + `/assets/img/solvalutions-og-image.png",
      "description": "` + desc + `",
      "founder": {
        "@type": "Person",
        "name": "Rob Meijerink",
        "url": "https://robmeijerink.nl"
      },
      "sameAs": [
        "https://linkedin.com/company/solvalutions",
        "https://github.com/robmeijerink"
      ],
      "address": {
        "@type": "PostalAddress",
        "addressCountry": "NL"
      }
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
	title := "Contact | Tell Us Your Problem | Solvalutions"
	desc := "Got a business problem you want solved with software? Contact Solvalutions and talk directly with an engineer."
	url := "https://solvalutions.com"

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "ContactPage",
	  "name": "Contact Solvalutions",
	  "description": "` + desc + `",
	  "url": "` + url + `/contact",
	  "mainEntity": {
	    "@type": "Organization",
	    "name": "Solvalutions",
	    "url": "` + url + `",
	    "logo": "` + url + `/assets/img/solvalutions-logo.png",
	    "contactPoint": {
	      "@type": "ContactPoint",
	      "telephone": "+31-6-49691374",
	      "contactType": "customer support",
	      "email": "info@solvalutions.nl",
	      "availableLanguage": ["English", "Dutch"]
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

// func services(c fiber.Ctx) error {
// 	region, _ := c.Locals("Region").(string)
// 	currentPath := c.Path()
//
// 	if region == "nl" && currentPath != "/diensten" {
// 		return c.Redirect("/diensten", fiber.StatusMovedPermanently)
// 	} else if region == "en" && currentPath != "/services" {
// 		return c.Redirect("/services", fiber.StatusMovedPermanently)
// 	}
//
// 	title := "Services | Solvalutions"
// 	if region == "nl" {
// 		title = "Diensten | Solvalutions"
// 	}
//
// 	return web.Render(c, "services", fiber.Map{
// 		"Title": title,
// 	})
// }

func about(c fiber.Ctx) error {
	title := "About | Software That Solves Problems | Solvalutions"
	desc := "Solvalutions is a software engineering agency that builds digital solutions for businesses. We solve problems and deliver solutions."
	url := "https://solvalutions.com"

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
      "@context": "https://schema.org",
      "@type": "AboutPage",
      "mainEntity": {
        "@type": "Organization",
        "name": "Solvalutions",
        "url": "` + url + `",
        "logo": "` + url + `/assets/img/solvalutions-logo.png",
        "description": "` + desc + `",
        "founder": {
          "@type": "Person",
          "name": "Rob Meijerink",
          "url": "https://robmeijerink.nl"
        },
        "sameAs": [
          "https://linkedin.com/in/robm89"
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
	title := "Approach | From Problem to Solution | Solvalutions"
	desc := "How we turn business problems into working solutions: understand the domain, build lean with Rust and Go, and run it cloud-native."

	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "HowTo",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "step": [
	    {
	      "@type": "HowToStep",
	      "name": "Understand the problem",
	      "text": "We map your business reality to code and define the domain boundaries before building anything."
	    },
	    {
	      "@type": "HowToStep",
	      "name": "Build the solution",
	      "text": "We pick Rust or Go for the problem and build lean, type-safe APIs and services."
	    },
	    {
	      "@type": "HowToStep",
	      "name": "Run it reliably",
	      "text": "Infrastructure as Code automates deployments so the system is available, scalable and reproducible."
	    }
	  ]
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
	title := "Case Studies | Problems Solved | Solvalutions"
	desc := "Business problems we solved with software, from slow batch imports to cloud migrations, built with Rust and Go."
	url := "https://solvalutions.com"

	structuredData := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "CollectionPage",
	  "name": "` + title + `",
	  "description": "` + desc + `",
	  "url": "` + url + `/cases",
	  "mainEntity": {
	    "@type": "ItemList",
	    "itemListElement": [
	      {
	        "@type": "ListItem",
	        "position": 1,
	        "name": "Data Aggregation Engine in Rust",
	        "description": "Data from many sources merged per request in a concurrent Rust service with predictable latency."
	      },
	      {
	        "@type": "ListItem",
	        "position": 2,
	        "name": "Batch Processing Rewrite in Rust",
	        "description": "A PHP nightly import rewritten in Rust, with identical output and a fraction of the runtime."
	      },
	      {
	        "@type": "ListItem",
	        "position": 3,
	        "name": "Shared Rust Core via FFI and WebAssembly",
	        "description": "Business-critical calculation logic extracted into one Rust library used by PHP and the browser."
	      },
	      {
	        "@type": "ListItem",
	        "position": 4,
	        "name": "High-Throughput API Layer in Rust",
	        "description": "A public API rebuilt in Rust with async I/O and caching for stable tail latency under load."
	      },
	      {
	        "@type": "ListItem",
	        "position": 5,
	        "name": "Cloud Infrastructure Migration",
	        "description": "A fragmented PHP monolith migrated to Go microservices on Kubernetes."
	      },
	      {
	        "@type": "ListItem",
	        "position": 6,
	        "name": "WooCommerce to Go API Refactor",
	        "description": "Transforming a slow legacy store into a high-performance headless commerce engine."
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

// lastModified holds the date each page's content last changed. Update the
// date for a page when its content changes, so lastmod in the sitemap stays honest.
var lastModified = map[string]string{
	"/":         "2026-11-07",
	"/about":    "2026-11-07",
	"/approach": "2026-11-07",
	"/cases":    "2026-11-07",
	"/contact":  "2026-11-07",
}

var sitemapPages = []struct {
	loc      string
	priority string
}{
	{loc: "/", priority: "1.0"},
	{loc: "/about", priority: "0.8"},
	{loc: "/approach", priority: "0.8"},
	{loc: "/cases", priority: "0.8"},
	{loc: "/contact", priority: "0.8"},
}

func sitemap(c fiber.Ctx) error {
	host := c.Hostname()

	// No hreflang alternates: .com (EN showcase) and .nl (NL commercial) are
	// distinct content, not translations of each other. Each domain self-canonicals.

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
