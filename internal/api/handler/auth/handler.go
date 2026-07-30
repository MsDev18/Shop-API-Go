package auth

import (
	authrepository "shop/internal/repository/mysql/auth"
	authservice "shop/internal/service/auth"
	authvalidator "shop/internal/validator/auth"
)

type Handler struct {
	service    authservice.Service
	validator  authvalidator.Validator
	repository authrepository.Repository
}


func New (repository authrepository.Repository, service authservice.Service, validator authvalidator.Validator) Handler {
	return Handler{
		service:    service,
		validator:  validator,
		repository: repository,
	}
}