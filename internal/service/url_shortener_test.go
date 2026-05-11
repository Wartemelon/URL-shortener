package service

import (
	"testing"

	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/repository"
)

func TestURLShortener_Shorten(t *testing.T) {
	repo := repository.NewMemoryRepository()
	baseURL := "http://localhost:8080/"
	s := NewURLShortener(repo, baseURL)

	tests := []struct {
		name    string
		longURL string
		wantErr bool
	}{
		{
			name:    "Valid URL",
			longURL: "https://google.com",
			wantErr: false,
		},
		{
			name:    "Invalid URL format",
			longURL: "not-a-url",
			wantErr: true,
		},
		{
			name:    "Missing scheme",
			longURL: "google.com",
			wantErr: true,
		},
		{
			name:    "Empty URL",
			longURL: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			short, err := s.Shorten(tt.longURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && short == "" {
				t.Error("Shorten() returned empty string for valid URL")
			}
		})
	}
}

func TestURLShortener_Resolve(t *testing.T) {
	repo := repository.NewMemoryRepository()
	baseURL := "http://localhost:8080/"
	s := NewURLShortener(repo, baseURL)

	longURL := "https://yandex.ru"
	shortFull, _ := s.Shorten(longURL)
	short := shortFull[len(baseURL):]

	t.Run("Resolve existing", func(t *testing.T) {
		got, err := s.Resolve(short)
		if err != nil {
			t.Errorf("Resolve() unexpected error: %v", err)
		}
		if got != longURL {
			t.Errorf("Resolve() = %v, want %v", got, longURL)
		}
	})

	t.Run("Resolve non-existing", func(t *testing.T) {
		_, err := s.Resolve("nonexistent")
		if err == nil {
			t.Error("Resolve() expected error for non-existing short code")
		}
	})
}
