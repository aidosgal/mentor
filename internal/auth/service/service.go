package service

import (
	"context"
	"log/slog"

	"github.com/aidosgal/mentor/internal/auth/data"
	"github.com/aidosgal/mentor/internal/auth/repository"
)

type Service interface {
	Create(ctx context.Context, user *data.UserModel) (int64, error)
}

type service struct {
	log *slog.Logger
	repository repository.Repository
}

func NewService(log *slog.Logger, repository repository.Repository) Service {
	return &service{
		log: log,
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, user *data.UserModel) (int64, error) {
	return s.repository.Create(ctx, user)
}
