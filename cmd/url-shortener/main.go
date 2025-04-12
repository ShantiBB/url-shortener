package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"url-shortener/internal/config"
	urlRouter "url-shortener/internal/http-server/router"
	"url-shortener/internal/lib/logger/slog"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
	"url-shortener/internal/storage/postgres"

	"github.com/go-chi/chi/v5"
)

func main() {
	configPath := filepath.Join("config", "local.yaml")
	cfg := config.MustLoad(configPath)

	log := sl.SetupLogger(cfg.Env)
	log.Debug("debug messages are enabled")

	storage, err := postgres.New(cfg.StorageURL)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	repo := repository.Storage{Storage: storage}
	svc := service.New(&repo, log)

	router := chi.NewRouter()
	urlRouter.URLRouter(router, log, svc)

	log.Info("starting url-shortener server", slog.String(
		"env", cfg.Env),
		slog.String("address", cfg.HTTPServer.Address),
	)

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))
		os.Exit(1)
	}

	log.Error("stopped server")
}
