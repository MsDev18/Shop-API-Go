package main

import (
	handler "shop/internal/api/handler/address"
	"shop/internal/repository/mysql"
	repository "shop/internal/repository/mysql/address"
	service "shop/internal/service/address"
	validator "shop/internal/validator/address"
)

func setupAddressModule(mysqlRepo mysql.Connection) handler.Handler {
	repository := repository.New(mysqlRepo)
	service := service.New(repository)
	validator := validator.New()
	return handler.New(service, validator)
}
