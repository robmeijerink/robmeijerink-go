package middleware

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

// TestRateLimiterMiddleware verifies that the rate limiter restricts clients
// to a maximum of 100 requests per minute and returns the custom Dutch error message.
func TestRateLimiterMiddleware(t *testing.T) {
	// Initialize a fresh Fiber app for the test
	app := fiber.New()

	// Register the RateLimiter middleware
	app.Use(RateLimiter())

	// Register a dummy route to hit
	app.Get("/dummy", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Fire 100 valid requests to exhaust the limit
	for i := 0; i < 100; i++ {
		req := httptest.NewRequest("GET", "/dummy", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Request %d should succeed", i+1)
	}

	// Request 101 should trigger the rate limit block
	reqLimit := httptest.NewRequest("GET", "/dummy", nil)
	respLimit, err := app.Test(reqLimit)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusTooManyRequests, respLimit.StatusCode, "Request 101 should be blocked")

	// Assert the custom error message is returned
	bodyBytes, err := io.ReadAll(respLimit.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(bodyBytes), "429 Too Many Requests: Je bent even geblokkeerd")
}
