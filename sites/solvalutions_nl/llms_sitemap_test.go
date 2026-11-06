package solvalutions_nl

import (
	"testing"

	"github.com/robmeijerink/robmeijerink-go/internal/web/webtest"
)

func TestLlmsCoversSitemap(t *testing.T) {
	webtest.AssertLlmsCoversSitemap(t, "solvalutions_nl", "solvalutions.nl", Setup)
}
