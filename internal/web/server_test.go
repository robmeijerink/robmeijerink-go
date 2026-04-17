package web

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
	"github.com/stretchr/testify/assert"
)

// setupTestServer is a helper function to generate a clean server instance for each test.
func setupTestServer() *Server {
	cfg := &config.AppConfig{
		IsProd: false, // Use dev mode to bypass manifest loading in tests
		Port:   "3000",
	}
	return NewServer(cfg)
}

// ---------------------------------------------------------
// Middleware Tests: WWW Redirect
// ---------------------------------------------------------
func TestServer_WWWRedirectMiddleware(t *testing.T) {
	server := setupTestServer()

	tests := []struct {
		name           string
		requestHost    string
		requestPath    string
		expectedStatus int
		expectedTarget string
	}{
		{
			name:           "Strips www. and forces HTTPS with exact path",
			requestHost:    "www.solvalutions.nl",
			requestPath:    "/about-us",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://solvalutions.nl/about-us",
		},
		{
			name:           "Preserves query parameters during redirect",
			requestHost:    "www.robmeijerink.nl",
			requestPath:    "/portfolio?tag=go",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://robmeijerink.nl/portfolio?tag=go",
		},
		{
			name:           "Ignores requests without www.",
			requestHost:    "robmeijerink.nl",
			requestPath:    "/",
			expectedStatus: fiber.StatusNotFound, // 404 because no route handles "/" in NewServer natively
			expectedTarget: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			req.Host = tt.requestHost

			resp, err := server.App.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			if tt.expectedTarget != "" {
				assert.Equal(t, tt.expectedTarget, resp.Header.Get("Location"))
			}
		})
	}
}

// ---------------------------------------------------------
// Middleware Tests: Canonical Host & Locals
// ---------------------------------------------------------
func TestServer_CanonicalHostMiddleware(t *testing.T) {
	server := setupTestServer()

	// Inject a temporary test route to extract and verify c.Locals state
	server.App.Get("/test-locals", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Site":          c.Locals("Site"),
			"CanonicalHost": c.Locals("CanonicalHost"),
			"Path":          c.Locals("Path"),
		})
	})

	tests := []struct {
		name                  string
		requestHost           string
		expectedSite          string
		expectedCanonicalHost string
	}{
		{
			name:                  "Resolves primary domain correctly",
			requestHost:           "robmeijerink.nl",
			expectedSite:          "robmeijerink",
			expectedCanonicalHost: "robmeijerink.nl",
		},
		{
			name:                  "Resolves tenant domain correctly",
			requestHost:           "solvalutions.nl",
			expectedSite:          "solvalutions",
			expectedCanonicalHost: "solvalutions.nl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test-locals", nil)
			req.Host = tt.requestHost

			resp, err := server.App.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			// Parse the JSON response back into a map
			var body map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&body)

			assert.Equal(t, tt.expectedSite, body["Site"])
			assert.Equal(t, tt.expectedCanonicalHost, body["CanonicalHost"])
			assert.Equal(t, "/test-locals", body["Path"])
		})
	}
}

// ---------------------------------------------------------
// Middleware Tests: Rate Limiter
// ---------------------------------------------------------
func TestServer_RateLimiter(t *testing.T) {
	server := setupTestServer()

	// Add a dummy route to hit
	server.App.Get("/dummy", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	// The limiter in NewServer is set to Max: 100.
	// We fire 100 valid requests first.
	for i := 0; i < 100; i++ {
		req := httptest.NewRequest("GET", "/dummy", nil)
		resp, err := server.App.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	}

	// Request 101 should be blocked by the rate limiter
	reqLimit := httptest.NewRequest("GET", "/dummy", nil)
	respLimit, err := server.App.Test(reqLimit)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusTooManyRequests, respLimit.StatusCode)

	// Assert the custom Dutch error message
	bodyBytes, _ := io.ReadAll(respLimit.Body)
	assert.Contains(t, string(bodyBytes), "429 Too Many Requests: Je bent even geblokkeerd")
}

// ---------------------------------------------------------
// Router Tests: Static Files & Caching Headers
// ---------------------------------------------------------
func TestServer_StaticFileRouting(t *testing.T) {
	server := setupTestServer()

	tests := []struct {
		name                 string
		path                 string
		expectedCacheControl string
	}{
		{
			name:                 "Asset route applies immutable cache-control",
			path:                 "/assets/css/main.css",
			expectedCacheControl: "public, max-age=31536000, immutable",
		},
		{
			name:                 "Whitelisted root file applies must-revalidate cache-control",
			path:                 "/robots.txt",
			expectedCacheControl: "public, max-age=3600, must-revalidate",
		},
		{
			name:                 "Non-whitelisted root file ignores cache-control and falls through",
			path:                 "/config.json", // Not in the rootFiles map
			expectedCacheControl: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			req.Host = "robmeijerink.nl"

			resp, err := server.App.Test(req)
			assert.NoError(t, err)

			// We check the Cache-Control header. Note: Fiber's c.SendFile might return a 404
			// if the physical file does not exist during testing, but c.Set is executed beforehand,
			// allowing us to verify the routing logic executed correctly.
			assert.Equal(t, tt.expectedCacheControl, resp.Header.Get("Cache-Control"))
		})
	}
}
