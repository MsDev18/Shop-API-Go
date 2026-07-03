package auth

import "context"

type Repository interface {
	IsPhoneNumberUnique(ctx context.Context ,phoneNumber string) (bool, error)
}

type Validator struct {
	repository Repository
}

func New (repository Repository) Validator {
	return Validator{
		repository: repository,
	}
}