package service

import (
	"context"
	"github.com/YANcomp/platform_common/pkg/db/pg/pgquery"
	"github.com/YANcomp/yanco-backend/internal/domain"
	"github.com/YANcomp/yanco-backend/internal/repository"
	"strings"
)

type BannersService struct {
	repo repository.Banners
}

func NewBannersService(repo repository.Banners) *BannersService {
	return &BannersService{repo: repo}
}

func (s *BannersService) GetAll(ctx context.Context, getsQuery domain.GetsQuery) ([]domain.Banner, error) {
	// add select params
	var selectQ pgquery.SelectQuery
	if getsQuery.SelectQuery.Selects != nil {
		for _, element := range getsQuery.SelectQuery.Selects {
			err := selectQ.SetField(element)
			if err != nil {
				return nil, err
			}
		}
	}

	// add filter params
	var mainFilter pgquery.Filter
	iFilters := 0
	for index, element := range getsQuery.Filters {
		for subIndex, subElem := range element {
			element[subIndex] = strings.ReplaceAll(subElem, "\"", "")
		}
		if iFilters == 0 {
			mainFilter = pgquery.NewFilter(index, pgquery.FilterTypeEQ, element)
		} else {
			filter := pgquery.NewFilter(index, pgquery.FilterTypeEQ, element)
			mainFilter.WithFilters(filter)
		}
		iFilters++
	}

	//add pagination params
	var paginationQ pgquery.PaginationQuery
	err := paginationQ.SetLimit(getsQuery.PaginationQuery.Limit)
	if err != nil {
		return nil, err
	}
	err = paginationQ.SetOffset(getsQuery.PaginationQuery.Offset)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(ctx, &selectQ, &mainFilter, &paginationQ, &pgquery.Sort{})
}
