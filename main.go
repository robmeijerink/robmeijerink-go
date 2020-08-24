package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Public
	app.Static("/", "./public")

	// Routes
	app.Get("/", home)

	// Start server
	log.Fatal(app.Listen(8080))
}

// Handler
func home(c *fiber.Ctx) {
	c.Send("New website coming soon 👋!")
}
