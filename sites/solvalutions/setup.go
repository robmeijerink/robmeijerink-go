package solvalutions

import (
	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

func Setup(router *web.DomainRouter) {
	router.Get("/", home)
}

func home(c fiber.Ctx) error {
	return web.Render(c, "home", fiber.Map{"Title": "Solvalutions | Jouw B2B Tech Partner"})
}
