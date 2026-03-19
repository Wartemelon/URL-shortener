package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/handler"
	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/repository"
	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/service"
)

func main() {
	repo := repository.NewMemoryRepository()
	baseURL := "http://localhost:8080/"
	shortener := service.NewURLShortener(repo, baseURL)

	h := handler.NewHandler(shortener)

	r := chi.NewRouter()
	r.Post("/", h.HandlePost)
	r.Get("/{short}", h.HandleGet)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start: ", err)
		os.Exit(1)
	}
}
