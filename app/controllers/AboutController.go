package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type AboutController struct {
}

func (AboutController) Index(c *fiber.Ctx) error {
    return c.Render("pages/about", fiber.Map{}, "master")
}
