package address

import (
	"context"
	"shop/internal/entity"
	"shop/internal/service/province"
)

type Service struct {
	repository     Repository
	provinceService province.Service
}

type Repository interface {
	Create(ctx context.Context, address entity.Address) (entity.Address, error)
	GetAll(ctx context.Context, userID uint) ([]entity.Address, error)
}

func New(repository Repository, provinceService province.Service) Service {
	return Service{
		repository:     repository,
		provinceService: provinceService,
	}
}
