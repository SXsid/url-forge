package internal

import (
	"github/SXsid/url-forge/internal/api/handler"
	"github/SXsid/url-forge/internal/config"
	"github/SXsid/url-forge/internal/logger"
	repository "github/SXsid/url-forge/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Applicaton struct {
	Config     *config.ServerConfig
	Logger     *logger.Logger
	UrlHandler *handler.URLHandler
	Pgpool     *pgxpool.Pool
}

// merge all
func NewApplication(logger *logger.Logger, config *config.ServerConfig, pg *pgxpool.Pool) *Applicaton {
	repo := repository.NewUrlRepository(pg)
	handler := handler.NewURLHandler(repo, logger)
	return &Applicaton{
		Logger:     logger,
		Config:     config,
		UrlHandler: handler,
		Pgpool:     pg,
	}
}
