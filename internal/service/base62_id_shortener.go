package service

import (
	"fmt"
	netURL "net/url"
	"sync"
)

type urlShortener struct {
	counter  int
	repo     Repository
	baseURL  string
	alphabet string
	mu       sync.RWMutex
}

func NewURLShortener(repo Repository, baseURL string) URLShortener {
	return &urlShortener{
		repo:     repo,
		baseURL:  baseURL,
		alphabet: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
}

func (s *urlShortener) base62Encode(n int) string {
	if n == 0 {
		return string(s.alphabet[0])
	}
	result := ""
	for n > 0 {
		result = string(s.alphabet[n%62]) + result
		n /= 62
	}
	return result
}

func (s *urlShortener) Shorten(long string) (string, error) {
	if long == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := netURL.Parse(long)
	if err != nil {
		return "", fmt.Errorf("invalid URL format: %v", err)
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", fmt.Errorf("URL must include scheme and host")
	}

	s.mu.Lock()
	s.counter++
	short := s.base62Encode(s.counter)
	s.mu.Unlock()

	if err = s.repo.Save(short, long); err != nil {
		return "", err
	}

	return s.baseURL + short, nil
}

func (s *urlShortener) Resolve(short string) (string, error) {
	return s.repo.Get(short)
}
