package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Olamilekan-12/go-passenger/internal/api"
	"github.com/Olamilekan-12/go-passenger/internal/config"
)

func TestHealthEndpoint(t *testing.T) {
	server := api.NewServer(config.Config{
		Port: "8080",
		Env:  "development",
	})
	handler := server.Routes()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer req.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status= %d, want %d", res.StatusCode, http.StatusOK)
	}

	if got := res.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q want %q", got, "application/json")
	}

	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decoding body: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("status = %q, want %q", body.Status, "ok")
	}
}

func TestRouting(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{"health with GET", http.MethodGet, "/health", http.StatusOK},
		{"health with POST", http.MethodPost, "/health", http.StatusMethodNotAllowed},
		{"unknown path", http.MethodGet, "/nope", http.StatusNotFound},
	}

	handler := api.NewServer(config.Config{
		Port: "8080",
		Env:  "development",
	}).Routes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))

			if rec.Code != tt.wantStatus {
				t.Errorf("%s %s = %d, want %d", tt.method, tt.path, rec.Code, tt.wantStatus)
			}
		})
	}
}
