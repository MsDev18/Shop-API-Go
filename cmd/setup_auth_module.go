package main

import (
	handler "shop/internal/api/handler/auth"
	repository "shop/internal/repository/mysql/auth"
	service "shop/internal/service/auth"
	validator "shop/internal/validator/auth"
)

func SetupAuthModule(authRepostory repository.Repository, cfg service.Config) handler.Handler {
	service := service.New(authRepostory, cfg)
	validator := validator.New()
	handler := handler.New(service, validator)
	return handler
}