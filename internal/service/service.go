package service

import (
	"authTest/internal/domain"
	"context"
	"fmt"
)

type UserStore interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserService struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserService {
	return &UserService{repo: repo}
}

// nolint: forbidigo
func (s *UserService) Start() {
	fmt.Println("db запущена")
}

func (s *UserService) Register(ctx context.Context, name, email string) (*domain.User, error) {
	user := &domain.User{Name: name, Email: email}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
