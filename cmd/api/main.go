package main

import (
	"log"
	"net/http"

	"github.com/Olamilekan-12/go-passenger/internal/api"
	"github.com/Olamilekan-12/go-passenger/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	server := api.NewServer(cfg)
	addr := ":" + cfg.Port
	log.Printf("go-passenger listening on %s in %s mode", addr, cfg.Env)

	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
