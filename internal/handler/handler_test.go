package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/repository"
	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/service"
)

func TestHandler_HandlePost(t *testing.T) {
	repo := repository.NewMemoryRepository()
	s := service.NewURLShortener(repo, "http://localhost:8080/")
	h := NewHandler(s)

	t.Run("Valid POST", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://google.com"))
		w := httptest.NewRecorder()

		h.HandlePost(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status 201, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "text/plain" {
			t.Errorf("expected content-type text/plain, got %s", w.Header().Get("Content-Type"))
		}

		if w.Body.String() == "" {
			t.Error("expected non-empty response body")
		}
	})

	t.Run("Empty Body POST", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		w := httptest.NewRecorder()

		h.HandlePost(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})
}

func TestHandler_HandleGet(t *testing.T) {
	repo := repository.NewMemoryRepository()
	baseURL := "http://localhost:8080/"
	s := service.NewURLShortener(repo, baseURL)
	h := NewHandler(s)

	longURL := "https://yandex.ru"
	shortFull, _ := s.Shorten(longURL)
	short := shortFull[len(baseURL):]

	t.Run("Existing short URL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/"+short, nil)
		w := httptest.NewRecorder()

		h.HandleGet(w, req)

		if w.Code != http.StatusTemporaryRedirect {
			t.Errorf("expected status 307, got %d", w.Code)
		}

		if w.Header().Get("Location") != longURL {
			t.Errorf("expected location %s, got %s", longURL, w.Header().Get("Location"))
		}
	})

	t.Run("Non-existing short URL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
		w := httptest.NewRecorder()

		h.HandleGet(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})
}
