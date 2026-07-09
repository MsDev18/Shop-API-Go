package auth

import (
	"context"
	"fmt"
	authdto "shop/internal/dto/auth"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
	"time"
)

func (s Service) SendOtp(ctx context.Context, req authdto.SendOtpRequest) (authdto.SendOtpResponse, error) {
	const op = "auth-service.SendOtp"
	// 1. get user from database
	var user entity.User
	var err error
	if user, err = s.repository.GetUserByPhoneNumber(ctx, req.PhoneNumber); err != nil {
		richErr, ok := err.(*richerror.RichError)
		if !ok || richErr.GetKind() != richerror.KindNotFoundErr {
			return authdto.SendOtpResponse{}, richerror.New().
				SetOp(op).
				SetMsg("unexpected error").
				SetKind(richerror.KindUnexpectedErr).
				SetErr(err)
		}
		u := entity.User{PhoneNumber: req.PhoneNumber}
		user, err = s.repository.CreateUser(ctx, u)
		if err != nil {
			// beacuse in previous step
			// we returned error as type richerror
			return authdto.SendOtpResponse{}, err
		}
	}
	fmt.Println(user)
	// if exist generate new otp code
	otp := entity.Otp{
		ID:        0,
		UserID:    user.ID,
		Code:      "",
		ExpiresAt: time.Now().Add(time.Second * 2),
	}
	s.repository.CreateOtp(ctx, otp)
	// if not exist generaete new user and otp code
	return authdto.SendOtpResponse{}, nil
}
