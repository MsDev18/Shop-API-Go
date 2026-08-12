package category

import (
	categoryservice "shop/internal/service/category"
	categoryvalidator "shop/internal/validator/category"

)

type Handler struct {
	service categoryservice.Service
	validator categoryvalidator.Validator
}

func New (service categoryservice.Service, validator categoryvalidator.Validator) Handler {
	return Handler{
		service: service,
		validator: validator,
	}
}