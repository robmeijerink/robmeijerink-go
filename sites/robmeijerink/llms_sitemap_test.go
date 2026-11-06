package robmeijerink

import (
	"testing"

	"github.com/robmeijerink/robmeijerink-go/internal/web/webtest"
)

func TestLlmsCoversSitemap(t *testing.T) {
	webtest.AssertLlmsCoversSitemap(t, "robmeijerink", "robmeijerink.nl", Setup)
}
