package service

import (
	"context"
	"github.com/YANcomp/platform_common/pkg/db/pg/pgquery"
	"github.com/YANcomp/yanco-backend/internal/domain"
	"github.com/YANcomp/yanco-backend/internal/repository"
)

type BannersService struct {
	repo repository.Banners
}

func NewBannersService(repo repository.Banners) *BannersService {
	return &BannersService{repo: repo}
}

func (s *BannersService) GetAll(ctx context.Context) ([]domain.Banner, error) {
	// add select params
	var selectQ pgquery.SelectQuery
	selectQ.SetField("header")
	selectQ.SetField("description")

	// add filter params
	var mainFilter pgquery.Filter
	mainFilter = pgquery.NewFilter("id", pgquery.FilterTypeEQ, 1)

	return s.repo.GetAll(ctx, &selectQ, &mainFilter, &pgquery.PaginationQuery{}, &pgquery.Sort{})
}
