package internal

import (
	"github/SXsid/url-forge/internal/api/handler"
	"github/SXsid/url-forge/internal/config"
	"github/SXsid/url-forge/internal/logger"
	repository "github/SXsid/url-forge/internal/repository/postgres"
)

type Applicaton struct {
	Config     *config.ServerConfig
	Logger     *logger.Logger
	UrlHandler *handler.URLHandler
}

// merge all
func NewApplication(logger *logger.Logger, config *config.ServerConfig) *Applicaton {
	repo := repository.NewUrlRepository()
	handler := handler.NewURLHandler(repo, logger)
	return &Applicaton{
		Logger:     logger,
		Config:     config,
		UrlHandler: handler,
	}
}
