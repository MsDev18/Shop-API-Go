package product

import (
	"context"
	dto "shop/internal/dto/product"
	"shop/internal/pkg/richerror"
)

func (s Service) Create(ctx context.Context, req dto.CreateRequest) (dto.CreateResponse, error) {
	const op = "product-service-Create"
	// check category 
	
	// check exists slug
	exists, err := s.repository.IsExistsSlug(ctx, req.Slug)
	if err != nil {
		return dto.CreateResponse{}, err
	}
	if exists {
		return dto.CreateResponse{}, richerror.New().
			SetOp(op).
			SetMsg("conflict slug , this slug already exists").
			SetKind(richerror.KindConflictErr)
	}
	// 
}
