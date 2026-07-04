package auth

import (
	authdto "shop/internal/dto/auth"
)

func (s Service) SendOtp (req authdto.SendOtpRequest) (authdto.SendOtpResponse, error) {
	return authdto.SendOtpResponse{}, nil
}