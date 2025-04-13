package save

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/slog"
	"url-shortener/internal/lib/random"
	"url-shortener/internal/storage"
)

const aliasLength = 6

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(url, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		getReqID := middleware.GetReqID(r.Context())

		log = log.With(slog.String("op", op), slog.String("request_id", getReqID))

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", sl.Err(err))

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err = validator.New().Struct(req); err != nil {
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)

			log.Error("invalid request", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomAlias(aliasLength)
		}

		err = urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			w.WriteHeader(http.StatusConflict) // 409
			render.JSON(w, r, resp.Error("url already exists"))

			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) // 500
			render.JSON(w, r, resp.Error("failed to save URL"))

			return
		}

		w.WriteHeader(http.StatusCreated) // 201
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
