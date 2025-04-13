package redirect

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "url-shortener/internal/lib/api/response"
	sl "url-shortener/internal/lib/logger/slog"
	"url-shortener/internal/storage"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"
		getReqID := middleware.GetReqID(r.Context())

		log = log.With(slog.String("op", op), slog.String("request_id", getReqID))

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			w.WriteHeader(http.StatusNotFound) // 404
			render.JSON(w, r, resp.Error("not found"))

			return
		}

		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			w.WriteHeader(http.StatusNotFound) // 404
			render.JSON(w, r, resp.Error("not found"))

			return
		}

		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError) // 500
			render.JSON(w, r, resp.Error("internal server error"))

			return
		}

		log.Info("success redirect", slog.String("url", resURL))

		http.Redirect(w, r, resURL, http.StatusFound) // 302

	}
}
