package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"url-shortener/cmd/internal/http-server/handlers/url/save"
	mwLogger "url-shortener/cmd/internal/http-server/middleware/logger"
	"url-shortener/cmd/internal/service"
)

func URLRouter(router *chi.Mux, log *slog.Logger, service *service.URLService) {
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, service))
}
