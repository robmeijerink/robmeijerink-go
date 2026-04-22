package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
	"github.com/stretchr/testify/assert"
)

// setupTestServer generates a clean server instance for integration tests.
// Added 'isProd' parameter so we can test production-specific security behaviors.
func setupTestServer(isProd bool) *Server {
	cfg := &config.AppConfig{
		IsProd: isProd,
		Port:   "3000",
	}
	return NewServer(cfg)
}

// ---------------------------------------------------------
// Integration Tests: Error Handling (Security & Leak Prevention)
// ---------------------------------------------------------
func TestServer_ErrorHandlingIntegration(t *testing.T) {
	// Helper function to spin up a server, inject a crashing route, and return the response
	runCrashTest := func(isProd bool) (int, string) {
		server := setupTestServer(isProd)

		// Inject a temporary route that simulates a fatal internal error containing highly sensitive data
		server.App.Get("/crash", func(c fiber.Ctx) error {
			return errors.New("FATAL_DB_ERROR: SENSITIVE_PASSWORD_12345_LEAKED")
		})

		req := httptest.NewRequest("GET", "/crash", nil)
		req.Host = "solvalutions.nl" // Required to pass the SiteContext middleware

		resp, err := server.App.Test(req)
		assert.NoError(t, err)

		bodyBytes, _ := io.ReadAll(resp.Body)
		return resp.StatusCode, string(bodyBytes)
	}

	t.Run("Development mode leaks debug info intentionally", func(t *testing.T) {
		status, body := runCrashTest(false)

		assert.Equal(t, fiber.StatusInternalServerError, status)
		assert.Contains(t, body, "SENSITIVE_PASSWORD", "In dev mode, the raw error should be visible to the developer")
	})

	t.Run("Production mode securely hides debug info", func(t *testing.T) {
		status, body := runCrashTest(true)

		assert.Equal(t, fiber.StatusInternalServerError, status)
		// SECURITY CHECK: Ensure the sensitive data is absolutely not in the output
		assert.NotContains(t, body, "SENSITIVE_PASSWORD", "CRITICAL SECURITY FAILURE: Production mode leaked sensitive data!")
		// Check if the generic fallback is presented instead
		assert.Contains(t, body, "500", "Production mode should show the safe fallback message or view")
	})

	t.Run("404 Not Found returns safe string or view", func(t *testing.T) {
		server := setupTestServer(true)

		req := httptest.NewRequest("GET", "/this-route-does-not-exist-at-all", nil)
		req.Host = "solvalutions.nl"

		resp, err := server.App.Test(req)
		assert.NoError(t, err)

		bodyBytes, _ := io.ReadAll(resp.Body)
		body := string(bodyBytes)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		assert.Contains(t, body, "404")
	})
}

