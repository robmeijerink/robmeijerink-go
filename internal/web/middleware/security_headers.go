package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// InlineJSClassScript is the one inline script in the page layouts. Its hash is
// allowed in the Content-Security-Policy, so any other inline script is blocked.
const InlineJSClassScript = "document.documentElement.classList.add('js')"

// SecurityHeaders adds browser hardening headers. The Content-Security-Policy and
// HSTS are only sent in production, where assets come from the same origin.
func SecurityHeaders(isProd bool) fiber.Handler {
	sum := sha256.Sum256([]byte(InlineJSClassScript))
	scriptHash := "'sha256-" + base64.StdEncoding.EncodeToString(sum[:]) + "'"

	csp := strings.Join([]string{
		"default-src 'self'",
		"script-src 'self' " + scriptHash,
		"style-src 'self' 'unsafe-inline'",
		"img-src 'self' data:",
		"font-src 'self'",
		"connect-src 'self'",
		"object-src 'none'",
		"base-uri 'self'",
		"form-action 'self'",
		"frame-ancestors 'self'",
	}, "; ")

	return func(c fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Set("Permissions-Policy", "camera=(), microphone=(), geolocation=(), payment=(), usb=()")
		c.Set("Cross-Origin-Opener-Policy", "same-origin")

		if isProd {
			c.Set("Content-Security-Policy", csp)
			c.Set("Strict-Transport-Security", "max-age=31536000")
		}

		return c.Next()
	}
}
