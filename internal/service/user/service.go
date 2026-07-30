package user

import (
	"context"
	"shop/internal/entity"
)

type Service struct {
	repository Repository
}

type Repository interface {
	GetUserByID(ctx context.Context , userID uint) (entity.User, error)
}

func New (repository Repository) Service {
	return Service{
		repository: repository,
	}
}