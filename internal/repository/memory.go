package repository

import (
	"fmt"
	"sync"
)

type MemoryRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]string),
	}
}

func (r *MemoryRepository) Save(short, original string) error {
	r.mu.Lock()
	r.data[short] = original
	r.mu.Unlock()

	return nil
}

func (r *MemoryRepository) Get(short string) (string, error) {
	r.mu.RLock()
	val, ok := r.data[short]
	r.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("there is no short URL for %s", short)
	}

	return val, nil
}
