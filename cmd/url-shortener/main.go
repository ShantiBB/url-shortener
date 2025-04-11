package main

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"url-shortener/cmd/internal/config"
	mwLogger "url-shortener/cmd/internal/http-server/middleware/logger"
	sl "url-shortener/cmd/internal/lib/logger/slog"
	"url-shortener/cmd/internal/repository"
	svc "url-shortener/cmd/internal/service"
	"url-shortener/cmd/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgres.New(cfg.StorageURL)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	service := svc.NewURLService(&repository.Storage{Storage: storage}, log)

	if err = service.SaveURL("https://google_delete.com", "google_delete"); err != nil {
		os.Exit(1)
	}

	if err = service.GetURL("google_delete"); err != nil {
		os.Exit(1)
	}

	if err = service.DeleteURL("google_delete"); err != nil {
		os.Exit(1)
	}

	if err = service.SaveURL("https://google.com", "google"); err != nil {
		os.Exit(1)
	}

	if err = service.GetURL("google"); err != nil {
		os.Exit(1)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(mwLogger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// TODO: run server
}
