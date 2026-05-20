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
	runCrashTest := func(isProd bool) (int, string) {
		server := setupTestServer(isProd)

		server.App.Get("/crash", func(c fiber.Ctx) error {
			return errors.New("FATAL_DB_ERROR: SENSITIVE_PASSWORD_12345_LEAKED")
		})

		req := httptest.NewRequest("GET", "/crash", nil)
		req.Host = "solvalutions.nl"

		resp, err := server.App.Test(req)
		assert.NoError(t, err)

		bodyBytes, _ := io.ReadAll(resp.Body)
		return resp.StatusCode, string(bodyBytes)
	}

	t.Run("Development mode leaks debug info intentionally", func(t *testing.T) {
		status, body := runCrashTest(false)

		assert.Equal(t, fiber.StatusInternalServerError, status)
		assert.Contains(t, body, "SENSITIVE_PASSWORD")
	})

	t.Run("Production mode securely hides debug info", func(t *testing.T) {
		status, body := runCrashTest(true)

		assert.Equal(t, fiber.StatusInternalServerError, status)
		assert.NotContains(t, body, "SENSITIVE_PASSWORD")
		assert.Contains(t, body, "500")
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
// Integration Tests: Tenant Context
// ---------------------------------------------------------
func TestServer_ContextAndIntegration(t *testing.T) {
	server := setupTestServer(false)

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
			name:                  "Resolves primary domain and defaults to Dutch based on .nl",
			requestHost:           "robmeijerink.nl",
			expectedSite:          "robmeijerink",
			expectedCanonicalHost: "robmeijerink.nl",
		},
		{
			name:                  "Resolves tenant domain and defaults to Dutch based on .nl",
			requestHost:           "solvalutions.nl",
			expectedSite:          "solvalutions_nl",
			expectedCanonicalHost: "solvalutions.nl",
		},
		{
			name:                  "Resolves tenant domain and defaults to English based on .com",
			requestHost:           "solvalutions.com",
			expectedSite:          "solvalutions_com",
			expectedCanonicalHost: "solvalutions.com",
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
			assert.Equal(t, "/test-locals", body["Path"])
		})
	}
}

// ---------------------------------------------------------
// Integration Tests: Rate Limiter
// ---------------------------------------------------------
func TestServer_RateLimiterIntegration(t *testing.T) {
	server := setupTestServer(false)

	server.App.Get("/dummy", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	hitRateLimit := false
	var limitBody string

	for i := 0; i < 150; i++ {
		req := httptest.NewRequest("GET", "/dummy", nil)
		req.Host = "solvalutions.com"

		resp, err := server.App.Test(req)
		assert.NoError(t, err)

		if resp.StatusCode == fiber.StatusTooManyRequests {
			hitRateLimit = true
			bodyBytes, _ := io.ReadAll(resp.Body)
			limitBody = string(bodyBytes)
			break
		} else {
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		}
	}

	assert.True(t, hitRateLimit, "The server should return a 429 Too Many Requests status")
	assert.Contains(t, limitBody, "429 Too Many Requests: Je bent even geblokkeerd")
}
