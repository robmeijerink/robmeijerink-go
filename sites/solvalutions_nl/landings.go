package solvalutions_nl

const (
	introLead    = "Vertel kort wat je nodig hebt en wat je ermee wilt bereiken. Je krijgt een eerlijk advies en daarna een voorstel met een vaste prijs per fase. Liever eerst even sparren? Bel of mail gerust, dat is vrijblijvend."
	introButton  = "Plan een kennismaking"
	introSubject = "Kennismaking Solvalutions"
)

var websitesPage = landing{
	Path:        "/websites",
	Crumb:       "Website laten maken",
	Title:       "Website laten maken Zutphen: webdesign & webshop | Solvalutions",
	Description: "Een website laten maken in Zutphen? Webdesign, webshop en formulieren die snel laden, goed gevonden worden en aanvragen opleveren. Vaste prijs per fase.",
	Service:     "Website laten maken",
	ServiceType: "Webdesign en webontwikkeling",
	CTATitle:    "Een nieuwe website nodig?",
	CTALead:     introLead,
	CTAButton:   introButton,
	CTASubject:  introSubject,
	OGImage:     "/assets/img/solvalutions-og-web.png",
	OGAlt:       "Website laten maken in Zutphen, Solvalutions",
	View:        "websites",
	FAQ: []faqItem{
		{"Wat kost een website laten maken?", "Dat hangt af van wat de website moet doen: een eenvoudige bedrijfswebsite is iets anders dan een webshop of een site met koppelingen. Daarom begin ik met een korte analyse en maak ik daarna een voorstel met per fase een vaste prijs, zodat je vooraf weet wat je betaalt en zelf bepaalt hoe ver je gaat."},
		{"Hoe lang duurt het om een website te laten bouwen?", "Dat verschilt per website. In het voorstel staat per fase wat er wordt opgeleverd en wanneer, zodat je geen vage einddatum krijgt."},
		{"Kan mijn bestaande website worden verbeterd in plaats van vervangen?", "Ja. Eerst kijk ik waar de website tegenvalt: snelheid, opbouw, vindbaarheid of het aantal aanvragen. Daarna bepalen we of verbeteren genoeg is of dat een nieuwe website beter past."},
		{"Wordt mijn website straks goed gevonden in Google?", "Snelheid, een heldere opbouw, goede titels en beschrijvingen en toegankelijkheid horen bij de oplevering. Daarmee leg ik een stevige basis. Een positie in Google kan niemand garanderen, want die hangt ook af van concurrentie, inhoud en tijd."},
		{"Kan ik een webshop of formulier op mijn website krijgen?", "Ja. Denk aan een webshop, een offerteformulier, een afsprakenmodule of een contactformulier dat de aanvragen netjes doorstuurt naar jouw systemen."},
		{"Werk je alleen voor bedrijven in Zutphen?", "Ik zit in Zutphen en werk voor bedrijven in heel Oost-Nederland. Kennismaken kan bij je op locatie, het vervolg kan ook op afstand."},
	},
}

