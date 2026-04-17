package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlRepository struct {
	db *pgxpool.Pool
}

func NewUrlRepository(pg *pgxpool.Pool) *UrlRepository {
	return &UrlRepository{
		db: pg,
	}
}

func (r *UrlRepository) GetURL(ctx context.Context, code string) (string, error) {
	return "", nil
}

func (r *UrlRepository) GetCode(ctx context.Context, url string) (string, error) {
	return "", nil
}

func (r *UrlRepository) Insert(ctx context.Context, url, code string) error {
	return nil
}

func (r *UrlRepository) UpdateAnylitics(ctx context.Context, code string) error {
	return nil
}
