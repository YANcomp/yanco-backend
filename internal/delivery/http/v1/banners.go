package v1

import (
	"encoding/json"
	"github.com/YANcomp/yanco-backend/internal/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func (h *Handler) initBannersRoutes(api chi.Router) {
	api.Route("/banners", func(r chi.Router) {
		r.Get("/", h.GetBanners)
	})
	api.Get("/{select}@banners", h.GetBanners)
}

func (h *Handler) GetBanners(w http.ResponseWriter, r *http.Request) {
	getsQuery := domain.GetsQuery{}

	//search select params
	selectParam := chi.URLParam(r, "select")
	if selectParam != "" {
		getsQuery.SelectQuery.Selects = strings.Split(selectParam, ",")
	}

	// search pagination in context
	pgLimitFromCtx, ok := r.Context().Value(pgLimitCtx).(string)
	if ok {
		getsQuery.PaginationQuery.Limit = pgLimitFromCtx
	}
	pgOffsetFromCtx, ok := r.Context().Value(pgOffsetCtx).(string)
	if ok {
		getsQuery.PaginationQuery.Offset = pgOffsetFromCtx
	}

	//search filters in query
	getsQuery.FiltersQuery.Filters = r.URL.Query()

	banners, err := h.services.Banners.GetAll(r.Context(), getsQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banners)
}
