package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (h *Handler) initBannersRoutes(api chi.Router) {
	api.Route("/banners", func(r chi.Router) {
		r.Get("/", h.GetBanners)
	})
	api.Get("/{select}@banners", h.GetBanners)
}

func (h *Handler) GetBanners(w http.ResponseWriter, r *http.Request) {
	selectParam := chi.URLParam(r, "select")
	log.Printf(selectParam)

	banners, err := h.services.Banners.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banners)
}
