package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/raboliotlegris/Tulkas/core"
	"github.com/raboliotlegris/Tulkas/routes"
)

func main() {
	log.SetLevel(log.DebugLevel)

	cfg, err := core.NewConfig()
	if err != nil {
		log.Fatal("Parsing env config: ", err)
	}
	log.SetLevel(cfg.Log_level)

	// TCP WORKER HERE

	store := core.NewSessionStore("tulkas")

	// Start router
	log.Info("Creating routes")
	if err = routes.Launch(cfg, routes.Create_router(cfg, store)); err != nil {
		log.Fatal("Tulkas crash with error: ", err)
	}
}
