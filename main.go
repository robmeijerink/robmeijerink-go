package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", home)

	// Start server
	log.Fatal(app.Listen(80))
}

// Handler
func home(c *fiber.Ctx) {
	c.Send("New website coming soon 👋!")
}
