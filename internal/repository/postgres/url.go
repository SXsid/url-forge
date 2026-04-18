package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transcation : %v", err)
	}
	defer tx.Rollback(ctx)
	tx.Exec(ctx, `INSERT INTO url (code,original_url) VALUES ($1,$2) RETURNING id`)
	tx.Exec(ctx, `UPDATE url SET `)
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting transcation :%w", err)
	}

	return nil
}

func (r *UrlRepository) UpdateAnylitics(ctx context.Context, code string) error {
	return nil
}
