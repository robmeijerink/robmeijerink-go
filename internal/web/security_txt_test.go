package web

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// security.txt (RFC 9116) is ignored by scanners once Expires has passed, so
// this fails a month early as a reminder to move the date forward.
func TestSecurityTxtIsValid(t *testing.T) {
	sites := map[string]string{
		"robmeijerink":     "robmeijerink.nl",
		"solvalutions_com": "solvalutions.com",
		"solvalutions_nl":  "solvalutions.nl",
	}

	for site, host := range sites {
		t.Run(site, func(t *testing.T) {
			raw, err := os.ReadFile("../../sites/" + site + "/public/.well-known/security.txt")
			require.NoError(t, err)

			fields := map[string][]string{}
			for _, line := range strings.Split(string(raw), "\n") {
				if key, value, ok := strings.Cut(line, ": "); ok && !strings.HasPrefix(line, "#") {
					fields[key] = append(fields[key], value)
				}
			}

			require.Len(t, fields["Contact"], 1)
			require.Len(t, fields["Expires"], 1)
			require.Len(t, fields["Canonical"], 1)

			assert.True(t, strings.HasPrefix(fields["Contact"][0], "mailto:"), "Contact moet een mailto zijn")
			assert.Equal(t, "https://"+host+"/.well-known/security.txt", fields["Canonical"][0])
			assert.Contains(t, fields["Encryption"], "https://robmeijerink.nl/.well-known/pgp-key.txt")

			expires, err := time.Parse(time.RFC3339, fields["Expires"][0])
			require.NoError(t, err, "Expires is geen RFC 3339-datum")
			assert.True(t, expires.After(time.Now().AddDate(0, 1, 0)), "security.txt verloopt binnen een maand: zet Expires op de drie sites een jaar vooruit")
			assert.True(t, expires.Before(time.Now().AddDate(1, 1, 0)), "Expires ligt verder dan een jaar vooruit")
		})
	}

	key, err := os.ReadFile("../../sites/robmeijerink/public/.well-known/pgp-key.txt")
	require.NoError(t, err)
	assert.Contains(t, string(key), "BEGIN PGP PUBLIC KEY BLOCK")
	assert.NotContains(t, string(key), "PRIVATE")
}
