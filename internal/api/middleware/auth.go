package middleware

import (
	"shop/internal/pkg/response"
	"shop/internal/pkg/richerror"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	const op = "auth-middleware"
	return func(ctx *gin.Context) {
		// get bearer token form autorization
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			response.New(ctx).Error(richerror.New().
				SetOp(op).
				SetMsg("login required").
				SetKind(richerror.KindUnauthorizeErr),
			)
			ctx.Abort()
			return 
		}

		
	}
}
