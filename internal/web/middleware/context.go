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

		// /over/ and /over are the same page: send crawlers to one address.
		if path := c.Path(); len(path) > 1 && strings.HasSuffix(path, "/") && (c.Method() == fiber.MethodGet || c.Method() == fiber.MethodHead) {
			target := strings.TrimRight(path, "/")
			if target == "" {
				target = "/"
			}

			if query := strings.IndexByte(c.OriginalURL(), '?'); query >= 0 {
				target += c.OriginalURL()[query:]
			}

			return c.Redirect().Status(fiber.StatusMovedPermanently).To(target)
		}

		// Batch inject all locals. The canonical host never carries "www.",
		// so www and non-www do not compete as two versions of the same page.
		c.Locals("Site", site)
		c.Locals("CanonicalHost", strings.TrimPrefix(host, "www."))
		c.Locals("IsProd", isProd)
		c.Locals("Path", c.Path())

		return c.Next()
	}
}
