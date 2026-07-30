package user

import (
	userservice "shop/internal/service/user"
	uservalidator "shop/internal/validator/user"
	userrepository "shop/internal/repository/mysql/user"
)

type Handler struct {
	repository userrepository.Repository 
	service userservice.Service
	validator uservalidator.Validator
}

func New (repository userrepository.Repository, service userservice.Service , validator uservalidator.Validator) Handler {
	return Handler{
		repository: repository,
		service: service,
		validator: validator,
	}
}