package auth

import (
	authdto "shop/internal/dto/auth"
	"shop/internal/pkg/response"
	"shop/internal/pkg/richerror"

	"github.com/gin-gonic/gin"
)

func (h Handler) CheckOtp(ctx *gin.Context) {
	const op = "auth-handler.CheckOtp"

	// bind data
	var req authdto.CheckOtpRequest
	if bindingErr := ctx.ShouldBindJSON(&req); bindingErr != nil {
		response.New(ctx).Error(
			richerror.New().
				SetOp(op).
				SetMsg("can't bind body of request data").
				SetKind(richerror.KindBadRequestErr).
				SetErr(bindingErr),
		)
		return
	}

	// validation data
	if validationErr := h.validator.CheckOtp(ctx.Request.Context() ,req) ; validationErr != nil {
		response.New(ctx).Error(validationErr)
		return
	}

	// call service
	_, serviceErr := h.service.CheckOtp(ctx.Request.Context(), req)
	if serviceErr != nil {
		response.New(ctx).Error(serviceErr)
		return
	}
	response.New(ctx).OK("login successfully" , nil)
}
