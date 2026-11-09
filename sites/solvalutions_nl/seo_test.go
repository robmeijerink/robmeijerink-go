package solvalutions_nl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var allLandings = []landing{websitesPage, conversionPage, appsPage, integrationsPage, dashboardsPage, automationPage, solutionsPage, logisticsPage, areaPage}

func TestLandingMetaAreUniqueAndBounded(t *testing.T) {
	titles := map[string]bool{}
	descriptions := map[string]bool{}

	for _, page := range allLandings {
		assert.LessOrEqual(t, utf8.RuneCountInString(page.Title), 66, page.Path)
		assert.GreaterOrEqual(t, utf8.RuneCountInString(page.Description), 100, page.Path)
		assert.LessOrEqual(t, utf8.RuneCountInString(page.Description), 165, page.Path)
		assert.False(t, titles[page.Title], "dubbele title: "+page.Path)
		assert.False(t, descriptions[page.Description], "dubbele description: "+page.Path)
		titles[page.Title] = true
		descriptions[page.Description] = true
	}
}

func TestExtraFaqListsFollowCopyRules(t *testing.T) {
	for name, list := range map[string][]faqItem{"home": homeFAQ, "diensten": servicesFAQ, "werkgebied": areaFAQ, "contact": contactFAQ, "aanpak": approachFAQ} {
		require.NotEmpty(t, list, name)

		for _, item := range list {
			text := item.Question + " " + item.Answer

			assert.NotContains(t, text, "—", name)
			assert.NotContains(t, text, "–", name)

			for _, banned := range []string{"Rust", "PHP", "Laravel", "vastloop", "wrijving"} {
				assert.NotContains(t, text, banned, name)
			}
		}
	}
}

func TestSplitFaqListsHaveEqualColumns(t *testing.T) {
	lists := map[string][]faqItem{"home": homeFAQ, "diensten": servicesFAQ, "werkgebied": areaFAQ, "contact": contactFAQ, "aanpak": approachFAQ}
	for _, page := range allLandings {
		lists[page.Path] = page.FAQ
	}

	for name, list := range lists {
		if faqRows(list) == 0 {
			continue
		}

		assert.Zero(t, len(list)%2, "oneven aantal vragen in een FAQ met twee kolommen: "+name)
		assert.Equal(t, len(list)/2, faqRows(list), name)
	}
}

func TestFaqRowsKeepsShortListsInOneColumn(t *testing.T) {
	assert.Zero(t, faqRows(make([]faqItem, faqSplitMin)))
	assert.Equal(t, 4, faqRows(make([]faqItem, 8)))
}

func TestLandingCopyRules(t *testing.T) {
	for _, page := range allLandings {
		text := page.Title + " " + page.Description
		for _, item := range page.FAQ {
			text += " " + item.Question + " " + item.Answer
		}

		assert.NotContains(t, text, "—", page.Path)
		assert.NotContains(t, text, "–", page.Path)

		for _, banned := range []string{"Rust", "PHP", "Laravel", "vastloop", "wrijving"} {
			assert.NotContains(t, text, banned, page.Path)
		}
	}
}

func TestFaqStructuredDataMatchesList(t *testing.T) {
	for _, page := range allLandings {
		if len(page.FAQ) == 0 {
			assert.Empty(t, string(faqStructuredData(page.FAQ)))
			continue
		}

		markup := string(faqStructuredData(page.FAQ))
		require.True(t, strings.HasPrefix(markup, `<script type="application/ld+json">`))

		payload := strings.TrimSuffix(strings.TrimPrefix(markup, `<script type="application/ld+json">`), `</script>`)

		var decoded struct {
			Type       string `json:"@type"`
			MainEntity []struct {
				Name           string `json:"name"`
				AcceptedAnswer struct {
					Text string `json:"text"`
				} `json:"acceptedAnswer"`
			} `json:"mainEntity"`
		}

		require.NoError(t, json.Unmarshal([]byte(payload), &decoded))
		assert.Equal(t, "FAQPage", decoded.Type)
		require.Len(t, decoded.MainEntity, len(page.FAQ))

		for i, item := range page.FAQ {
			assert.Equal(t, item.Question, decoded.MainEntity[i].Name)
			assert.Equal(t, item.Answer, decoded.MainEntity[i].AcceptedAnswer.Text)
		}
	}
}

func TestLandingStructuredDataIsValidJSON(t *testing.T) {
	for _, page := range allLandings {
		markup := string(landingStructuredData(page))
		payload := strings.TrimSuffix(strings.TrimPrefix(markup, `<script type="application/ld+json">`), `</script>`)

		var decoded map[string]any
		require.NoError(t, json.Unmarshal([]byte(payload), &decoded), page.Path)
	}
}

func TestSitemapHasRealLastModified(t *testing.T) {
	require.NotEmpty(t, sitemapPages)

	for _, page := range sitemapPages {
		date, ok := lastModified[page.loc]
		require.True(t, ok, "geen lastmod voor "+page.loc)

		parsed, err := time.Parse("2006-01-02", date)
		require.NoError(t, err, page.loc)
		assert.False(t, parsed.After(time.Now().AddDate(0, 0, 1)), page.loc)
	}
}

func TestBreadcrumbsAreValidJSON(t *testing.T) {
	markup := string(breadcrumbs("/over", "Over"))
	require.True(t, strings.HasPrefix(markup, `<script type="application/ld+json">`))

	payload := strings.TrimSuffix(strings.TrimPrefix(markup, `<script type="application/ld+json">`), `</script>`)

	var decoded struct {
		Context string `json:"@context"`
		Type    string `json:"@type"`
		Items   []struct {
			Position int    `json:"position"`
			Name     string `json:"name"`
			Item     string `json:"item"`
		} `json:"itemListElement"`
	}

	require.NoError(t, json.Unmarshal([]byte(payload), &decoded))
	assert.Equal(t, "BreadcrumbList", decoded.Type)
	require.Len(t, decoded.Items, 2)
	assert.Equal(t, "https://solvalutions.nl/over", decoded.Items[1].Item)
}

func TestOGImagesExist(t *testing.T) {
	for _, page := range allLandings {
		if page.OGImage == "" {
			continue
		}

		_, err := os.Stat("public" + page.OGImage)
		assert.NoError(t, err, page.Path)
		assert.NotEmpty(t, page.OGAlt, page.Path)
	}
}

func TestImagesHaveAltAndTitle(t *testing.T) {
	files, err := filepath.Glob("views/*.html")
	require.NoError(t, err)

	more, err := filepath.Glob("views/layouts/*.html")
	require.NoError(t, err)

	imgTag := regexp.MustCompile(`<img\b[^>]*>`)
	svgTag := regexp.MustCompile(`<svg\b[^>]*role="img"[^>]*>(<title>[^<]+</title>)?`)

	for _, file := range append(files, more...) {
		raw, err := os.ReadFile(file)
		require.NoError(t, err)

		for _, tag := range imgTag.FindAllString(string(raw), -1) {
			assert.Regexp(t, `\balt="[^"]+"`, tag, file)
			assert.Regexp(t, `\btitle="[^"]+"`, tag, file)
		}

		for _, tag := range svgTag.FindAllString(string(raw), -1) {
			assert.Contains(t, tag, "<title>", file)
		}
	}
}