var conversionPage = landing{
	Path:        "/conversie",
	Crumb:       "Conversie optimalisatie",
	Title:       "Conversie optimalisatie website Zutphen | Solvalutions",
	Description: "Meer aanvragen, bellers of verkopen uit je bestaande website. Conversie optimalisatie in Zutphen: meten waar bezoekers afhaken en dat gericht verbeteren.",
	Service:     "Conversie optimalisatie",
	ServiceType: "Conversie-optimalisatie van websites",
	CTATitle:    "Meer aanvragen uit je website?",
	CTALead:     introLead,
	CTAButton:   introButton,
	CTASubject:  introSubject,
	OGImage:     "/assets/img/solvalutions-og-conversie.png",
	OGAlt:       "Meer aanvragen uit je website, Solvalutions",
	View:        "conversion",
	FAQ: []faqItem{
		{"Wat is conversie optimalisatie?", "Het verbeteren van je website zodat een groter deel van de bezoekers doet wat jij wilt: bellen, een formulier invullen, een afspraak plannen of iets kopen. Je haalt dus meer uit het verkeer dat je al hebt."},
		{"Hoe weet je wat er beter kan?", "Ik meet eerst hoe bezoekers je website gebruiken: waar ze binnenkomen, welke pagina's ze lezen en op welk punt ze afhaken. Daaruit volgt een lijst met verbeteringen, geordend op wat het meeste oplevert voor de minste moeite."},
		{"Hoeveel meer aanvragen kan ik verwachten?", "Dat beloof ik vooraf niet, omdat het afhangt van je aanbod, je verkeer en je huidige website. Wel leggen we de beginsituatie vast, zodat elke verbetering zichtbaar en meetbaar is."},
		{"Heb ik veel bezoekers nodig?", "Hoe meer bezoekers, hoe sneller je betrouwbare uitkomsten van tests ziet. Heb je weinig bezoekers, dan richt ik me op duidelijke verbeteringen in plaats van op tests."},
		{"Pas je ook mijn bestaande website aan?", "Ja, dat is meestal het uitgangspunt. Kleine aanpassingen kunnen veel doen, zoals een duidelijkere pagina of een eenvoudiger formulier. Is de basis te zwak, dan zeg ik dat eerlijk."},
	},
}

var appsPage = landing{
	Path:        "/apps",
	Crumb:       "App of webapplicatie",
	Title:       "App laten maken Zutphen: webapp & mobiele app | Solvalutions",
	Description: "Een app of webapplicatie laten maken in Zutphen? Portalen, klantomgevingen en bedrijfsapps die ook op je telefoon werken, gekoppeld aan je systemen.",
	Service:     "App en webapplicatie op maat",
	ServiceType: "Softwareontwikkeling",
	CTATitle:    "Een app of webapplicatie nodig?",
	CTALead:     introLead,
	CTAButton:   introButton,
	CTASubject:  introSubject,
	OGImage:     "/assets/img/solvalutions-og-apps.png",
	OGAlt:       "App of webapplicatie op maat, Solvalutions",
	View:        "apps",
	FAQ: []faqItem{
		{"Wat is het verschil tussen een app en een webapplicatie?", "Een webapplicatie werkt in de browser, op een computer en op een telefoon, zonder installatie. Een app installeer je vanuit een appstore. Voor veel bedrijfstoepassingen is een webapplicatie die goed op een telefoon werkt voldoende."},
		{"Kan een webapplicatie op mijn telefoon werken?", "Ja. De applicatie past zich aan het scherm aan en kan als snelkoppeling op het beginscherm staan, zodat medewerkers of klanten er snel bij kunnen."},
		{"Kan de applicatie koppelen met mijn bestaande systemen?", "Ja. Een nieuwe applicatie wordt vaak pas echt nuttig als gegevens vanzelf uit je boekhouding, ERP of webshop komen. Zie ook de pagina over koppelingen en API's."},
		{"Wat kost een app of webapplicatie?", "Dat hangt af van wat hij moet kunnen. Ik begin met een korte analyse en maak daarna een voorstel met per fase een vaste prijs. Je kunt dus klein beginnen en zelf bepalen hoe ver je gaat."},
		{"Kunnen klanten of medewerkers inloggen?", "Ja. Een omgeving met inloggen en verschillende rechten is een veelgebruikte vorm, bijvoorbeeld een klantportaal of een medewerkersomgeving."},
	},
}

