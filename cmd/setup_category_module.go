package main

import (
	handler "shop/internal/api/handler/category"
	repository "shop/internal/repository/mysql/category"
	service "shop/internal/service/category"
	validator "shop/internal/validator/category"
)
	

func SetupCategoryModule (categoryRepository repository.Repository) handler.Handler {
	service := service.New(categoryRepository)
	validator := validator.New() 
	handler := handler.New(service, validator)
	return handler
}