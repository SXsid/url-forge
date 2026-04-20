package internal

import (
	"github/SXsid/url-forge/internal/api/handler"
	"github/SXsid/url-forge/internal/config"
	"github/SXsid/url-forge/internal/logger"
	repository "github/SXsid/url-forge/internal/repository/postgres"
	redisChache "github/SXsid/url-forge/internal/repository/redis"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Applicaton struct {
	Config     *config.ServerConfig
	Logger     *logger.Logger
	UrlHandler *handler.URLHandler
	Pgpool     *pgxpool.Pool
	rdc        *redis.Client
}

// merge all
func NewApplication(logger *logger.Logger, config *config.ServerConfig, pg *pgxpool.Pool, rdc *redis.Client) *Applicaton {
	repo := repository.NewUrlRepository(pg)
	chache := redisChache.NewRedisCache(rdc)
	handler := handler.NewURLHandler(repo, logger, chache)
	return &Applicaton{
		Logger:     logger,
		Config:     config,
		UrlHandler: handler,
		Pgpool:     pg,
		rdc:        rdc,
	}
}
