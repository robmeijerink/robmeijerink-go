package solvalutions

import (
	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
	router.Get("/contact", contact)

	// en
	router.Get("/approach", approach)
}

func home(c fiber.Ctx) error {
	region, _ := c.Locals("Region").(string)

	title := "Solvalutions | Your B2B Tech Partner"
	desc := "We engineer high-performance backend systems."

	if region == "nl" {
		title = "Solvalutions | Jouw B2B Tech Partner"
		desc = "Wij bouwen high-performance backend systemen."
	}

	return web.Render(c, "home", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
	})
}

func contact(c fiber.Ctx) error {
	region, _ := c.Locals("Region").(string)

	title := "Contact | Your Experienced Tech Partner | Solvalutions"
	desc := "Facing a complex technical challenge or need a scalable application? Contact Solvalutions. Your experienced tech partner for high-performance solutions"

	if region == "nl" {
		title = "Contact | Jouw Ervaren Tech Partner | Solvalutions"
		desc = "Complex technisch vraagstuk of een schaalbare applicatie nodig? Neem contact op met Solvalutions. Dé B2B tech partner voor high-performance oplossingen"
	}

	return web.Render(c, "contact", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
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

func approach(c fiber.Ctx) error {
	region, _ := c.Locals("Region").(string)

	title := "Approach | High-Performance Tech Architecture | Solvalutions"
	desc := "Discover our lean methodology for building scalable, high-performance Go backends and cloud-native systems. Domain-driven design with brutal efficiency."

	if region == "nl" {
		title = "Werkwijze | High-Performance Tech Architectuur | Solvalutions"
		desc = "Ontdek onze lean werkwijze voor het bouwen van schaalbare, high-performance Go backends en cloud-native systemen. B2B software zonder onnodige overhead."
	}

	return web.Render(c, "approach", fiber.Map{
		"Title":           title,
		"MetaDescription": desc,
	})
}
