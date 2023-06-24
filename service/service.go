package service

import (
	"context"
	"gomicroservice/repository"
)

type Service interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	return s.repository.Get(ctx, key)
}

func (s *service) Set(ctx context.Context, key string, value string) error {
	return s.repository.Set(ctx, key, value)
}
