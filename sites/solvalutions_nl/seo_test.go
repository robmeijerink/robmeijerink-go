package solvalutions_nl

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var allLandings = []landing{websitesPage, conversionPage, appsPage, integrationsPage, areaPage}

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
