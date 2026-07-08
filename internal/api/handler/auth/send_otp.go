package auth

import (
	authdto "shop/internal/dto/auth"
	"shop/internal/pkg/response"
	"shop/internal/pkg/richerror"

	"github.com/gin-gonic/gin"
)

func (h Handler) SendOtp(ctx *gin.Context) {
	const op = "auth-handler.SendOtp"
	// bind request body to json
	var req authdto.SendOtpRequest
	// in this section
	// we need to convert error type to richerror
	if bErr := ctx.ShouldBindJSON(&req); bErr != nil {
		response.New(ctx).Error(richerror.New().
			SetOp(op).
			SetMsg("can't bind body of request data").
			SetKind(richerror.KindBadRequestErr).
			SetErr(bErr),
		)
		return
	}
	// validate request body
	if vErr := h.validator.SendOtp(ctx.Request.Context(), req); vErr != nil {
		response.New(ctx).Error(vErr)
		return
	}
	// call service
	res, sErr := h.service.SendOtp(ctx.Request.Context(), req)
	if sErr != nil {
		response.New(ctx).Error(sErr)
		return
	}
	response.New(ctx).OK("OTP sent successfully", res)
}
