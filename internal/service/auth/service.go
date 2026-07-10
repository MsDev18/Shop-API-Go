package auth

import (
	"context"
	"shop/internal/entity"
)

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	CreateOtp(ctx context.Context, otp entity.Otp) (entity.Otp, error)
	UpdateOtp(ctx context.Context, otp entity.Otp) error
}

type Service struct {
	repository Repository
}

func New(repository Repository) Service {
	return Service{
		repository: repository,
	}
}
