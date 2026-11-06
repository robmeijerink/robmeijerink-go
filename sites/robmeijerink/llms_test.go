package robmeijerink

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLlmsFullContainsEveryFaq(t *testing.T) {
	text := llmsFullText("# Rob Meijerink\n")

	assert.True(t, strings.HasPrefix(text, "# Rob Meijerink\n"))

	for _, item := range faqItems {
		assert.Contains(t, text, "### "+item.Question+"\n\n"+item.Answer+"\n")
	}
}
