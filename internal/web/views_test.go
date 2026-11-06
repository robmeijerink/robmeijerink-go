package web

import (
	"testing"

	"github.com/gofiber/template/html/v2"
	"github.com/stretchr/testify/require"
)

// One broken template stops the engine from loading any view, so every page
// of every site returns a 500. Parse them all here so that fails in CI instead.
func TestAllViewsParse(t *testing.T) {
	engine := html.New("../../sites", ".html")

	require.NoError(t, engine.Load())
}
