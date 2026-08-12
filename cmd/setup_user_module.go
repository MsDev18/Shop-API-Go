package main

import (
	handler "shop/internal/api/handler/user"
	"shop/internal/pkg/imageprocessor"
	"shop/internal/repository/mysql"
	repository "shop/internal/repository/mysql/user"
	service "shop/internal/service/user"
	validator "shop/internal/validator/user"
)

func SetupUserModule(mysqlRepo  mysql.Connection, imageProcessor imageprocessor.Processor) handler.Handler {
	repository := repository.New(mysqlRepo)
	service  := service.New(repository, imageProcessor)
	validator := validator.New()
	handler := handler.New(service , validator)
	return handler
}