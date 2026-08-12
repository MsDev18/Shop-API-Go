package main

import (
	handler "shop/internal/api/handler/user"
	"shop/internal/pkg/imageprocessor"
	repository "shop/internal/repository/mysql/user"
	service "shop/internal/service/user"
	validator "shop/internal/validator/user"
)

func SetupUserModule(userRepository  repository.Repository, imageProcessor imageprocessor.Processor) handler.Handler {
	service  := service.New(userRepository, imageProcessor)
	validator := validator.New()
	handler := handler.New(service , validator)
	return handler
}