package handler

import (
	"context"

	"github/SXsid/url-forge/internal/logger"
)

type URLHandler struct {
	repo    URLRepository
	loggeer *logger.Logger
}

type URLRepository interface {
	GetURL(ctx context.Context, code string) (string, error)
	GetCode(ctx context.Context, url string) (string, error)
	Insert(ctx context.Context, url, code string) error
	UpdateAnylitics(ctx context.Context, code string) error
}

func NewURLHandler(repo URLRepository, logger *logger.Logger) *URLHandler {
	return &URLHandler{
		repo:    repo,
		loggeer: logger,
	}
}
