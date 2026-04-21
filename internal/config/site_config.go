package config

type SiteConfig struct {
	IsMultiRegion bool
	DefaultRegion string
}

var SiteRegistry = map[string]SiteConfig{
	"solvalutions": {
		IsMultiRegion: true,
		DefaultRegion: "en",
	},
	"robmeijerink": {
		IsMultiRegion: false,
		DefaultRegion: "nl", // Optional
	},
}
