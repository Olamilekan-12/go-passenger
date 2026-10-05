package main

import (
	"log"
	"net/http"

	"github.com/Olamilekan-12/go-passenger/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	addr := ":" + cfg.Port
	log.Printf("go-passenger listening on %s in %s mode", addr, cfg.Env)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
