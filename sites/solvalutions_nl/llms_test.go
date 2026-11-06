package solvalutions_nl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLlmsFullContainsEveryFaq(t *testing.T) {
	text := llmsFullText("# Solvalutions\n")

	assert.True(t, strings.HasPrefix(text, "# Solvalutions\n"))

	for _, page := range allLandings {
		assert.Contains(t, text, siteURL+page.Path, "landing ontbreekt in faqSections: "+page.Path)
	}

	for _, section := range faqSections {
		assert.NotEmpty(t, section.Items, section.Path)

		for _, item := range section.Items {
			assert.Contains(t, text, "#### "+item.Question+"\n\n"+item.Answer+"\n")
		}
	}
}