var integrationsPage = landing{
	Path:        "/integraties",
	Crumb:       "API en systeemkoppelingen",
	Title:       "API koppeling laten maken: ERP & systemen | Solvalutions",
	Description: "API-, ERP- en systeemkoppelingen op maat, ook tussen je website en CRM, boekhouding of betaling. Gegevens lopen vanzelf door. Vanuit Zutphen, vaste prijs per fase.",
	Service:     "API-, systeem- en ERP-koppelingen",
	ServiceType: "Systeemintegratie",
	CTATitle:    "Praten je systemen nog niet met elkaar?",
	OGImage:     "/assets/img/solvalutions-og-integraties.png",
	OGAlt:       "API- en ERP-koppelingen, Solvalutions",
	View:        "integrations",
	FAQ: []faqItem{
		{"Welke systemen kun je aan elkaar koppelen?", "Ik kan elk systeem aan elkaar koppelen: de achterliggende techniek is hetzelfde. Het gaat om boekhouding, ERP, webshop, magazijn, CRM of je eigen website."},
		{"Kan mijn website gekoppeld worden aan CRM, boekhouding of betaling?", "Ja. Een aanvraag of bestelling op je website kan automatisch in je CRM, je boekhouding of je magazijnsysteem terechtkomen, zonder dat iemand iets overtypt."},
		{"Kan ik een koppeling laten maken met een systeem zonder API?", "Meestal wel, bijvoorbeeld via een export of rechtstreeks via de database. Dat bekijken we in de analyse, zodat je vooraf weet welke route het beste past."},
		{"Wat gebeurt er als een koppeling een keer uitvalt?", "Een goede koppeling meldt fouten en probeert het opnieuw, zodat er geen gegevens verloren gaan. Hoe dat precies werkt leg ik vooraf met je vast."},
		{"Kun je een nieuwe API voor mij bouwen?", "Ja. Heb je een eigen systeem dat gegevens moet aanbieden aan anderen, dan bouw ik een API op maat."},
	},
}

var areaPage = landing{
	Path:        "/werkgebied",
	Crumb:       "Werkgebied",
	Title:       "Webdesign en software in Zutphen en omgeving | Solvalutions",
	Description: "Gevestigd in Zutphen en actief voor bedrijven in Gelderland, Overijssel en Utrecht. Kennismaken bij je op locatie, het vervolg ook op afstand.",
	CTATitle:    "Zullen we kennismaken?",
	CTALead:     introLead,
	CTAButton:   introButton,
	CTASubject:  introSubject,
	View:        "area",
	FAQ:         areaFAQ,
}

var dashboardsPage = landing{
	Path:        "/dashboards",
	Crumb:       "Dashboards en data",
	Title:       "Dashboard laten maken Zutphen: data visualisatie | Solvalutions",
	Description: "Een dashboard laten maken in Zutphen? Data uit je boekhouding, ERP en webshop vanzelf samengebracht en helder in beeld, altijd actueel en op elk scherm.",
	Service:     "Dashboards en datavisualisatie",
	ServiceType: "Datavisualisatie en rapportage",
	CTATitle:    "Cijfers die niet bij elkaar komen?",
	CTALead:     introLead,
	CTAButton:   introButton,
	CTASubject:  introSubject,
	OGImage:     "/assets/img/solvalutions-og-dashboards.png",
	OGAlt:       "Dashboards en data, Solvalutions",
	View:        "dashboards",
	FAQ: []faqItem{
		{"Wat is een dashboard?", "Een overzicht dat je belangrijkste cijfers op één plek toont, in grafieken en tabellen, zodat je in één oogopslag ziet waar je op moet sturen. Denk aan omzet, orders, voorraad, aanvragen of uren."},
		{"Gebruik je een bestaande tool of bouw je het dashboard zelf?", "Beide kan. Ik kan je bronnen koppelen aan een bestaande tool voor rapportages, of een dashboard bouwen als webapplicatie op maat. Welke route past hangt af van je wensen, je budget en wie het dashboard gaat gebruiken. Dat bepalen we in de analyse."},
		{"Uit welke systemen kan de data komen?", "Uit vrijwel elke bron: je boekhouding, ERP, webshop, CRM, website of zelfs losse Excel-bestanden. Ik koppel de bronnen, zodat de gegevens vanzelf bij elkaar komen en niemand ze meer hoeft samen te voegen."},
		{"Hoe actueel zijn de cijfers?", "Dat hangt af van de bron en van wat je nodig hebt. Waar het kan worden de gegevens automatisch bijgewerkt, bijvoorbeeld dagelijks of vrijwel direct. Hoe vaak dat gebeurt leggen we vooraf vast."},
		{"Kan ik het dashboard ook op mijn telefoon bekijken?", "Ja. Het dashboard is een webapplicatie die zich aanpast aan het scherm, dus je kunt het op een computer, tablet en telefoon bekijken."},
		{"Wat kost een dashboard?", "Dat hangt af van het aantal bronnen en van wat je wilt laten zien. Ik begin met een korte analyse en maak daarna een voorstel met per fase een vaste prijs, zodat je klein kunt beginnen en zelf bepaalt hoe ver je gaat."},
	},
}

