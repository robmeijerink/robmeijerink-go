package robmeijerink

import (
	"encoding/json"
	"html/template"
)

type faqItem struct {
	Question string
	Answer   string
}

var faqItems = []faqItem{
	{
		Question: "Wat voor opdrachten pak je op?",
		Answer:   "Van development teams die vastlopen, op techniek of op de hoeveelheid werk, tot extra capaciteit in een bestaand team. Ook pak ik complexe business cases zelfstandig op en werk ik als onderaannemer voor bureaus.",
	},
	{
		Question: "Adviseer je alleen, of bouw je ook zelf?",
		Answer:   "Beide: ik adviseer over architectuur en werk hands-on mee in je team, van full stack tot DevOps.",
	},
	{
		Question: "Kun je meewerken aan bestaande projecten?",
		Answer:   "Ja, ook in bestaande projecten. PHP/Laravel is een vast onderdeel van mijn werk, naast Rust en Go.",
	},
	{
		Question: "Hoe borg je de kwaliteit van de code?",
		Answer:   "Met automatische unit- en integratietests die via CI/CD-pipelines bij elke wijziging draaien voordat iets live gaat. Zo komen er minder fouten in productie.",
	},
	{
		Question: "Hoe werk je met AI, en wie controleert de code?",
		Answer:   "AI zet ik beheerst in om sneller op te leveren, zonder risico's te nemen. Geen vibe coding, maar stap voor stap met agentic development. De code controleer ik zelf, nauwkeurig met eigen code reviews en bijbehorende tests.",
	},
	{
		Question: "Hoe functioneer je in een team?",
		Answer:   "Zelfstandig en met afstemming waar nodig. Samen met het team bepaal ik de kern van het probleem en wat als eerste moet. Onderweg houd ik developers gemotiveerd en help ik ze groeien.",
	},
	{
		Question: "Hoe ga je om met verschillen van inzicht?",
		Answer:   "Eerst luisteren, daarna mijn afweging uitleggen. Beargumenteerde andere inzichten zijn welkom. Samen werken we de use case uit en zoeken we de beste optie.",
	},
	{
		Question: "Hoe kan mijn team zonder jou verder?",
		Answer:   "Ik lever duidelijke, geteste code met complete documentatie. Daarnaast leg ik in een gesprek uit hoe het in elkaar zit en ben ik daarna beschikbaar voor vragen.",
	},
	{
		Question: "Wanneer kun je starten, en hoe werken we samen?",
		Answer:   "Dat verschilt per opdracht en er is geen vast minimum. Ik werk remote, hybride of op locatie. Samen leggen we de afspraken vooraf schriftelijk vast.",
	},
	{
		Question: "Voldoe je aan de Wet DBA?",
		Answer:   "Ja. Ik werk als zelfstandige en autonoom: ik bepaal zelf hoe en wanneer ik mijn werk uitvoer, zonder gezag van de opdrachtgever. De afspraken leggen we vooraf schriftelijk vast.",
	},
}

func faqStructuredData(items []faqItem) template.HTML {
	type answer struct {
		Type string `json:"@type"`
		Text string `json:"text"`
	}

	type question struct {
		Type           string `json:"@type"`
		Name           string `json:"name"`
		AcceptedAnswer answer `json:"acceptedAnswer"`
	}

	entities := make([]question, 0, len(items))

	for _, item := range items {
		entities = append(entities, question{
			Type:           "Question",
			Name:           item.Question,
			AcceptedAnswer: answer{Type: "Answer", Text: item.Answer},
		})
	}

	payload, err := json.Marshal(map[string]any{
		"@context":   "https://schema.org",
		"@type":      "FAQPage",
		"mainEntity": entities,
	})
	if err != nil {
		return ""
	}

	return template.HTML(`<script type="application/ld+json">` + string(payload) + `</script>`)
}
