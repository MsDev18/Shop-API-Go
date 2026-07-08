package auth

import (
	"context"
	authdto "shop/internal/dto/auth"
)

func (s Service) SendOtp (cxt context.Context,req authdto.SendOtpRequest) (authdto.SendOtpResponse, error) {
	return authdto.SendOtpResponse{}, nil
}