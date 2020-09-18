package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/romeijerink/robmeijerink-go/routes"
)

func main() {
	// Templates
	engine := django.New("./resources/views", ".html")

	// Fiber instance
    app := fiber.New(fiber.Config{
        Views: engine,
    })

	// Public
	app.Static("/", "./public")

	// Routes
	routes.RegisterWeb(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

// Routes
func home(c *fiber.Ctx) error {
    return c.Render("pages/home", fiber.Map{
              "Title": "New website coming soon 👋!",
           }, "master")
}
