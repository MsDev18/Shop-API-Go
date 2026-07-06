package auth

import (
	authRepository "shop/internal/repository/mysql/auth"
	authservice "shop/internal/service/auth"
	authvalidator "shop/internal/validator/auth"
)

type Handler struct {
	service    authservice.Service
	validator  authvalidator.Validator
	repository authRepository.Repository
}


func New (repository authRepository.Repository, service authservice.Service, validator authvalidator.Validator) Handler {
	return Handler{
		service:    service,
		validator:  validator,
		repository: repository,
	}
}