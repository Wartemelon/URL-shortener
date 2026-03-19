package handler

import (
	"github.com/lekan-pvp/go-f1-factual-1-tpl/internal/service"
)

type Handler struct {
	shortener service.URLShortener
}

func NewHandler(urlShortener service.URLShortener) *Handler {
	return &Handler{
		shortener: urlShortener,
	}
}
