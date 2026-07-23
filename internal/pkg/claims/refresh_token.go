package claims

import (
	"shop/internal/pkg/richerror"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func CreateRefreshToken(userID uint, secretKey string, duration uint) (string, error) {
	const op = "claims.CreateRefreshToken"

	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(duration)))
	claims := RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			Subject: strconv.Itoa(int(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(secretKey)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("signed jwt refresh token err").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	return tokenStr, nil
}
