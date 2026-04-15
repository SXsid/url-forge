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

	"github/SXsid/url-forge/internal"
	"github/SXsid/url-forge/internal/api/router"
	"github/SXsid/url-forge/internal/config"
	"github/SXsid/url-forge/internal/logger"
)

func main() {
	port := flag.Int64("p", 8080, "port for the app")
	flag.Parse()
	config, err := config.NewServerConfig()
	if err != nil {
		panic(err)
	}
	logger := logger.NewLogger()
	app := internal.NewApplication(logger, config)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      router.NewRouter(app),
		IdleTimeout:  30 * time.Second,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
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
