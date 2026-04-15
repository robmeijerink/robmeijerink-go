package main

import (
	"log"
	"strings"
	"time"

	// Internal "Engine"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
	"github.com/robmeijerink/robmeijerink-go/internal/web"

	// Domain Specific Sites
	"github.com/robmeijerink/robmeijerink-go/sites/robmeijerink"
	"github.com/robmeijerink/robmeijerink-go/sites/solvalutions"

	// External
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func main() {
	cfg := config.Load()

	engine := html.New("./sites", ".html")
	if !cfg.IsProd {
		engine.Reload(true)
	}

	if cfg.IsProd {
		_ = web.LoadManifest("robmeijerink", "./sites/robmeijerink/public/dist/.vite/manifest.json")
		_ = web.LoadManifest("solvalutions", "./sites/solvalutions/public/dist/.vite/manifest.json")
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendString("429 Too Many Requests: Je bent even geblokkeerd want er waren te veel verzoeken. Neem even pauze.")
		},
	}))

	app.Use(func(c fiber.Ctx) error {
		rawHost := c.Hostname()
		canonicalHost := strings.TrimPrefix(rawHost, "www.")

		site := "robmeijerink"
		if strings.Contains(canonicalHost, "solvalutions") {
			site = "solvalutions"
		}

		c.Locals("Site", site)
		c.Locals("CanonicalHost", canonicalHost)
		c.Locals("IsProd", cfg.IsProd)
		c.Locals("Path", c.Path())

		return c.Next()
	})

	if !cfg.IsProd {
		app.Get("/assets/rob/*", static.New("./sites/robmeijerink/public/dist"))
		app.Get("/assets/sol/*", static.New("./sites/solvalutions/public/dist"))

		app.Get("/img/rob/*", static.New("./sites/robmeijerink/public/img"))
		app.Get("/img/sol/*", static.New("./sites/solvalutions/public/img"))
	}

	robmeijerink.Setup(&web.DomainRouter{
		App:    app,
		Domain: "robmeijerink",
	})

	solvalutions.Setup(&web.DomainRouter{
		App:    app,
		Domain: "solvalutions",
	})

	log.Fatal(app.Listen(":" + cfg.Port))
}
