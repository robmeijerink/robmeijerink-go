package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

// TestRedirectWWWMiddleware verifies that the middleware correctly intercepts
// requests with a 'www.' prefix and issues a 301 redirect to the root domain.
func TestRedirectWWWMiddleware(t *testing.T) {
	// Define test cases covering various hostname and path combinations
	tests := []struct {
		name           string
		requestHost    string
		requestPath    string
		expectedStatus int
		expectedTarget string // Expected Location header for redirects
	}{
		{
			name:           "Redirects www domain to root domain",
			requestHost:    "www.solvalutions.nl",
			requestPath:    "/",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://solvalutions.nl/",
		},
		{
			name:           "Preserves path and query parameters during redirect",
			requestHost:    "www.robmeijerink.nl",
			requestPath:    "/about?theme=dark",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://robmeijerink.nl/about?theme=dark",
		},
		{
			name:           "Ignores root domain (no redirect)",
			requestHost:    "solvalutions.nl",
			requestPath:    "/contact",
			expectedStatus: fiber.StatusOK, // Reaches the dummy handler
			expectedTarget: "",
		},
		{
			name:           "Ignores subdomains other than www",
			requestHost:    "api.solvalutions.nl",
			requestPath:    "/v1/health",
			expectedStatus: fiber.StatusOK,
			expectedTarget: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a fresh Fiber app for the test
			app := fiber.New()

			// Register the middleware being tested
			app.Use(RedirectWWW())

			// Register a dummy route to simulate the next handler in the chain
			app.Get("*", func(c fiber.Ctx) error {
				return c.SendStatus(fiber.StatusOK)
			})

			// Create a mock HTTP request
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			req.Host = tt.requestHost

			// Execute the request through the Fiber app
			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Assert the HTTP status code matches expectations
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// If a redirect is expected, assert the Location header is correct
			if tt.expectedTarget != "" {
				assert.Equal(t, tt.expectedTarget, resp.Header.Get("Location"))
			}
		})
	}
}
