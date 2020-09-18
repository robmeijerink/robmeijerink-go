package routes

import (
	"github.com/gofiber/fiber"
	"log"

	HomeController "github.com/romeijerink/robmeijerink-go/app/controllers/HomeController.go"
	AboutController "github.com/romeijerink/robmeijerink-go/app/controllers/AboutController.go"
)

func RegisterWeb(app *fiber.App) {
	// Homepage
	app.Get("/", HomeController.Index)
    app.Get("/about", AboutController.Index())

    // Page not found
    app.Use(func (c *fiber.Ctx) error {
        c.SendStatus(404)
        return c.Render("errors/404", fiber.Map{}, "master")
    })
}
