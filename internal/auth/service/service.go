package service

import (
	"context"
	"log/slog"

	"github.com/aidosgal/mentor/internal/auth/data"
	"github.com/aidosgal/mentor/internal/auth/repository"
)

type Service interface {
	Create(ctx context.Context, user *data.UserModel) (int64, bool, error)
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

func (s *service) Create(ctx context.Context, user *data.UserModel) (int64, bool, error) {
	isNewUser := true

	user, err := s.repository.Get(ctx, user.ChatID)
	if err != nil {
		return 0, true, err
	}

	if user != nil {
		isNewUser = false
	}

	var id int64
	if isNewUser {
		id, err = s.repository.Create(ctx, user)
	}

	return id, isNewUser, err
}
