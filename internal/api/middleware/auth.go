package middleware

import (
	"context"
	"shop/internal/entity"
	"shop/internal/pkg/claims"
	"shop/internal/pkg/response"
	"shop/internal/pkg/richerror"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	USER_ID_KEY    = "user-id"
	SESSION_ID_KEY = "session-id"
)

type AuthMiddleware struct {
	accessTokenSecret string
	repository        Repository
}

type Repository interface {
	GetSessionByID(ctx context.Context, sessionID uint) (entity.Session, error)
}

func NewAuthMiddileware(accessTokenSecret string, repository Repository) AuthMiddleware {
	return AuthMiddleware{
		accessTokenSecret: accessTokenSecret,
		repository:        repository,
	}
}

func (a AuthMiddleware) AuthRequired() gin.HandlerFunc {
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

		session, err := a.repository.GetSessionByID(ctx, accessClaims.SessionID)
		if err != nil {
			response.New(ctx).Error(err)
			ctx.Abort()
			return 
		}

		if session.RevokeAt != nil || !session.ExpiresAt.After(time.Now()) {
			response.New(ctx).Error(richerror.New().
				SetOp(op).
				SetMsg("session revoked or expired").
				SetKind(richerror.KindUnauthorizeErr),
			)
			ctx.Abort()
			return
		}

		ctx.Set(SESSION_ID_KEY, session.ID)
		ctx.Set(USER_ID_KEY, userID)

		ctx.Next()
	}
}
