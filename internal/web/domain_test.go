package web

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestDomainRouter_StrictIsolation(t *testing.T) {
	// 1. Setup Fiber App
	app := fiber.New()

	// 2. Mock the Host/Tenant Middleware reflecting the new Strict Site Matching
	app.Use(func(c fiber.Ctx) error {
		host := c.Hostname()
		var site string

		if host == "solvalutions.nl" || host == "www.solvalutions.nl" {
			site = "solvalutions_nl"
		} else if host == "solvalutions.com" || host == "www.solvalutions.com" {
			site = "solvalutions_com"
		} else {
			site = "robmeijerink"
		}

		c.Locals("Site", site)
		return c.Next()
	})

	// 3. Initialize Domain Routers
	robRouter := &DomainRouter{
		App:    app,
		Domain: "robmeijerink",
	}

	solNlRouter := &DomainRouter{
		App:    app,
		Domain: "solvalutions_nl",
	}

	// 4. Register strictly the routes we want to test
	robRouter.Get("/expertise", func(c fiber.Ctx) error {
		return c.SendString("Rob's Expertise Page")
	})

	solNlRouter.Get("/services", func(c fiber.Ctx) error {
		return c.SendString("Solvalutions NL Services Page")
	})

	// 5. Define exact test scenarios
	tests := []struct {
		name           string
		host           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		// Valid Access
		{
			name:           "Rob can access his own /expertise route",
			host:           "robmeijerink.nl",
			path:           "/expertise",
			expectedStatus: fiber.StatusOK,
			expectedBody:   "Rob's Expertise Page",
		},
		{
			name:           "Solvalutions NL can access its own /services route",
			host:           "solvalutions.nl",
			path:           "/services",
			expectedStatus: fiber.StatusOK,
			expectedBody:   "Solvalutions NL Services Page",
		},

		// Unauthorized Cross-Tenant Access
		{
			name:           "Solvalutions NL CANNOT access Rob's /expertise route",
			host:           "solvalutions.nl",
			path:           "/expertise",
			expectedStatus: fiber.StatusNotFound, // Fiber returns 404 because c.Next() falls through
			expectedBody:   "Not Found",
		},
		{
			name:           "Rob CANNOT access Solvalutions NL's /services route",
			host:           "robmeijerink.nl",
			path:           "/services",
			expectedStatus: fiber.StatusNotFound,
			expectedBody:   "Not Found",
		},
	}

	// 6. Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			req.Host = tt.host // Crucial: sets the domain for the middleware

			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			bodyBytes, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(bodyBytes), tt.expectedBody)
		})
	}
}

// ---------------------------------------------------------
// Mock Views Engine
// ---------------------------------------------------------
type MockViews struct {
	Template       string
	Layout         string
	Data           fiber.Map
	FailOnTemplate string // Added to simulate a missing template file
}

func (m *MockViews) Load() error { return nil }

func (m *MockViews) Render(w io.Writer, template string, data any, layouts ...string) error {
	// Simulate Fiber returning an error when a file is missing
	if m.FailOnTemplate != "" && template == m.FailOnTemplate {
		return errors.New("template not found on disk")
	}

	m.Template = template
	if len(layouts) > 0 {
		m.Layout = layouts[0]
	}

	if mapData, ok := data.(fiber.Map); ok {
		m.Data = mapData
	}

	_, _ = w.Write([]byte("mock HTML content"))
	return nil
}

// ---------------------------------------------------------
// Render Function Tests
// ---------------------------------------------------------

func TestRender_FlatStructure_Success(t *testing.T) {
	mockEngine := &MockViews{}
	app := fiber.New(fiber.Config{Views: mockEngine})

	// To effectively test the Render function without relying on the actual config.SiteRegistry
	// which is parsed in runtime, we override the strict check in a test environment, or we
	// mock the local config map. Assuming config.SiteRegistry["solvalutions_nl"] is initialized in init() or setup.

	// Note: You must ensure your test runner initializes config.SiteRegistry["solvalutions_nl"]
	// for this handler not to panic/throw 500 based on the current domain.go logic.

	app.Get("/test-flat", func(c fiber.Ctx) error {
		// Mock locals for a flat site visitor
		c.Locals("Site", "solvalutions_nl")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "solvalutions.nl")
		c.Locals("Path", "/test-flat")

		customData := fiber.Map{"PageTitle": "Engineering Backbone"}
		return Render(c, "pages/home", customData)
	})

	req := httptest.NewRequest("GET", "/test-flat", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Assert it used the flat, isolated structure
	assert.Equal(t, "solvalutions_nl/views/pages/home", mockEngine.Template)
	// Assert it loaded the root layout file for this specific site
	assert.Equal(t, "solvalutions_nl/views/layouts/master", mockEngine.Layout)
	assert.Equal(t, "Engineering Backbone", mockEngine.Data["PageTitle"])
}
