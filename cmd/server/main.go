package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github/SXsid/url-forge/internal"
	"github/SXsid/url-forge/internal/api/router"
	"github/SXsid/url-forge/internal/cache"
	"github/SXsid/url-forge/internal/config"
	"github/SXsid/url-forge/internal/db"
	"github/SXsid/url-forge/internal/logger"
)

func setupDB(ctx context.Context, postgresCfg config.DBConfig) (*pgxpool.Pool, error) {
	pgPool, err := db.NewPool(ctx, postgresCfg)
	if err != nil {
		return nil, err
	}
	if err := db.Test(ctx, pgPool); err != nil {
		return nil, err
	}

	return pgPool, nil
}

func main() {
	port := flag.Int("port", 8080, "port")
	flag.Parse()
	ctx := context.Background()
	config := config.NewServerConfig()
	logger := logger.NewLogger()

	pgpool, err := setupDB(ctx, config.Database)
	if err != nil {
		fmt.Printf("db is not connecteed:%v\n", err)
		os.Exit(-1)

	}
	rdb, err := cache.RedisClient(ctx, config.Redis.DSN)
	if err != nil {
		fmt.Printf("reids is not connected :w,\n", err)
		os.Exit(-1)
	}
	app := internal.NewApplication(logger, config, pgpool, rdb)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      router.NewRouter(app),
		IdleTimeout:  time.Duration(config.IdealTime) * time.Second,
		WriteTimeout: time.Duration(config.WriteTime) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTime) * time.Second,
	}

	go func() {
		fmt.Printf("server is up and running at %d\n", *port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Print(err)
		}
	}()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	<-ctx.Done()
	fmt.Println("gracefully shutting down server")
	timectx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(timectx); err != nil {
		fmt.Print(err)
	}
	fmt.Println("server closed")
}