// ---------------------------------------------------------
// Integration Tests: WWW Redirect
// ---------------------------------------------------------
func TestServer_WWWRedirectIntegration(t *testing.T) {
	server := setupTestServer(false)

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
			requestPath:    "/contact",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://solvalutions.nl/contact",
		},
		{
			name:           "Preserves query parameters during redirect",
			requestHost:    "www.robmeijerink.nl",
			requestPath:    "/expertise?tag=test",
			expectedStatus: fiber.StatusMovedPermanently,
			expectedTarget: "https://robmeijerink.nl/expertise?tag=test",
		},
		{
			name:           "Ignores requests without www.",
			requestHost:    "robmeijerink.nl",
			requestPath:    "/",
			expectedStatus: fiber.StatusNotFound, // 404 because no native route handles "/"
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
// Integration Tests: Tenant Context & Region Detection
// ---------------------------------------------------------
func TestServer_ContextAndRegionIntegration(t *testing.T) {
	server := setupTestServer(false)

	// Inject a temporary route to extract Locals injected by SiteContext and RegionMiddleware
	server.App.Get("/test-locals", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Site":          c.Locals("Site"),
			"CanonicalHost": c.Locals("CanonicalHost"),
			"Region":        c.Locals("Region"),
			"Path":          c.Locals("Path"),
		})
	})

	tests := []struct {
		name                  string
		requestHost           string
		expectedSite          string
		expectedCanonicalHost string
		expectedRegion        string
	}{
		{
			name:                  "Resolves primary domain and defaults to Dutch based on .nl",
			requestHost:           "robmeijerink.nl",
			expectedSite:          "robmeijerink",
			expectedCanonicalHost: "robmeijerink.nl",
			expectedRegion:        "nl",
		},
		{
			name:                  "Resolves tenant domain and defaults to Dutch based on .nl",
			requestHost:           "solvalutions.nl",
			expectedSite:          "solvalutions",
			expectedCanonicalHost: "solvalutions.nl",
			expectedRegion:        "nl",
		},
		{
			name:                  "Resolves tenant domain and defaults to English based on .com",
			requestHost:           "solvalutions.com",
			expectedSite:          "solvalutions",
			expectedCanonicalHost: "solvalutions.com",
			expectedRegion:        "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test-locals", nil)
			req.Host = tt.requestHost

			resp, err := server.App.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			var body map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedSite, body["Site"])
			assert.Equal(t, tt.expectedCanonicalHost, body["CanonicalHost"])
			assert.Equal(t, tt.expectedRegion, body["Region"])
			assert.Equal(t, "/test-locals", body["Path"])
		})
	}
}

// ---------------------------------------------------------
// Integration Tests: Rate Limiter
// ---------------------------------------------------------
func TestServer_RateLimiterIntegration(t *testing.T) {
	server := setupTestServer(false)

	// Add a dummy route to hit
	server.App.Get("/dummy", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	hitRateLimit := false
	var limitBody string

	// Fire up to 100 valid requests. We expect a 429 somewhere around request 31.
	for i := 0; i < 150; i++ {
		req := httptest.NewRequest("GET", "/dummy", nil)
		// Crucial: set the host so our RequestContext middleware allows it through!
		req.Host = "solvalutions.com"

		resp, err := server.App.Test(req)
		assert.NoError(t, err)

		if resp.StatusCode == fiber.StatusTooManyRequests {
			// We successfully hit the rate limit!
			hitRateLimit = true

			// Read the body to check the custom message later
			bodyBytes, _ := io.ReadAll(resp.Body)
			limitBody = string(bodyBytes)

			// Stop spamming requests, the test has successfully triggered the limiter
			break
		} else {
			// If it is not a 429, it MUST be a 200 OK.
			// (If it returns a 404, the Host routing failed).
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		}
	}

	// 1. Assert that the rate limiter actually kicked in during our loop
	assert.True(t, hitRateLimit, "The server should return a 429 Too Many Requests status")

	// 2. Assert the custom error message was returned
	assert.Contains(t, limitBody, "429 Too Many Requests: Je bent even geblokkeerd")
}

// ---------------------------------------------------------
// Integration Tests: Static File Routes & Headers
// ---------------------------------------------------------
func TestServer_StaticFileRoutingIntegration(t *testing.T) {
	server := setupTestServer(false)

	tests := []struct {
		name                 string
		path                 string
		expectedCacheControl string
	}{
		{
			name:                 "Whitelisted root file applies must-revalidate cache-control",
			path:                 "/robots.txt",
			expectedCacheControl: "public, max-age=3600, must-revalidate",
		},
		{
			name:                 "Non-whitelisted root file ignores cache-control and falls through",
			path:                 "/config.json", // Not in the allowed map
			expectedCacheControl: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			req.Host = "robmeijerink.nl" // Provide a valid host for the SiteContext middleware

			resp, err := server.App.Test(req)
			assert.NoError(t, err)

			// Validate Cache-Control header injection by the routing logic
			assert.Equal(t, tt.expectedCacheControl, resp.Header.Get("Cache-Control"))
		})
	}
}
