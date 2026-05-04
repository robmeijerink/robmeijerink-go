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

type MockViews struct{}

func (m *MockViews) Load() error { return nil }

func (m *MockViews) Render(w io.Writer, template string, data any, layouts ...string) error {
	_, _ = w.Write([]byte("mocked view: " + template))
	return nil
}

func TestSetupRoutes(t *testing.T) {
	app := fiber.New(fiber.Config{
		Views: &MockViews{},
	})

	app.Use(func(c fiber.Ctx) error {
		c.Locals("Site", "robmeijerink")
		c.Locals("Region", "nl")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "robmeijerink.nl")
		c.Locals("Path", c.Path())
		return c.Next()
	})

	router := &web.DomainRouter{
		App:    app,
		Domain: "robmeijerink",
	}

	Setup(router)

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
			name:           "Workstation setup redirects to GitHub script",
			method:         "GET",
			path:           "/workstation-setup",
			host:           "robmeijerink.nl",
			expectedStatus: fiber.StatusFound, // 302 Redirect
			checkResponse: func(t *testing.T, resp *http.Response) {
				expectedURL := "https://raw.githubusercontent.com/robmeijerink/workstation-setup/main/workstation-setup.sh"
				assert.Equal(t, expectedURL, resp.Header.Get("Location"))
			},
		},
		{
			name:           "Sitemap dynamically generates valid XML",
			method:         "GET",
			path:           "/sitemap.xml",
			host:           "robmeijerink.nl",
			expectedStatus: fiber.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, "application/xml; charset=utf-8", resp.Header.Get("Content-Type"))

				bodyBytes, _ := io.ReadAll(resp.Body)
				bodyString := string(bodyBytes)

				assert.Contains(t, bodyString, `<?xml version="1.0" encoding="UTF-8"?>`)
				assert.Contains(t, bodyString, "<loc>https://robmeijerink.nl/</loc>")

				today := time.Now().Format("2006-01-02")
				assert.Contains(t, bodyString, "<lastmod>"+today+"</lastmod>")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			req.Host = tt.host

			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.checkResponse != nil {
				tt.checkResponse(t, resp)
			}
		})
	}
}
