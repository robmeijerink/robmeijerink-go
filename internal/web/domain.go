package web

import (
	"github.com/gofiber/fiber/v3"
)

type DomainRouter struct {
	App    *fiber.App
	Domain string
}

func (r *DomainRouter) Get(path string, handler fiber.Handler) {
	r.App.Get(path, func(c fiber.Ctx) error {
		if c.Locals("Site") != r.Domain {
			return c.Next()
		}
		return handler(c)
	})
}

func Render(c fiber.Ctx, template string, data fiber.Map) error {
	site := c.Locals("Site").(string)
	isProd := c.Locals("IsProd").(bool)
	canonicalHost := c.Locals("CanonicalHost").(string)
	path := c.Locals("Path").(string)

	fullData := fiber.Map{
		"Site":          site,
		"Vite":          ViteHelper{Site: site},
		"IsProd":        isProd,
		"CanonicalHost": canonicalHost,
		"Path":          path,
	}

	layoutPath := site + "/views/layouts/master"
	templatePath := site + "/views/" + template

	return c.Render(templatePath, fullData, layoutPath)
}
