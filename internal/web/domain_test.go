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

	// 2. Mock the Host/Tenant Middleware exactly as it works in production now
	app.Use(func(c fiber.Ctx) error {
		host := c.Hostname()
		site := "robmeijerink"
		region := "en"

		if host == "solvalutions.nl" || host == "www.solvalutions.nl" {
			site = "solvalutions"
			region = "nl"
		}

		c.Locals("Site", site)
		c.Locals("Region", region)
		return c.Next()
	})

	// 3. Initialize Domain Routers
	robRouter := &DomainRouter{
		App:    app,
		Domain: "robmeijerink",
	}

	solRouter := &DomainRouter{
		App:    app,
		Domain: "solvalutions",
	}

	// 4. Register strictly the routes we want to test
	robRouter.Get("/expertise", func(c fiber.Ctx) error {
		return c.SendString("Rob's Expertise Page")
	})

	solRouter.Get("/services", func(c fiber.Ctx) error {
		return c.SendString("Solvalutions Services Page")
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
			name:           "Solvalutions can access its own /services route",
			host:           "solvalutions.nl",
			path:           "/services",
			expectedStatus: fiber.StatusOK,
			expectedBody:   "Solvalutions Services Page",
		},

		// Unauthorized Cross-Tenant Access
		{
			name:           "Solvalutions CANNOT access Rob's /expertise route",
			host:           "solvalutions.nl",
			path:           "/expertise",
			expectedStatus: fiber.StatusNotFound, // Fiber returns 404 because c.Next() falls through
			expectedBody:   "Not Found",
		},
		{
			name:           "Rob CANNOT access Solvalutions' /services route",
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

func TestRender_MultiRegion_Success(t *testing.T) {
	mockEngine := &MockViews{}
	app := fiber.New(fiber.Config{Views: mockEngine})

	app.Get("/test-nl", func(c fiber.Ctx) error {
		// Mock locals for a Dutch Solvalutions visitor
		c.Locals("Site", "solvalutions")
		c.Locals("Region", "nl")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "solvalutions.nl")
		c.Locals("Path", "/test-nl")

		return Render(c, "pages/home", nil)
	})

	req := httptest.NewRequest("GET", "/test-nl", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Assert it successfully targeted the "nl" folder
	assert.Equal(t, "solvalutions/views/nl/pages/home", mockEngine.Template)
	// Assert it successfully targeted the isolated layout folder based on SharedMasterTemplate: false
	assert.Equal(t, "solvalutions/views/nl/layouts/master", mockEngine.Layout)
	assert.Equal(t, "nl", mockEngine.Data["Region"])
}

func TestRender_MultiRegion_Fallback(t *testing.T) {
	mockEngine := &MockViews{
		// Force the engine to fail on the NL template to trigger the fallback
		FailOnTemplate: "solvalutions/views/nl/pages/portfolio",
	}
	app := fiber.New(fiber.Config{Views: mockEngine})

	app.Get("/test-fallback", func(c fiber.Ctx) error {
		c.Locals("Site", "solvalutions")
		c.Locals("Region", "nl") // User wants NL
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "solvalutions.nl")
		c.Locals("Path", "/test-fallback")

		return Render(c, "pages/portfolio", nil)
	})

	req := httptest.NewRequest("GET", "/test-fallback", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Assert the Render function caught the error and fell back to the DefaultRegion ("en")
	assert.Equal(t, "solvalutions/views/en/pages/portfolio", mockEngine.Template)
}

func TestRender_FlatStructure_Success(t *testing.T) {
	mockEngine := &MockViews{}
	app := fiber.New(fiber.Config{Views: mockEngine})

	app.Get("/test-flat", func(c fiber.Ctx) error {
		// Mock locals for a flat site visitor
		c.Locals("Site", "robmeijerink")
		c.Locals("Region", "nl")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "robmeijerink.nl")
		c.Locals("Path", "/test-flat")

		customData := fiber.Map{"PageTitle": "About Rob"}
		return Render(c, "pages/about", customData)
	})

	req := httptest.NewRequest("GET", "/test-flat", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Assert it used the flat structure (NO region folder in the path)
	assert.Equal(t, "robmeijerink/views/pages/about", mockEngine.Template)
	// Assert it loaded the root layout file based on SharedMasterTemplate: true
	assert.Equal(t, "robmeijerink/views/layouts/master", mockEngine.Layout)
	assert.Equal(t, "About Rob", mockEngine.Data["PageTitle"])
}
