package http

import (
	"github.com/YANcomp/yanco-backend/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(cfg *config.Config) *chi.Mux {
	// Init chi handler
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		corsMiddleware,
	)

	// Init router
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	return router
}
