package middleware

import (
	"shop/internal/pkg/claims"
	"shop/internal/pkg/response"
	"shop/internal/pkg/richerror"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	USER_ID_KEY = "user-id"
)

type AuthMiddleware struct {
	accessTokenSecret string
}

func NewAuthMiddileware (accessTokenSecret string) AuthMiddleware {
	return AuthMiddleware{
		accessTokenSecret: accessTokenSecret,
	}
}

func (a AuthMiddleware) AuthRequired () gin.HandlerFunc {
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

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.New(ctx).Error(richerror.New().
				SetOp(op).
				SetMsg("invalid authorization header format").
				SetKind(richerror.KindUnauthorizeErr),
			)
			ctx.Abort()
			return
		}

		token := parts[1]
		accessClaims, err := claims.ParseAccessToken(token, a.accessTokenSecret)
		if err != nil {
			response.New(ctx).Error(err)
			ctx.Abort()
			return
		}

		userID, err := strconv.ParseUint(accessClaims.Subject, 10, 64)
		if err != nil {
			response.New(ctx).Error(richerror.New().
				SetOp(op).
				SetMsg("invalid subject claim in access token").
				SetKind(richerror.KindUnauthorizeErr).
				SetErr(err),
			)
			ctx.Abort()
			return
		}

		ctx.Set(USER_ID_KEY, userID)
		ctx.Next()
	}
}
