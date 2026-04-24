package solvalutions

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
	router.Get("/cases", cases)
	router.Get("/contact", contact)

	// en
	router.Get("/about", about)
	router.Get("/approach", approach)
}

func home(c fiber.Ctx) error {
	region, _ := c.Locals("Region").(string)

	title := "Solvalutions | High-Performance B2B Software Engineering"
	desc := "Specialized B2B tech agency for high-performance digital infrastructure and scalable software solutions. Digital Engineering Solutions."
	url := "https://solvalutions.com"

	if region == "nl" {
		title = "Solvalutions | High-Performance B2B Software Engineering"
		desc = "Gespecialiseerde B2B tech agency voor high-performance digitale infrastructuur en schaalbare software oplossingen. Bewezen Digitale Oplossingen."
		url = "https://solvalutions.nl"
	}

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
	region, _ := c.Locals("Region").(string)

	title := "Contact | Your Experienced Tech Partner | Solvalutions"
	desc := "Facing a complex technical challenge or need a scalable application? Contact Solvalutions. Your experienced tech partner for high-performance solutions"
	url := "https://solvalutions.com"

	if region == "nl" {
		title = "Contact | Jouw Ervaren Tech Partner | Solvalutions"
		desc = "Complex technisch vraagstuk of een schaalbare applicatie nodig? Neem contact op met Solvalutions. Dé B2B tech partner voor high-performance oplossingen"
		url = "https://solvalutions.nl"
	}

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
	region, _ := c.Locals("Region").(string)

	title := "B2B Software Engineering Agency | About Solvalutions"
	desc := "Solvalutions is a specialized B2B software engineering agency. We build lean, high-performance digital infrastructure that scales with your business."
	url := "https://solvalutions.com"

	if region == "nl" {
		title = "B2B Software Engineering Agency | Over Solvalutions"
		desc = "Solvalutions is een gespecialiseerde B2B tech partner. Wij bouwen lean, razendsnelle digitale infrastructuur die schaalt met jouw bedrijfsvoering."
		url = "https://solvalutions.nl"
	}

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
	region, _ := c.Locals("Region").(string)

	title := "Approach | High-Performance Tech Architecture | Solvalutions"
	desc := "Discover our lean methodology for building scalable, high-performance Go backends and cloud-native systems. Domain-driven design with brutal efficiency."

	if region == "nl" {
		title = "Werkwijze | High-Performance Tech Architectuur | Solvalutions"
		desc = "Ontdek onze lean werkwijze voor het bouwen van schaalbare, high-performance Go backends en cloud-native systemen. B2B software zonder onnodige overhead."
	}

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
	      "name": "Analysis & Strategy",
	      "text": "We analyze complex business problems to find the most lean and effective technical solution."
	    },
	    {
	      "@type": "HowToStep",
	      "name": "Architecture Design",
	      "text": "Designing solid, scalable foundations using proven technologies like Go and cloud-native principles."
	    },
	    {
	      "@type": "HowToStep",
	      "name": "High-Performance Engineering",
	      "text": "Building the solution with a focus on speed, security, and maintainability without unnecessary overhead."
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
	region, _ := c.Locals("Region").(string)

	title := "Case Studies | Proven Engineering Success | Solvalutions"
	desc := "Explore our portfolio of high-performance B2B software solutions. From cloud-native migrations to custom Go-driven backend architectures."
	url := "https://solvalutions.nl"

	if region == "nl" {
		title = "Cases | Bewezen Digitale Successen | Solvalutions"
		desc = "Bekijk onze portfolio van high-performance B2B softwareoplossingen. Van cloud-native migraties tot schaalbare Go-architecturen."
		url = "https://solvalutions.com"
	}

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
	        "name": "Cloud Infrastructure Migration",
	        "description": "Scalable architecture for a logistics partner."
	      },
	      {
	        "@type": "ListItem",
	        "position": 2,
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

func sitemap(c fiber.Ctx) error {
	host := c.Hostname()
	now := time.Now().Format("2006-01-02")

	pages := []struct {
		loc      string
		priority string
	}{
		{loc: "/", priority: "1.0"},
		{loc: "/about", priority: "0.8"},
		{loc: "/approach", priority: "0.8"},
		{loc: "/cases", priority: "0.8"},
		{loc: "/contact", priority: "0.8"},
	}

	xml := `<?xml version="1.0" encoding="UTF-8"?>`
	xml += `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`

	for _, page := range pages {
		xml += fmt.Sprintf(`
	<url>
		<loc>https://%s%s</loc>
		<lastmod>%s</lastmod>
		<changefreq>monthly</changefreq>
		<priority>%s</priority>
	</url>`, host, page.loc, now, page.priority)
	}

	xml += `</urlset>`

	c.Type("xml")
	return c.SendString(xml)
}
