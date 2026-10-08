package api

import "net/http"

// healthResponse is the body returned by the health endpoint.
type healthResponse struct {
	Status string `json:"status"`
}

// handleHealth reports that the process is alive and serving requests.
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Status: "ok",
	})
}
