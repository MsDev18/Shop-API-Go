package main

import (
	handler "shop/internal/api/handler/category"
	"shop/internal/repository/mysql"
	repository "shop/internal/repository/mysql/category"
	service "shop/internal/service/category"
	validator "shop/internal/validator/category"
)
	

func SetupCategoryModule (mysqlRepo mysql.Connection) handler.Handler {
	repository := repository.New(mysqlRepo)
	service := service.New(repository)
	validator := validator.New() 
	handler := handler.New(service, validator)
	return handler
}