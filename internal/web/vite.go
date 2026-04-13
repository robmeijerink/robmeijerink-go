package web

import (
	"encoding/json"
	"fmt"
	"os"
)

// ViteHelper is injected into the templates to provide site-aware asset resolution.
// Public, so the template engine can access it.
type ViteHelper struct {
	Site string
}

// Asset calls the internal assetLookup function.
// Public method, used in HTML as {{ .Vite.Asset "..." }}
func (v ViteHelper) Asset(assetName string) string {
	return assetLookup(v.Site, assetName)
}

// Css calls the internal cssLookup function.
// Public method, used in HTML as {{ .Vite.Css "..." }}
func (v ViteHelper) Css(entryName string) string {
	return cssLookup(v.Site, entryName)
}

// ViteChunk perfectly matches the structure in Vite's manifest.json.
type ViteChunk struct {
	File string   `json:"file"`
	Css  []string `json:"css"` // Vite places the extracted CSS files here.
}

// manifests is an in-memory map mapping SiteName -> EntryName -> ViteChunk.
var manifests = make(map[string]map[string]ViteChunk)

// LoadManifest is called once in main.go during startup.
// This remains public (capital L) because main.go needs to call it.
func LoadManifest(site, manifestPath string) error {
	file, err := os.ReadFile(manifestPath)
	if err != nil {
		return err // In development there is no manifest, which is expected and fine.
	}

	var manifest map[string]ViteChunk
	if err := json.Unmarshal(file, &manifest); err != nil {
		return err
	}

	manifests[site] = manifest
	return nil
}

// assetLookup searches for the hashed JS or image file.
// Now private (lowercase 'a'): only used within this package.
func assetLookup(site, assetName string) string {
	manifest, ok := manifests[site]
	if !ok {
		return "/assets/" + site + "/" + assetName
	}

	chunk, ok := manifest[assetName]
	if !ok {
		return fmt.Sprintf("ASSET-NOT-FOUND-[%s]", assetName)
	}

	return fmt.Sprintf("/assets/%s/%s", site, chunk.File)
}

// cssLookup specifically searches for the CSS belonging to a JS entry file.
// Now private (lowercase 'c'): only used within this package.
func cssLookup(site, entryName string) string {
	manifest, ok := manifests[site]
	if !ok {
		return "" // Locally we return nothing, Vite's HMR handles the CSS injection via JS.
	}

	chunk, ok := manifest[entryName]
	if !ok || len(chunk.Css) == 0 {
		return "" // No CSS found for this specific chunk.
	}

	// Return the hashed CSS file path.
	return fmt.Sprintf("/assets/%s/%s", site, chunk.Css[0])
}
