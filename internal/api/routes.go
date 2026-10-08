package api

import "net/http"

// Routes returns an http.Handler with every route of the service registered
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.handleHealth)
	return mux
}
