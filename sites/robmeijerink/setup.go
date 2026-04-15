package robmeijerink

import (
	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
	router.Get("/expertise", work)
	// router.Get("/blog", blog)
}

func home(c fiber.Ctx) error {
	return web.Render(c, "home", fiber.Map{
		"Title":           "Rob Meijerink | Software Developer & B2B Tech Partner",
		"MetaDescription": "Beschikbaar voor freelance projecten (ZZP). Gespecialiseerd in schaalbare backends (Go/PHP), cloud-architectuur en DevOps automatisering (Kubernetes/AWS).",
		"Keywords":        "Freelance, ZZP, Software Engineer, DevOps, Backend Developer, Go, Golang, PHP, Laravel, Kubernetes, Systeem Architectuur, B2B",
	})
}

func work(c fiber.Ctx) error {
	return web.Render(c, "work", fiber.Map{
		"Title":           "Rob Meijerink | Go, Laravel, PHP, JavaScript & Cloud Development",
		"MetaDescription": "Ontdek mijn expertise als Freelance Software Architect (ZZP). Ervaring in schaalbare software (Go/PHP/Laravel), Domain-Driven Design en DevOps (Kubernetes)",
		"Keywords":        "Backend Development, Ervaring, CV, ZZP, DevOps Engineering, Systeem Architectuur, Domain Driven Design, Event Sourcing, Go, Golang, PHP, Laravel, Kubernetes, CI/CD",
	})
}

// func blog(c fiber.Ctx) error {
// 	return web.Render(c, "blog", fiber.Map{"Title": "Mijn Blog"})
// }
