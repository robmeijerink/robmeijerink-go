package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/django"
)

func main() {
	// Templates
	engine := django.New("./resources/views", ".html")

	// Fiber instance
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	// Public
	app.Static("/", "./public")

	// Routes
	app.Get("/", home)

	// Start server
	log.Fatal(app.Listen(8080))
}

// Routes

func home(c *fiber.Ctx) {
	c.Render("pages/home", fiber.Map{
		"Title": "New website coming soon 👋!",
	}, "master")
}
