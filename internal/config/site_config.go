package config

type SiteConfig struct {
	IsMultiRegion bool
	DefaultRegion string
}

var SiteRegistry = map[string]SiteConfig{
	"solvalutions": {
		IsMultiRegion: true,
		DefaultRegion: "nl", // @todo switch to en later
	},
	"robmeijerink": {
		IsMultiRegion: false,
		DefaultRegion: "nl", // Optional
	},
}
