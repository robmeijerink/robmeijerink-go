package web

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestDomainRouter_StrictIsolation(t *testing.T) {
	// 1. Setup Fiber App
	app := fiber.New()

	// 2. Mock the Host/Tenant Middleware exactly as it works in production
	app.Use(func(c fiber.Ctx) error {
		host := c.Hostname()
		site := "robmeijerink" // Default

		if host == "solvalutions.nl" || host == "www.solvalutions.nl" {
			site = "solvalutions"
		}

		c.Locals("Site", site)
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
// This intercepts Fiber's template rendering so we can assert
// the exact data and paths the Render function generated.
type MockViews struct {
	Template string
	Layout   string
	Data     fiber.Map
}

func (m *MockViews) Load() error { return nil }

func (m *MockViews) Render(w io.Writer, template string, data any, layouts ...string) error {
	m.Template = template
	if len(layouts) > 0 {
		m.Layout = layouts[0]
	}

	// Type assert the data back to a fiber.Map so we can inspect it
	if mapData, ok := data.(fiber.Map); ok {
		m.Data = mapData
	}

	// Write a fake response so Fiber thinks it succeeded
	_, _ = w.Write([]byte("mock HTML content"))
	return nil
}

// ---------------------------------------------------------
// Render Function Test
// ---------------------------------------------------------
func TestRender_Success(t *testing.T) {
	// 1. Initialize our mock engine
	mockEngine := &MockViews{}

	// 2. Setup Fiber App with the mock engine registered
	app := fiber.New(fiber.Config{
		Views: mockEngine,
	})

	// 3. Create a test route that utilizes your Render function
	app.Get("/test-render", func(c fiber.Ctx) error {
		// Inject the required Locals.
		// (In production, your middleware handles this. If these are missing,
		// the type assertions in your Render func will cause a panic).
		c.Locals("Site", "solvalutions")
		c.Locals("IsProd", true)
		c.Locals("CanonicalHost", "solvalutions.nl")
		c.Locals("Path", "/test-render")

		// Define some custom data to see if it merges correctly
		customData := fiber.Map{
			"PageTitle": "Contact Us",
		}

		// Call the function we are testing
		return Render(c, "pages/contact", customData)
	})

	// 4. Perform the HTTP Request
	req := httptest.NewRequest("GET", "/test-render", nil)
	resp, err := app.Test(req)

	// 5. Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Assert that the paths were constructed perfectly for the specific tenant
	assert.Equal(t, "solvalutions/views/pages/contact", mockEngine.Template)
	assert.Equal(t, "solvalutions/views/layouts/master", mockEngine.Layout)

	// Assert that the global data AND custom data were merged correctly
	assert.NotNil(t, mockEngine.Data)
	assert.Equal(t, true, mockEngine.Data["IsProd"])
	assert.Equal(t, "solvalutions.nl", mockEngine.Data["CanonicalHost"])
	assert.Equal(t, "/test-render", mockEngine.Data["Path"])
	assert.Equal(t, time.Now().Year(), mockEngine.Data["CurrentYear"])

	// Verify the custom data was appended successfully
	assert.Equal(t, "Contact Us", mockEngine.Data["PageTitle"])

	// Verify that the ViteHelper was injected (checking if it exists)
	assert.Contains(t, mockEngine.Data, "Vite")
	assert.Contains(t, mockEngine.Data, "AvailableForWorkQ")
}
