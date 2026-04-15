package repository

import "context"

type UrlRepository struct{}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{}
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
