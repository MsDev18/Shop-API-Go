package auth

import (
	"net/http"
	authdto "shop/internal/dto/auth"
	"shop/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h Handler) SendOtp(ctx *gin.Context) {
	// bind request body to json
	var req authdto.SendOtpRequest
	// in this section 
	// type error is not richerror 
	// and if use package response 
	// return internal srever error 500
	if bErr := ctx.ShouldBindJSON(&req); bErr != nil {
		response.New(ctx).WithErr(http.StatusBadRequest, "can't bind body of request data")
		return
	}
	// validate rquest body
	if vErr := h.validator.SendOtp(ctx.Request.Context(), req); vErr != nil {
		response.New(ctx).WithRichErr(vErr)
		return
	}
	// call service
	res, sErr := h.service.SendOtp(ctx.Request.Context(), req)
	if sErr != nil {
		response.New(ctx).WithRichErr(sErr)
		return
	}
	response.New(ctx).OK("OTP sent successfully", res)
}
