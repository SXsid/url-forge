package postgres

import (
	"context"
	"errors"
	"fmt"

	"github/SXsid/url-forge/internal/domain"
	"github/SXsid/url-forge/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlPostgresRepository struct {
	db *pgxpool.Pool
}

func NewUrlRepository(pg *pgxpool.Pool) *UrlPostgresRepository {
	return &UrlPostgresRepository{
		db: pg,
	}
}

func (r *UrlPostgresRepository) BeginTx(ctx context.Context) (repository.Transaction, error) {
	return r.db.Begin(ctx)
}

func (r *UrlPostgresRepository) IsUrlExist(ctx context.Context, url string) (string, error) {
	var code string
	err := r.db.QueryRow(ctx, ` SELECT code FROM url WHERE original_url = $1`, url).Scan(&code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return code, nil
}

func (r *UrlPostgresRepository) GetURL(ctx context.Context, code string) (string, error) {
	var url string
	if err := r.db.QueryRow(ctx, `SELECT original_url FROM url WHERE code=$1`, code).Scan(&url); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", domain.ErrCodeNotFound
		}
		return "", err
	}
	return url, nil
}

func (r *UrlPostgresRepository) InsertURL(ctx context.Context, tx repository.Transaction, url string) (uint64, error) {
	var id uint64
	if err := tx.QueryRow(ctx, `INSERT INTO url (original_url ) VALUES ($1) RETURNING id`, url).Scan(&id); err != nil {
		return 0, fmt.Errorf("error create a url entry:%w", err)
	}
	return id, nil
}

func (r *UrlPostgresRepository) UpdateCode(ctx context.Context, tx repository.Transaction, id uint64, code string) error {
	commandtag, err := tx.Exec(ctx, `UPDATE url SET code=$1 where id=$2`, code, id)
	if err != nil {
		return fmt.Errorf("error updating code for id %d :%w", id, err)
	}
	if commandtag.RowsAffected() == 0 {
		return domain.ErrCodeNotFound
	}
	return nil
}

func (r *UrlPostgresRepository) UpdateAnylitics(ctx context.Context, code string) error {
	return nil
}

func (r *UrlPostgresRepository) GetAnyltics(ctx context.Context, code string) (domain.Urls, error) {
	return domain.Urls{}, nil
}
