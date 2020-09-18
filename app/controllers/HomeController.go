package web

import (
	"log"

	"github.com/gofiber/fiber/2"
)

func Index(c *fiber.Ctx) {
	 return c.Render("pages/home", fiber.Map{
               "Title": "New website coming soon 👋!",
            }, "master")
}
