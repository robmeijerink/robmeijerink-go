package robmeijerink

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/workstation-setup", workstationSetup)

	router.Get("/", home)
	router.Get("/expertise", work)
	// router.Get("/blog", blog)

	router.Get("/llms-full.txt", llmsFull)
	router.Get("/sitemap.xml", sitemap)
}

func home(c fiber.Ctx) error {
	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
      "@context": "https://schema.org",
      "@type": "ProfilePage",
      "mainEntity": {
        "@type": "Person",
        "name": "Rob Meijerink",
        "jobTitle": "Freelance Senior Software Developer · Platform & DevOps",
        "url": "https://robmeijerink.nl",
        "image": "https://robmeijerink.nl/assets/img/og-rob-meijerink.jpg",
        "knowsAbout": [
          "Rust",
          "Go",
          "Python",
          "PHP",
          "Laravel",
          "TypeScript",
          "Vue.js",
          "Software architectuur",
          "Microservices",
          "DevOps",
          "Infrastructure as Code",
          "CI/CD",
          "Docker",
          "Kubernetes",
          "Terraform",
          "GitOps"
        ],
        "worksFor": {
          "@type": "Organization",
          "name": "Solvalutions",
          "url": "https://solvalutions.nl"
        },
        "sameAs": [
          "https://linkedin.com/in/robm89",
          "https://github.com/robmeijerink",
          "https://solvalutions.nl",
          "https://solvalutions.com"
        ]
      }
    }
	</script>
	`)

	return web.Render(c, "home", fiber.Map{
		"Title":           "Freelance Software Developer · Platform & DevOps | Rob Meijerink",
		"MetaDescription": "Freelance senior software developer beschikbaar. Cloudplatforms en CI/CD met Kubernetes, Terraform en GitOps, plus Go, Rust, Python en PHP/Laravel.",
		"Keywords":        "Freelance, ZZP, Software Developer, inhuren, Platform Engineering, DevOps, CI/CD, Rust, Rust developer, Go, Golang, Python, PHP, Laravel, TypeScript, Vue.js, Kubernetes, Terraform, GitOps, Cloud, B2B",
		"StructuredData":  jsonLd + faqStructuredData(faqItems),
		"FAQ":             faqItems,
	})
}

func work(c fiber.Ctx) error {
	jsonLd := template.HTML(`
	<script type="application/ld+json">
	{
	  "@context": "https://schema.org",
	  "@type": "ProfilePage",
	  "breadcrumb": {
	    "@type": "BreadcrumbList",
	    "itemListElement": [{
	      "@type": "ListItem",
	      "position": 1,
	      "name": "Home",
	      "item": "https://robmeijerink.nl"
	    },{
	      "@type": "ListItem",
	      "position": 2,
	      "name": "Expertise"
	    }]
	  },
	  "mainEntity": {
	    "@type": "Person",
	    "name": "Rob Meijerink",
	    "jobTitle": "Freelance Senior Software Developer · Platform & DevOps",
	    "url": "https://robmeijerink.nl",
	    "knowsAbout": [
	      "Rust",
	      "Go",
	      "Python",
	      "PHP",
	      "Laravel",
	      "Systeem architectuur",
	      "DevOps",
	      "Docker",
	      "Kubernetes",
	      "Terraform",
	      "GitOps",
	      "CI/CD",
	      "Domain Driven Design",
	      "Event Sourcing",
	      "TypeScript",
	      "Vue.js",
	      "React"
	    ],
	    "worksFor": {
	      "@type": "Organization",
	      "name": "Solvalutions",
	      "url": "https://solvalutions.nl"
	    },
	    "sameAs": [
	      "https://linkedin.com/in/robm89",
	      "https://github.com/robmeijerink",
	      "https://solvalutions.nl",
	      "https://solvalutions.com"
	    ]
	  }
	}
	</script>
	`)

	return web.Render(c, "work", fiber.Map{
		"Title":           "Freelance Platform & DevOps Developer inhuren | Rob Meijerink",
		"MetaDescription": "Freelance Software Developer inhuren: cloudplatforms en CI/CD met Kubernetes, Terraform en GitOps, backends in Go, Rust, Python en PHP/Laravel.",
		"Keywords":        "Freelance, Developer, inhuren, Platform Engineering, Ervaring, CV, ZZP, DevOps, Systeem Architectuur, Domain Driven Design, Event Sourcing, Rust, Go, Golang, Python, PHP, Laravel, TypeScript, Vue.js, Kubernetes, Terraform, GitOps, CI/CD",
		"StructuredData":  jsonLd,
	})
}

// func blog(c fiber.Ctx) error {
// 	return web.Render(c, "blog", fiber.Map{"Title": "Mijn Blog"})
// }

func workstationSetup(c fiber.Ctx) error {
	// c.Set(fiber.HeaderStrictTransportSecurity, "max-age=31536000; includeSubDomains")
	return c.Redirect().Status(fiber.StatusFound).To("https://raw.githubusercontent.com/robmeijerink/workstation-setup/main/workstation-setup.sh")
}

func sitemap(c fiber.Ctx) error {
	host := c.Hostname()
	now := time.Now().Format("2006-01-02")

	pages := []struct {
		loc      string
		priority string
	}{
		{loc: "/", priority: "1.0"},
		{loc: "/expertise", priority: "0.8"},
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
