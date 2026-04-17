package web

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
)

type Server struct {
	App *fiber.App
	Cfg *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	engine := html.New("./sites", ".html")
	if !cfg.IsProd {
		engine.Reload(true)
	}

	if cfg.IsProd {
		_ = LoadManifest("robmeijerink", "./sites/robmeijerink/public/assets/dist/.vite/manifest.json")
		_ = LoadManifest("solvalutions", "./sites/solvalutions/public/assets/dist/.vite/manifest.json")
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// --- Middlewares ---
	app.Use(compress.New(compress.Config{Level: compress.LevelDefault}))
	app.Use(recover.New())
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendString("429 Too Many Requests: Je bent even geblokkeerd want er waren te veel verzoeken. Neem even pauze.")
		},
	}))

	// WWW Redirect Middleware
	app.Use(func(c fiber.Ctx) error {
		hostname := c.Hostname()
		if strings.HasPrefix(hostname, "www.") {
			newHost := strings.TrimPrefix(hostname, "www.")
			target := "https://" + newHost + c.OriginalURL()
			return c.Redirect().Status(fiber.StatusMovedPermanently).To(target)
		}
		return c.Next()
	})

	// Canonical Host Middleware
	app.Use(func(c fiber.Ctx) error {
		host := c.Hostname()
		site := "robmeijerink"
		if strings.Contains(host, "solvalutions") {
			site = "solvalutions"
		}

		c.Locals("Site", site)
		c.Locals("CanonicalHost", host)
		c.Locals("IsProd", cfg.IsProd)
		c.Locals("Path", c.Path())
		return c.Next()
	})

	// Asset Router
	app.Get("/assets/*", func(c fiber.Ctx) error {
		site := c.Locals("Site").(string)
		filePath := fmt.Sprintf("./sites/%s/public/assets/%s", site, c.Params("*"))
		c.Set("Cache-Control", "public, max-age=31536000, immutable")
		return c.SendFile(filePath)
	})

	// Root File Router
	app.Get("/:rootFile", func(c fiber.Ctx) error {
		file := c.Params("rootFile")
		rootFiles := map[string]bool{
			"favicon.ico":                  true,
			"favicon.svg":                  true,
			"favicon-96x96.png":            true,
			"apple-touch-icon.png":         true,
			"site.webmanifest":             true,
			"robots.txt":                   true,
			"web-app-manifest-192x192.png": true,
			"web-app-manifest-512x512.png": true,
		}

		if rootFiles[file] {
			site := c.Locals("Site").(string)
			c.Set("Cache-Control", "public, max-age=3600, must-revalidate")
			return c.SendFile(fmt.Sprintf("./sites/%s/public/%s", site, file))
		}
		return c.Next()
	})

	return &Server{
		App: app,
		Cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.App.Listen(":" + s.Cfg.Port)
}
