package repository

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/YANcomp/platform_common/pkg/db"
	"github.com/YANcomp/platform_common/pkg/db/pg/pgquery"
	"github.com/YANcomp/yanco-backend/internal/domain"
	"slices"
	"strings"
)

const (
	tableName = "public.banners"

	idColumn          = "id"
	headerColumn      = "header"
	descriptionColumn = "description"
	targetTypeColumn  = "targettype"
	targetColumn      = "target"
	imageColumn       = "image"
	imageMobileColumn = "imagemobile"
	isCatalogColumn   = "iscatalog"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

func getSupportedColumn() []string {
	return []string{"header", "description", "targettype", "target", "image", "imagemobile", "iscatalog"}
}

type BannersRepo struct {
	db db.Client
}

func NewBannersRepo(db db.Client) *BannersRepo {
	return &BannersRepo{
		db: db,
	}
}

func (r *BannersRepo) GetAll(ctx context.Context, selectQuery *pgquery.SelectQuery, filterQuery *pgquery.Filter, pg *pgquery.PaginationQuery, sorts ...*pgquery.Sort) ([]domain.Banner, error) {
	columns := ""
	if !selectQuery.IsEmpty() {
		for _, value := range selectQuery.GetFields() {
			if !slices.Contains(getSupportedColumn(), value) {
				return nil, errors.New("Unsupported select field: " + value)
			}
		}

		columns = strings.Join(selectQuery.GetFields(), ", ")
	} else {
		columns = strings.Join([]string{
			headerColumn,
			descriptionColumn,
			targetTypeColumn,
			targetColumn,
			imageColumn,
			imageMobileColumn,
			isCatalogColumn,
		}, ", ")
	}

	builder := sq.Select(idColumn).
		PlaceholderFormat(sq.Dollar).
		Columns(columns).
		From(tableName)

	// set filters
	if !filterQuery.IsEmpty() {
		builder = filterQuery.UseSelectBuilder(builder)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "banners_repository.Get",
		QueryRaw: query,
	}

	list := make([]domain.Banner, 0)

	err = r.db.DB().ScanAllContext(ctx, &list, q, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}