var homeFAQ = []faqItem{
	{"Wat doet Solvalutions?", "Solvalutions bouwt websites, webshops, apps en webapplicaties, dashboards en koppelingen tussen systemen, en maakt software op maat. De klanten zijn vooral MKB en logistiek, vanuit Zutphen. Je werkt direct met de engineer die het ook bouwt."},
	{"Voor welke bedrijven werkt Solvalutions?", "Vooral voor MKB en logistiek. Daarnaast ook voor industriële bedrijven, stichtingen, bureaus en teams met een digitaal product."},
	{"Waar begin ik als ik niet weet wat ik nodig heb?", "Met een knelpunt-analyse van twee tot drie werkdagen. Je krijgt een diagnose van waar het knelt, een plan met prioriteiten en een gesprek over de vervolgstappen. Liever eerst even sparren kan ook, dat is vrijblijvend."},
	{"Wat kost een website, app of koppeling?", "Dat hangt af van wat er nodig is, dus een bedrag noemen zonder gesprek zegt weinig. Na een korte analyse volgt een voorstel met een vaste prijs per fase. Je weet vooraf wat het kost en zit niet vast aan een licentie of aan één leverancier."},
	{"Werkt Solvalutions ook voor bedrijven buiten Zutphen?", "Ja, voor bedrijven in Gelderland, Overijssel en Utrecht. Kennismaken kan bij je op locatie, het vervolg kan ook op afstand."},
	{"Wordt er AI gebruikt bij het bouwen?", "Ja, maar niet als vervanging van vakwerk. AI helpt om gerichter en sneller te bouwen. Niets gaat ongezien live: alles wordt beoordeeld voordat het wordt opgeleverd."},
}

var servicesFAQ = []faqItem{
	{"Welke diensten biedt Solvalutions?", "Websites en webshops, conversie optimalisatie, apps en webapplicaties, dashboards, koppelingen tussen systemen en het moderniseren of automatiseren van bestaande software."},
	{"Kan ik meerdere diensten combineren?", "Ja, en dat is vaak slim. Een website die direct koppelt met je CRM of boekhouding, of een dashboard dat gegevens uit meerdere systemen toont, vraagt om beide kanten. Dat regelt één partij, dus je hoeft niet te schakelen tussen leveranciers."},
	{"Wat als ik nog niet weet welke dienst bij mijn probleem past?", "Gebruik de check op de homepage of bel of mail even. Samen kijken we waar het knelt en wat daarvoor het beste werkt."},
	{"Hoe werkt de prijs?", "Per fase een vaste prijs. Je weet wat een fase oplevert en kost voordat die begint."},
}

var areaFAQ = []faqItem{
	{"In welke regio werkt Solvalutions?", "Solvalutions is gevestigd in Zutphen en actief voor bedrijven in Gelderland, Overijssel en Utrecht, bijvoorbeeld in Apeldoorn, Deventer, Arnhem, Nijmegen, Enschede, Zwolle en Amersfoort."},
	{"Kom je bij mij op locatie?", "Voor een eerste kennismaking kan dat. Het vervolg loopt ook goed op afstand."},
	{"Kan ik ook werken met Solvalutions als ik buiten deze provincies zit?", "Ja, werken op afstand kan ook. Je hoeft dus niet in de buurt te zitten. Kennismaken gaat dan online of telefonisch."},
	{"Hoe snel krijg ik antwoord op een vraag?", "Bel of mail gerust. Meestal heb je binnen een werkdag antwoord."},
}
