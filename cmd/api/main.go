package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:              addr,
		Handler:           server.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("go-passenger listening on %s in %s mode", addr, cfg.Env)
		serverErrors <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("server failed: %v", err)
	case sig := <-shutdown:
		log.Printf("shutdown started: received %s", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)

			if err := srv.Close(); err != nil {
				log.Printf("forced close failed: %v", err)
			}
		}
		log.Println("shutdown complete")
	}
}
