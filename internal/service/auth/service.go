package auth

import "context"

type Repository interface {
	IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error)
}

type Service struct {
	repository Repository
}


func New(repository Repository) Service {
	return Service{
		repository: repository,
	}
}