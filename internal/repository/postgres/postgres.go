package postgres

import (
	"authTest/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	err := p.db.QueryRowContext(ctx, query, user.Name, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`

	var user domain.User
	err := p.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, nil
		}

		return nil, err
	}

	return &user, nil
}

func (p *Postgres) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password_hash, created_at FROM users WHERE email = $1`

	var user domain.User
	err := p.db.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, fmt.Errorf("пользователь не найден: %w", err)
		}

		return &domain.User{}, err
	}

	return &user, nil
}
