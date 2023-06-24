package repository

import (
	"context"
	"errors"
	"sync"
)

type Repository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type repository struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewRepository() Repository {
	return &repository{store: make(map[string]string)}
}

func (r *repository) Get(ctx context.Context, key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if value, ok := r.store[key]; ok {
		return value, nil
	}

	return "", errors.New("key not found")
}

func (r *repository) Set(ctx context.Context, key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Set the value for the given key
	r.store[key] = value
	return nil
}
