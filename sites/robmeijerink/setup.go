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
        "jobTitle": "Freelance Senior Software Engineer · Full Stack",
        "url": "https://robmeijerink.nl",
        "image": "https://robmeijerink.nl/assets/img/og-rob-meijerink.jpg",
        "knowsAbout": [
          "Full stack development",
          "PHP",
          "Laravel",
          "TypeScript",
          "Vue.js",
          "React",
          "API-koppelingen",
          "Procesautomatisering",
          "Go",
          "Rust",
          "Python",
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
		"Title":           "Freelance Senior Software Engineer · Full Stack | Rob Meijerink",
		"MetaDescription": "Freelance full stack developer inhuren. Bedrijfssoftware, API's en koppelingen in Laravel, Vue, TypeScript, Go en Rust, betrouwbaar live met CI/CD.",
		"Keywords":        "Freelance, ZZP, Full Stack Developer, Software Engineer, inhuren, PHP, Laravel, Vue.js, React, TypeScript, API, koppelingen, automatisering, Go, Golang, Rust, Python, DevOps, CI/CD, Docker, Kubernetes, B2B",
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
	    "jobTitle": "Freelance Senior Software Engineer · Full Stack",
	    "url": "https://robmeijerink.nl",
	    "knowsAbout": [
	      "Full stack development",
	      "PHP",
	      "Laravel",
	      "API-koppelingen",
	      "Procesautomatisering",
	      "Go",
	      "Rust",
	      "Python",
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
		"Title":           "Freelance Full Stack Developer inhuren | Rob Meijerink",
		"MetaDescription": "Senior full stack developer inhuren: backends en API's in PHP/Laravel, Go en Rust, frontends in Vue en React, koppelingen en betrouwbare releases.",
		"Keywords":        "Freelance, Full Stack Developer, Software Engineer, inhuren, Ervaring, CV, ZZP, API, koppelingen, DevOps, Systeem Architectuur, Domain Driven Design, Event Sourcing, Rust, Go, Golang, Python, PHP, Laravel, TypeScript, Vue.js, Kubernetes, Terraform, GitOps, CI/CD",
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
