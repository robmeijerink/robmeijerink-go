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

		site := tenantFor(host, isProd)

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

// tenantFor maps a hostname to a site. It matches whole domain labels, so a
// look-alike such as "evilsolvalutions.nl" or "solvalutions.nl.evil.com" can
// never select a tenant. In development, ".test" names and localhost work too.
func tenantFor(host string, isProd bool) string {
	host = strings.ToLower(strings.TrimSuffix(host, "."))

	if !isProd {
		if host == "localhost" {
			return "robmeijerink"
		}

		host = strings.TrimSuffix(host, ".test")
	}

	labels := strings.Split(host, ".")
	if len(labels) < 2 {
		return ""
	}

	switch strings.Join(labels[len(labels)-2:], ".") {
	case "solvalutions.nl":
		return "solvalutions_nl"
	case "solvalutions.com":
		return "solvalutions_com"
	}

	if labels[len(labels)-2] == "robmeijerink" {
		return "robmeijerink"
	}

	return ""
}
