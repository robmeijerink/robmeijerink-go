package robmeijerink

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const llmsPath = "sites/robmeijerink/public/llms.txt"

// llmsFull serves llms.txt with every FAQ appended. It is generated from
// faqItems, so it always matches the homepage.
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
	b.WriteString("\n\n## Veelgestelde vragen\n\nAlle vragen en antwoorden zoals ze op https://robmeijerink.nl/ staan.\n")

	for _, item := range faqItems {
		b.WriteString("\n### " + item.Question + "\n\n" + item.Answer + "\n")
	}

	return b.String()
}
