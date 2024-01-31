package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) initBannersRoutes(api chi.Router) {
	api.Route("/banners", func(r chi.Router) {
		r.Get("/", h.GetBanners)
	})
}

func (h *Handler) GetBanners(w http.ResponseWriter, r *http.Request) {
	banners, err := h.services.Banners.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banners)
}
