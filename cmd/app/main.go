package main

import (
	"log"

	"github.com/Demonyker/personal-assistant-telegram-gateway/config"
	"github.com/Demonyker/personal-assistant-telegram-gateway/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
