package middleware

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityHeaders(t *testing.T) {
	for _, isProd := range []bool{true, false} {
		app := fiber.New()
		app.Use(SecurityHeaders(isProd))
		app.Get("/", func(c fiber.Ctx) error { return c.SendString("ok") })

		resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
		require.NoError(t, err)

		assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
		assert.Equal(t, "SAMEORIGIN", resp.Header.Get("X-Frame-Options"))
		assert.NotEmpty(t, resp.Header.Get("Referrer-Policy"))

		if isProd {
			csp := resp.Header.Get("Content-Security-Policy")
			assert.Contains(t, csp, "object-src 'none'")
			assert.Contains(t, csp, "frame-ancestors 'self'")
			assert.NotContains(t, csp, "script-src 'self' 'unsafe-inline'")
			assert.NotEmpty(t, resp.Header.Get("Strict-Transport-Security"))
		} else {
			assert.Empty(t, resp.Header.Get("Content-Security-Policy"))
			assert.Empty(t, resp.Header.Get("Strict-Transport-Security"))
		}
	}
}

// Every inline script in the layouts must match the hash allowed by the CSP,
// otherwise the browser blocks it in production.
func TestInlineScriptsMatchCSPHash(t *testing.T) {
	files, err := filepath.Glob("../../../sites/*/views/*/*.html")
	require.NoError(t, err)
	require.NotEmpty(t, files)

	for _, file := range files {
		raw, err := os.ReadFile(file)
		require.NoError(t, err)

		for _, line := range strings.Split(string(raw), "\n") {
			if i := strings.Index(line, "<script>"); i >= 0 {
				body := line[i+len("<script>"):]
				body = body[:strings.Index(body, "</script>")]
				assert.Equal(t, InlineJSClassScript, body, file)
			}
		}
	}
}
