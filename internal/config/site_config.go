package config

type SiteConfig struct {
	IsMultiRegion        bool
	DefaultRegion        string
	SharedMasterTemplate bool
}

var SiteRegistry = map[string]SiteConfig{
	"solvalutions": {
		IsMultiRegion:        true,
		DefaultRegion:        "en",
		SharedMasterTemplate: false,
	},
	"robmeijerink": {
		IsMultiRegion:        false,
		DefaultRegion:        "nl", // Optional
		SharedMasterTemplate: true,
	},
}
