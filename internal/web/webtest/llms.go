// Package webtest holds helpers shared by the site tests.
package webtest

import (
	"io"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var sitemapLoc = regexp.MustCompile(`<loc>([^<]+)</loc>`)

// AssertLlmsCoversSitemap fails when a page from the site's sitemap is not
// linked in public/llms.txt, so a new page cannot be forgotten there.
func AssertLlmsCoversSitemap(t *testing.T, site, host string, setup func(*web.DomainRouter)) {
	t.Helper()

	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		c.Locals("Site", site)
		return c.Next()
	})
	setup(&web.DomainRouter{App: app, Domain: site})

	req := httptest.NewRequest("GET", "/sitemap.xml", nil)
	req.Host = host

	resp, err := app.Test(req)
	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	locs := sitemapLoc.FindAllStringSubmatch(string(body), -1)
	require.NotEmpty(t, locs, "sitemap bevat geen pagina's")

	llms, err := os.ReadFile("public/llms.txt")
	require.NoError(t, err)

	for _, loc := range locs {
		assert.True(t, strings.Contains(string(llms), "("+loc[1]+")"), "pagina uit de sitemap ontbreekt in llms.txt: "+loc[1])
	}
}
