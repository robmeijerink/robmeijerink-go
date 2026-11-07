package solvalutions_com

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSitemapHasRealLastModified(t *testing.T) {
	require.NotEmpty(t, sitemapPages)

	for _, page := range sitemapPages {
		date, ok := lastModified[page.loc]
		require.True(t, ok, "no lastmod for "+page.loc)

		parsed, err := time.Parse("2006-01-02", date)
		require.NoError(t, err, page.loc)
		assert.False(t, parsed.After(time.Now().AddDate(0, 0, 1)), page.loc)
	}
}
