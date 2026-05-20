package middleware

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

// TestSiteContextMiddleware verifies that the correct tenant and context variables
// are injected into the request lifecycle based on the hostname.
func TestSiteContextMiddleware(t *testing.T) {
	// Define test cases covering different hosts and environment flags
	tests := []struct {
		name         string
		requestHost  string
		requestPath  string
		isProdEnv    bool
		expectedSite string
		expectedHost string
		expectedProd bool
		expectedPath string
	}{
		{
			name:         "Default site resolution (robmeijerink)",
			requestHost:  "robmeijerink.nl",
			requestPath:  "/expertise",
			isProdEnv:    true,
			expectedSite: "robmeijerink",
			expectedHost: "robmeijerink.nl",
			expectedProd: true,
			expectedPath: "/expertise",
		},
		{
			name:         "Tenant site resolution (solvalutions)",
			requestHost:  "solvalutions.com",
			requestPath:  "/contact",
			isProdEnv:    false,
			expectedSite: "solvalutions_com",
			expectedHost: "solvalutions.com",
			expectedProd: false,
			expectedPath: "/contact",
		},
		{
			name:         "Tenant site resolution with subdomain (www.solvalutions.nl)",
			requestHost:  "www.solvalutions.nl",
			requestPath:  "/",
			isProdEnv:    true,
			expectedSite: "solvalutions_nl",
			expectedHost: "www.solvalutions.nl",
			expectedProd: true,
			expectedPath: "/",
		},
		{
			name:         "Fallback to default site on unknown host (localhost)",
			requestHost:  "localhost:3000",
			requestPath:  "/api/health",
			isProdEnv:    false,
			expectedSite: "robmeijerink",
			expectedHost: "localhost",
			expectedProd: false,
			expectedPath: "/api/health",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a fresh Fiber app for each test case
			app := fiber.New()

			// Register the middleware with the current test's IsProd flag
			app.Use(RequestContext(tt.isProdEnv))

			// Setup a dummy route to echo the Locals back as JSON
			app.Get("*", func(c fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"Site":          c.Locals("Site"),
					"CanonicalHost": c.Locals("CanonicalHost"),
					"IsProd":        c.Locals("IsProd"),
					"Path":          c.Locals("Path"),
				})
			})

			// Create a mock HTTP request
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			req.Host = tt.requestHost

			// Execute the request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			// Decode the JSON response body to inspect the Locals
			var body map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&body)
			assert.NoError(t, err)

			// Assert that all injected locals match our expectations
			assert.Equal(t, tt.expectedSite, body["Site"])
			assert.Equal(t, tt.expectedHost, body["CanonicalHost"])
			assert.Equal(t, tt.expectedProd, body["IsProd"])
			assert.Equal(t, tt.expectedPath, body["Path"])
		})
	}
}
