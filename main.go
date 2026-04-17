package main

import (
	"log"

	// Internal "Engine"
	"github.com/robmeijerink/robmeijerink-go/internal/config"
	"github.com/robmeijerink/robmeijerink-go/internal/web"

	// Domain Specific Sites
	"github.com/robmeijerink/robmeijerink-go/sites/robmeijerink"
	"github.com/robmeijerink/robmeijerink-go/sites/solvalutions"
)

func main() {
	cfg := config.Load()

	server := web.NewServer(cfg)

	robmeijerink.Setup(&web.DomainRouter{
		App:    server.App,
		Domain: "robmeijerink",
	})

	solvalutions.Setup(&web.DomainRouter{
		App:    server.App,
		Domain: "solvalutions",
	})

	log.Printf("Starting application on port %s (Production: %v)", cfg.Port, cfg.IsProd)
	log.Fatal(server.Start())
}
