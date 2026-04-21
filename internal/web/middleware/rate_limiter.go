package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// RateLimiter configures and returns the Fiber rate limiting middleware.
// It restricts clients to a maximum of 100 requests per minute.
func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			// Custom error message when the rate limit is exceeded
			return c.Status(fiber.StatusTooManyRequests).SendString("429 Too Many Requests: Je bent even geblokkeerd want er waren te veel verzoeken. Neem even pauze.")
		},
	})
}
