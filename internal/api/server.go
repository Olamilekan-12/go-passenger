// Package api implements the HTTP transport layer for the passenger service.
package api

import "github.com/Olamilekan-12/go-passenger/internal/config"

// Server holds the dependencies shared by every HTTP handler.
type Server struct {
	cfg config.Config
}

// NewServer returns a Server wired with the given configuration.
func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
