package robmeijerink

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFaqStructuredData(t *testing.T) {
	markup := string(faqStructuredData(faqItems))

	require.True(t, strings.HasPrefix(markup, `<script type="application/ld+json">`))
	require.True(t, strings.HasSuffix(markup, `</script>`))

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
	require.Len(t, decoded.MainEntity, len(faqItems))

	for i, item := range faqItems {
		assert.Equal(t, item.Question, decoded.MainEntity[i].Name)
		assert.Equal(t, item.Answer, decoded.MainEntity[i].AcceptedAnswer.Text)
	}
}
