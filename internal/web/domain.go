package web

import (
	"time"

	"github.com/robmeijerink/robmeijerink-go/internal/availability"
	"github.com/robmeijerink/robmeijerink-go/internal/config"

	"github.com/gofiber/fiber/v3"
)

type DomainRouter struct {
	App    *fiber.App
	Domain string
}

func (r *DomainRouter) Get(path string, handler fiber.Handler) {
	r.App.Get(path, func(c fiber.Ctx) error {
		if c.Locals("Site") != r.Domain {
			return c.Next()
		}
		return handler(c)
	})
}

func Render(c fiber.Ctx, template string, data fiber.Map) error {
	site := c.Locals("Site").(string)
	region := c.Locals("Region").(string)
	isProd := c.Locals("IsProd").(bool)
	canonicalHost := c.Locals("CanonicalHost").(string)
	path := c.Locals("Path").(string)

	availabilityDate := availability.NextQuarter(time.Now())

	fullData := fiber.Map{
		"Vite":              ViteHelper{Site: site},
		"IsProd":            isProd,
		"CanonicalHost":     canonicalHost,
		"Path":              path,
		"Region":            region,
		"AvailableForWorkQ": availabilityDate,
		"CurrentYear":       time.Now().Year(),
	}

	if data != nil {
		for key, value := range data {
			fullData[key] = value
		}
	}

	config, exists := config.SiteRegistry[site]
	if !exists {
		return fiber.NewError(fiber.StatusInternalServerError, "Site configuration missing for: "+site)
	}

	var templatePath string
	var layoutPath string

	if config.SharedMasterTemplate {
		layoutPath = site + "/views/layouts/master"
	} else if config.IsMultiRegion {
		layoutPath = site + "/views/" + region + "/layouts/master"
	} else {
		layoutPath = site + "/views/layouts/master"
	}

	if config.IsMultiRegion {
		templatePath = site + "/views/" + region + "/" + template

		err := c.Render(templatePath, fullData, layoutPath)
		if err != nil {
			if region != config.DefaultRegion {
				fallbackTemplate := site + "/views/" + config.DefaultRegion + "/" + template
				fallbackLayout := layoutPath

				if !config.SharedMasterTemplate {
					fallbackLayout = site + "/views/" + config.DefaultRegion + "/layouts/master"
				}

				return c.Render(fallbackTemplate, fullData, fallbackLayout)
			}
			return err
		}
		return nil

	} else {
		templatePath = site + "/views/" + template
		return c.Render(templatePath, fullData, layoutPath)
	}
}
