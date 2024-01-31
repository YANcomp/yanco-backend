package repository

import (
	"context"
	"github.com/YANcomp/platform_common/pkg/db"
	"github.com/YANcomp/platform_common/pkg/db/pg/pgquery"
	"github.com/YANcomp/yanco-backend/internal/domain"
)

type Banners interface {
	GetAll(ctx context.Context, selectQuery *pgquery.SelectQuery, filterQuery *pgquery.Filter, pg *pgquery.PaginationQuery, sorts ...*pgquery.Sort) ([]domain.Banner, error)
}

type Repositories struct {
	Banners Banners
}

func NewRepositories(db db.Client) *Repositories {
	return &Repositories{
		Banners: NewBannersRepo(db),
	}
}
