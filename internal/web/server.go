package web

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
	"github.com/robmeijerink/robmeijerink-go/internal/web/middleware"
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
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			log.Printf("ERROR [%s]: %v", c.Path(), err)

			isProd, _ := c.Locals("IsProd").(bool)

			if code == fiber.StatusNotFound {
				c.Status(code)
				if err := Render(c, "404", nil); err != nil {
					return c.SendString("404 - Page not found")
				}
				return nil
			}

			if isProd {
				c.Status(code)
				if err := Render(c, "500", nil); err != nil {
					return c.SendString("500 - Internal server error. The site admin has been notified.")
				}
				return nil
			}

			return c.Status(code).SendString(err.Error())
		},
	})

	// Standard Middlewares
	app.Use(compress.New(compress.Config{Level: compress.LevelDefault}))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())

	// Custom Modular Middlewares
	app.Use(middleware.RateLimiter())
	app.Use(middleware.RedirectWWW())
	app.Use(middleware.RequestContext(cfg.IsProd))

	// Routes
	registerFileRoutes(app)

	return &Server{
		App: app,
		Cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.App.Listen(":" + s.Cfg.Port)
}

// registerFileRoutes handles serving static assets and root-level files.
func registerFileRoutes(app *fiber.App) {
	// Asset Router: Handles everything inside the /assets/ directory
	app.Get("/assets/*", func(c fiber.Ctx) error {
		// The Site context is injected by our SiteContext middleware
		site := c.Locals("Site").(string)
		filePath := fmt.Sprintf("./sites/%s/public/assets/%s", site, c.Params("*"))

		// Immutable cache for long-term storage of hashed assets
		c.Set("Cache-Control", "public, max-age=31536000, immutable")
		return c.SendFile(filePath)
	})

	// Root File Router: Safely handles specific files requested at the domain root
	app.Get("/:rootFile", func(c fiber.Ctx) error {
		file := c.Params("rootFile")

		// Whitelist of allowed root files to prevent arbitrary file reading
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

		// If it's not a whitelisted root file, pass to the next handler
		return c.Next()
	})
}
