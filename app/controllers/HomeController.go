package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
}

func (HomeController) Index(c *fiber.Ctx) error {
    return c.Render("pages/home", fiber.Map{
              "Title": "New website coming soon 👋!",
           }, "master")
}
