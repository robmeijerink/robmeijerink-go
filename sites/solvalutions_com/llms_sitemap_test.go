package solvalutions_com

import (
	"testing"

	"github.com/robmeijerink/robmeijerink-go/internal/web/webtest"
)

func TestLlmsCoversSitemap(t *testing.T) {
	webtest.AssertLlmsCoversSitemap(t, "solvalutions_com", "solvalutions.com", Setup)
}
