package config

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name     string
		setEnv   bool
		envValue string
		fallback string
		want     string
	}{
		{
			name:     "returns fallback when variable is unset",
			setEnv:   false,
			fallback: "default-port",
			want:     "default-port",
		},
		{
			name:     "returns value when variable is set",
			setEnv:   true,
			envValue: "9090",
			fallback: "default-port",
			want:     "9090",
		},
		{
			name:     "returns fallback when variable is set but empty",
			setEnv:   true,
			envValue: "",
			fallback: "default-port",
			want:     "default-port",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const key = "GOPICK_TEST_VAR"

			if tt.setEnv {
				t.Setenv(key, tt.envValue)
			}

			got := getEnv(key, tt.fallback)

			if got != tt.want {
				t.Errorf("getEnv(%q, %q) = %q, want %q", key, tt.fallback, got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    Config
		wantErr bool
	}{
		{
			name: "uses defaults when variables are empty",
			env:  map[string]string{"PORT": "", "APP_ENV": ""},
			want: Config{Port: "8080", Env: "development"},
		},
		{
			name: "reads values from the environment",
			env:  map[string]string{"PORT": "9090", "APP_ENV": "production"},
			want: Config{Port: "9090", Env: "production"},
		},
		{
			name:    "rejects an unknown APP_ENV",
			env:     map[string]string{"APP_ENV": "banana"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			got, err := Load()

			if tt.wantErr {
				if err == nil {
					t.Fatal("Load() returned nil error, want an error")
				}
				return
			}

			if err != nil {
				t.Fatalf("Load() returned unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("Load() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
