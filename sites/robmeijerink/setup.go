package robmeijerink

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/workstation-setup", workstationSetup)

	router.Get("/", home)
	router.Get("/expertise", work)
	// router.Get("/blog", blog)

	router.Get("/sitemap.xml", sitemap)
}

func home(c fiber.Ctx) error {
	return web.Render(c, "home", fiber.Map{
		"Title":           "Rob Meijerink | Software Developer & B2B Tech Partner",
		"MetaDescription": "Beschikbaar voor freelance projecten (ZZP). Gespecialiseerd in schaalbare backends (Go/PHP), cloud-architectuur en DevOps automatisering (Kubernetes/AWS).",
		"Keywords":        "Freelance, ZZP, Software Engineer, inhuren, DevOps, Backend Developer, Go, Golang, PHP, Laravel, Kubernetes, Systeem Architectuur, B2B",
	})
}

func work(c fiber.Ctx) error {
	return web.Render(c, "work", fiber.Map{
		"Title":           "Rob Meijerink | Go, Laravel, PHP, JavaScript & Cloud Development",
		"MetaDescription": "Mijn expertise als Freelance Software Developer inhuren (ZZP). Ervaring in schaalbare software (Go/PHP/Laravel) Domain-Driven Design en DevOps (Kubernetes)",
		"Keywords":        "Developer, inhuren, Backend Development, Ervaring, CV, ZZP, DevOps Engineering, Systeem Architectuur, Domain Driven Design, Event Sourcing, Go, Golang, PHP, Laravel, Kubernetes, CI/CD",
	})
}

// func blog(c fiber.Ctx) error {
// 	return web.Render(c, "blog", fiber.Map{"Title": "Mijn Blog"})
// }

func workstationSetup(c fiber.Ctx) error {
	c.Set(fiber.HeaderStrictTransportSecurity, "max-age=31536000; includeSubDomains")
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
