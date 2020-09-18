package web

import (
	"log"

	"github.com/gofiber/fiber/2"
)

func Index(c *fiber.Ctx) {
	 return c.Render("pages/about", fiber.Map{}, "master")
}
