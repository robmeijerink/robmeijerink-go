package solvalutions_nl

import (
	"encoding/json"
	"html/template"

	"github.com/gofiber/fiber/v3"
	"github.com/robmeijerink/robmeijerink-go/internal/web"
)

const siteURL = "https://solvalutions.nl"

type faqItem struct {
	Question string
	Answer   string
}

type landing struct {
	Path        string
	Crumb       string
	Title       string
	Description string
	Service     string
	ServiceType string
	CTATitle    string
	OGImage     string
	OGAlt       string
	CTALead     string
	CTAButton   string
	CTASubject  string
	View        string
	FAQ         []faqItem
}

func organization() map[string]any {
	return map[string]any{
		"@type": "Organization",
		"name":  "Solvalutions",
		"url":   siteURL,
		"logo":  siteURL + "/assets/img/solvalutions-logo.png",
	}
}

func breadcrumbList(path, crumb string) map[string]any {
	return map[string]any{
		"@type": "BreadcrumbList",
		"itemListElement": []any{
			map[string]any{"@type": "ListItem", "position": 1, "name": "Solvalutions", "item": siteURL + "/"},
			map[string]any{"@type": "ListItem", "position": 2, "name": crumb, "item": siteURL + path},
		},
	}
}

func breadcrumbs(path, crumb string) template.HTML {
	list := breadcrumbList(path, crumb)
	list["@context"] = "https://schema.org"

	return jsonLD(list)
}

func landingStructuredData(page landing) template.HTML {
	graph := []any{
		map[string]any{
			"@type":       "WebPage",
			"name":        page.Title,
			"description": page.Description,
			"url":         siteURL + page.Path,
			"isPartOf":    map[string]any{"@type": "WebSite", "name": "Solvalutions", "url": siteURL},
		},
		breadcrumbList(page.Path, page.Crumb),
	}

	if page.Service != "" {
		graph = append(graph, map[string]any{
			"@type":       "Service",
			"name":        page.Service,
			"serviceType": page.ServiceType,
			"provider":    organization(),
			"areaServed":  []any{map[string]any{"@type": "City", "name": "Zutphen"}, map[string]any{"@type": "AdministrativeArea", "name": "Gelderland"}, map[string]any{"@type": "AdministrativeArea", "name": "Overijssel"}},
		})
	}

	return jsonLD(map[string]any{"@context": "https://schema.org", "@graph": graph})
}

func faqStructuredData(items []faqItem) template.HTML {
	if len(items) == 0 {
		return ""
	}

	entities := make([]any, 0, len(items))

	for _, item := range items {
		entities = append(entities, map[string]any{
			"@type":          "Question",
			"name":           item.Question,
			"acceptedAnswer": map[string]any{"@type": "Answer", "text": item.Answer},
		})
	}

	return jsonLD(map[string]any{"@context": "https://schema.org", "@type": "FAQPage", "mainEntity": entities})
}

func jsonLD(payload map[string]any) template.HTML {
	encoded, err := json.Marshal(payload)
	if err != nil {
		return ""
	}

	return template.HTML(`<script type="application/ld+json">` + string(encoded) + `</script>`)
}

func renderLanding(c fiber.Ctx, page landing) error {
	return web.Render(c, page.View, fiber.Map{
		"Title":           page.Title,
		"MetaDescription": page.Description,
		"StructuredData":  landingStructuredData(page) + faqStructuredData(page.FAQ),
		"CTATitle":        page.CTATitle,
		"OGImage":         page.OGImage,
		"OGAlt":           page.OGAlt,
		"CTALead":         page.CTALead,
		"CTAButton":       page.CTAButton,
		"CTASubject":      page.CTASubject,
		"FAQ":             page.FAQ,
	})
}

func websites(c fiber.Ctx) error     { return renderLanding(c, websitesPage) }
func conversion(c fiber.Ctx) error   { return renderLanding(c, conversionPage) }
func apps(c fiber.Ctx) error         { return renderLanding(c, appsPage) }
func area(c fiber.Ctx) error         { return renderLanding(c, areaPage) }
func integrations(c fiber.Ctx) error { return renderLanding(c, integrationsPage) }

func dashboards(c fiber.Ctx) error { return renderLanding(c, dashboardsPage) }
