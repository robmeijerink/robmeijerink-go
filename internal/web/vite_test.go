package web

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to reset the global map before/after each test.
// This prevents test pollution.
func resetManifests() {
	manifests = make(map[string]map[string]ViteChunk)
}

// ---------------------------------------------------------
// LoadManifest Tests
// ---------------------------------------------------------
func TestLoadManifest(t *testing.T) {
	resetManifests()
	defer resetManifests() // Ensure cleanup after the test finishes

	t.Run("Successfully loads and parses a valid manifest", func(t *testing.T) {
		// 1. Create a temporary directory and file for the test
		tempDir := t.TempDir()
		manifestPath := filepath.Join(tempDir, "manifest.json")

		// 2. Write valid Vite JSON data to the temp file
		validJSON := `{
			"resources/js/app.js": {
				"file": "app-123456.js",
				"css": ["app-abcdef.css"]
			}
		}`
		err := os.WriteFile(manifestPath, []byte(validJSON), 0644)
		assert.NoError(t, err)

		// 3. Test the LoadManifest function
		err = LoadManifest("solvalutions", manifestPath)

		// 4. Assertions
		assert.NoError(t, err)
		assert.NotNil(t, manifests["solvalutions"])

		// Verify the data was parsed correctly into the struct
		chunk := manifests["solvalutions"]["resources/js/app.js"]
		assert.Equal(t, "app-123456.js", chunk.File)
		assert.Equal(t, []string{"app-abcdef.css"}, chunk.Css)
	})

	t.Run("Returns error when manifest file does not exist (Dev Mode)", func(t *testing.T) {
		err := LoadManifest("robmeijerink", "/path/that/does/not/exist/manifest.json")

		// Expect an error because the file doesn't exist
		assert.Error(t, err)
		// Ensure the site was NOT added to the global map
		_, exists := manifests["robmeijerink"]
		assert.False(t, exists)
	})

	t.Run("Returns error when JSON is invalid", func(t *testing.T) {
		tempDir := t.TempDir()
		manifestPath := filepath.Join(tempDir, "invalid.json")

		// Write broken JSON
		err := os.WriteFile(manifestPath, []byte(`{ "broken": "json", }`), 0644)
		assert.NoError(t, err)

		err = LoadManifest("solvalutions", manifestPath)
		assert.Error(t, err) // JSON unmarshal should fail
	})
}

// ---------------------------------------------------------
// ViteHelper Asset Tests (covers assetLookup)
// ---------------------------------------------------------
func TestViteHelper_Asset(t *testing.T) {
	resetManifests()
	defer resetManifests()

	// Seed the global map with test data for "solvalutions"
	manifests["solvalutions"] = map[string]ViteChunk{
		"resources/img/logo.png": {File: "logo-hash99.png"},
		"resources/js/app.js":    {File: "app-hash88.js"},
	}

	tests := []struct {
		name      string
		site      string
		assetName string
		expected  string
	}{
		{
			name:      "Dev Mode: Site not in manifest returns raw local path",
			site:      "robmeijerink", // No manifest loaded for Rob
			assetName: "resources/img/logo.png",
			expected:  "/assets/resources/img/logo.png",
		},
		{
			name:      "Prod Mode: Valid asset returns hashed dist path",
			site:      "solvalutions",
			assetName: "resources/img/logo.png",
			expected:  "/assets/dist/logo-hash99.png",
		},
		{
			name:      "Prod Mode: Missing asset returns fallback error string",
			site:      "solvalutions",
			assetName: "resources/img/missing.jpg",
			expected:  "ASSET-NOT-FOUND-[resources/img/missing.jpg]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := ViteHelper{Site: tt.site}
			result := helper.Asset(tt.assetName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// ---------------------------------------------------------
// ViteHelper Css Tests (covers cssLookup)
// ---------------------------------------------------------
func TestViteHelper_Css(t *testing.T) {
	resetManifests()
	defer resetManifests()

	// Seed the global map with test data
	manifests["solvalutions"] = map[string]ViteChunk{
		"resources/js/app.js": {
			File: "app-hash88.js",
			Css:  []string{"app-hash88.css"}, // Has CSS
		},
		"resources/js/no-css.js": {
			File: "no-css-hash11.js",
			Css:  nil, // Does NOT have CSS
		},
	}

	tests := []struct {
		name      string
		site      string
		entryName string
		expected  string
	}{
		{
			name:      "Dev Mode: Site not in manifest returns empty string (Vite JS handles it)",
			site:      "robmeijerink",
			entryName: "resources/js/app.js",
			expected:  "",
		},
		{
			name:      "Prod Mode: Entry has CSS array, returns hashed CSS path",
			site:      "solvalutions",
			entryName: "resources/js/app.js",
			expected:  "/assets/dist/app-hash88.css",
		},
		{
			name:      "Prod Mode: Entry exists but has no CSS array, returns empty string",
			site:      "solvalutions",
			entryName: "resources/js/no-css.js",
			expected:  "",
		},
		{
			name:      "Prod Mode: Entry completely missing from manifest, returns empty string",
			site:      "solvalutions",
			entryName: "resources/js/missing.js",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := ViteHelper{Site: tt.site}
			result := helper.Css(tt.entryName)
			assert.Equal(t, tt.expected, result)
		})
	}
}
