package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

// RequestContext boots up the environment for the incoming request,
// detecting the tenant (Site) and the market (Region).
// It blocks any traffic from unknown domains or direct IP access.
func RequestContext(isProd bool) fiber.Handler {
	return func(c fiber.Ctx) error {
		host := c.Hostname()

		var site string
		if strings.Contains(host, "solvalutions.nl") {
			site = "solvalutions_nl"
		} else if strings.Contains(host, "solvalutions.com") {
			site = "solvalutions_com"
		} else if strings.Contains(host, "robmeijerink") {
			site = "robmeijerink"
		} else if !isProd && strings.Contains(host, "localhost") {
			// Allow localhost routing during development
			site = "robmeijerink"
		}

		// Abort early if the domain is completely unknown
		if site == "" {
			return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
		}

		// Batch inject all locals
		c.Locals("Site", site)
		c.Locals("CanonicalHost", host)
		c.Locals("IsProd", isProd)
		c.Locals("Path", c.Path())

		return c.Next()
	}
}
