package robmeijerink

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
	"github.com/stretchr/testify/assert"
)

// MockViews is a simple mock template engine to intercept c.Render calls
// and prevent Fiber from looking for actual physical HTML files on disk.
type MockViews struct{}

func (m *MockViews) Load() error { return nil }

func (m *MockViews) Render(w io.Writer, template string, data any, layouts ...string) error {
	// Write the template name back so we can assert the correct view was called
	_, _ = w.Write([]byte("mocked view: " + template))
	return nil
}

func TestSetupRoutes(t *testing.T) {
	// 1. Setup a clean Fiber instance with our Mock View Engine
	app := fiber.New(fiber.Config{
		Views: &MockViews{},
	})

	// 2. Mock the global middleware to inject required Locals.
	// Without this, web.Render would panic because of unhandled type assertions
	// like `c.Locals("Site").(string)`.
	app.Use(func(c fiber.Ctx) error {
		c.Locals("Site", "robmeijerink")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "robmeijerink.nl")
		c.Locals("Path", c.Path())
		return c.Next()
	})

	// 3. Initialize the DomainRouter
	router := &web.DomainRouter{
		App:    app,
		Domain: "robmeijerink",
	}

	// 4. Register the routes
	Setup(router)

	// 5. Define comprehensive test scenarios
	tests := []struct {
		name           string
		method         string
		path           string
		host           string
		expectedStatus int
		checkResponse  func(t *testing.T, resp *http.Response)
	}{
		{
			name:           "Home route renders the correct view successfully",
			method:         "GET",
			path:           "/",
			host:           "robmeijerink.nl",
			expectedStatus: fiber.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				bodyBytes, _ := io.ReadAll(resp.Body)
				// The path includes the tenant name because web.Render prepends it
				assert.Contains(t, string(bodyBytes), "mocked view: robmeijerink/views/home")
			},
		},
		{
			name:           "Expertise route renders the correct view successfully",
			method:         "GET",
			path:           "/expertise",
			host:           "robmeijerink.nl",
			expectedStatus: fiber.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				bodyBytes, _ := io.ReadAll(resp.Body)
				assert.Contains(t, string(bodyBytes), "mocked view: robmeijerink/views/work")
			},
		},
		{
			name:           "Workstation setup applies HSTS header and redirects correctly",
			method:         "GET",
			path:           "/workstation-setup",
			host:           "robmeijerink.nl",
			expectedStatus: fiber.StatusFound, // 302 Redirect
			checkResponse: func(t *testing.T, resp *http.Response) {
				// Assert Redirect Location
				expectedURL := "https://raw.githubusercontent.com/robmeijerink/workstation-setup/main/workstation-setup.sh"
				assert.Equal(t, expectedURL, resp.Header.Get("Location"))

				// Assert Security Header
				assert.Equal(t, "max-age=31536000; includeSubDomains", resp.Header.Get("Strict-Transport-Security"))
			},
		},
		{
			name:           "Sitemap dynamically generates valid XML based on the request host",
			method:         "GET",
			path:           "/sitemap.xml",
			host:           "robmeijerink.nl", // Simulating the request coming from the primary domain
			expectedStatus: fiber.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				// Assert Content-Type
				assert.Equal(t, "application/xml; charset=utf-8", resp.Header.Get("Content-Type"))

				bodyBytes, _ := io.ReadAll(resp.Body)
				bodyString := string(bodyBytes)

				// Assert strict XML formatting
				assert.Contains(t, bodyString, `<?xml version="1.0" encoding="UTF-8"?>`)
				assert.Contains(t, bodyString, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

				// Assert dynamic URLs are generated using the provided request Host
				assert.Contains(t, bodyString, "<loc>https://robmeijerink.nl/</loc>")
				assert.Contains(t, bodyString, "<loc>https://robmeijerink.nl/expertise</loc>")

				// Assert dynamic dates
				today := time.Now().Format("2006-01-02")
				assert.Contains(t, bodyString, "<lastmod>"+today+"</lastmod>")
			},
		},
	}

	// 6. Execute all test scenarios
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			req.Host = tt.host // Essential for the sitemap test and DomainRouter

			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Run specific assertions on headers and body content if defined
			if tt.checkResponse != nil {
				tt.checkResponse(t, resp)
			}
		})
	}
}
