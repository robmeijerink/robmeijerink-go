package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

// RedirectWWW strips the 'www.' prefix from the hostname and redirects to the root domain
func RedirectWWW() fiber.Handler {
	return func(c fiber.Ctx) error {
		hostname := c.Hostname()
		if strings.HasPrefix(hostname, "www.") {
			newHost := strings.TrimPrefix(hostname, "www.")
			target := "https://" + newHost + c.OriginalURL()
			return c.Redirect().Status(fiber.StatusMovedPermanently).To(target)
		}
		return c.Next()
	}
}
