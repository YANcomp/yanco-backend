package service

import (
	"context"
	"github.com/YANcomp/yanco-backend/internal/domain"
	"github.com/YANcomp/yanco-backend/internal/repository"
)

type Banners interface {
	GetAll(ctx context.Context, getsQuery domain.GetsQuery) ([]domain.Banner, error)
}

type Services struct {
	Banners Banners
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	bannersService := NewBannersService(deps.Repos.Banners)

	return &Services{
		Banners: bannersService,
	}
}
