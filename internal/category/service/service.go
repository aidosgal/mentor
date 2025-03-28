package service

import (
	"context"
	"log/slog"

	"github.com/aidosgal/mentor/internal/category/data"
	"github.com/aidosgal/mentor/internal/category/repository"
)

type Service interface {
	List(ctx context.Context) ([]*data.Category, error)
}

type service struct {
	log        *slog.Logger
	repository repository.Repository
}

func NewService(log *slog.Logger, repository repository.Repository) Service {
	return &service{
		log:        log,
		repository: repository,
	}
}

func (s *service) List(ctx context.Context) ([]*data.Category, error) {
	return s.repository.List(ctx)
}
