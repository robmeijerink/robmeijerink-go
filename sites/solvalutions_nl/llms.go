package solvalutions_nl

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const llmsPath = "sites/solvalutions_nl/public/llms.txt"

type faqSection struct {
	Title string
	Path  string
	Items []faqItem
}

// faqSections lists every FAQ on the site, in the order of the navigation.
// llms-full.txt is generated from it, so it always matches the pages.
var faqSections = []faqSection{
	{"Algemeen", "/", homeFAQ},
	{"Diensten", "/diensten", servicesFAQ},
	{websitesPage.Crumb, websitesPage.Path, websitesPage.FAQ},
	{conversionPage.Crumb, conversionPage.Path, conversionPage.FAQ},
	{appsPage.Crumb, appsPage.Path, appsPage.FAQ},
	{dashboardsPage.Crumb, dashboardsPage.Path, dashboardsPage.FAQ},
	{integrationsPage.Crumb, integrationsPage.Path, integrationsPage.FAQ},
	{automationPage.Crumb, automationPage.Path, automationPage.FAQ},
	{"Aanpak", "/aanpak", approachFAQ},
	{areaPage.Crumb, areaPage.Path, areaPage.FAQ},
	{"Contact", "/contact", contactFAQ},
}

func llmsFull(c fiber.Ctx) error {
	base, err := os.ReadFile(llmsPath)
	if err != nil {
		return err
	}

	c.Type("txt", "utf-8")

	return c.SendString(llmsFullText(string(base)))
}

func llmsFullText(base string) string {
	var b strings.Builder

	b.WriteString(strings.TrimRight(base, "\n"))
	b.WriteString("\n\n## Veelgestelde vragen\n\nAlle vragen en antwoorden zoals ze op de pagina's van solvalutions.nl staan.\n")

	for _, section := range faqSections {
		b.WriteString("\n### " + section.Title + "\n\nBron: " + siteURL + section.Path + "\n")

		for _, item := range section.Items {
			b.WriteString("\n#### " + item.Question + "\n\n" + item.Answer + "\n")
		}
	}

	return b.String()
}
