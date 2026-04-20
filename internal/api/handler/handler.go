package handler

import (
	"context"

	"github/SXsid/url-forge/internal/domain"
	"github/SXsid/url-forge/internal/logger"
	"github/SXsid/url-forge/internal/repository"
)

type URLHandler struct {
	repo    URLRepository
	loggeer *logger.Logger
	cahche  CahceRepository
}

type CahceRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
	// SetStream() error
	// GetStream() error
}
type URLRepository interface {
	GetURL(ctx context.Context, code string) (string, error)
	IsUrlExist(ctx context.Context, url string) (string, error)
	InsertURL(ctx context.Context, tx repository.Transaction, url string) (uint64, error)
	UpdateCode(ctx context.Context, tx repository.Transaction, id uint64, code string) error
	BeginTx(ctx context.Context) (repository.Transaction, error)
	GetAnyltics(ctx context.Context, code string) (domain.Urls, error)
	UpdateAnylitics(ctx context.Context, code string) error
}

func NewURLHandler(repo URLRepository, logger *logger.Logger, chache CahceRepository) *URLHandler {
	return &URLHandler{
		repo:    repo,
		loggeer: logger,
		cahche:  chache,
	}
}

// dto
type registerDTO struct {
	Url string `json:"url"`
}
