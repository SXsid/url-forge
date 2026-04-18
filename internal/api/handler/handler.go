package handler

import (
	"context"

	"github/SXsid/url-forge/internal/logger"
	"github/SXsid/url-forge/internal/repository"
)

type URLHandler struct {
	repo    URLRepository
	loggeer *logger.Logger
}

type URLRepository interface {
	GetURL(ctx context.Context, code string) (string, error)
	GetCode(ctx context.Context, url string) (string, error)
	InsertURL(ctx context.Context, tx repository.Transaction, url string) (int64, error)
	UpdateCode(ctx context.Context, tx repository.Transaction, id int64, code string) error
	BeginTx(ctx context.Context) (repository.Transaction, error)
	UpdateAnylitics(ctx context.Context, code string) error
}

func NewURLHandler(repo URLRepository, logger *logger.Logger) *URLHandler {
	return &URLHandler{
		repo:    repo,
		loggeer: logger,
	}
}
