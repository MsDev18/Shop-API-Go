package auth

import (
	"context"
	"errors"
	"math/rand"
	authdto "shop/internal/dto/auth"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
	"strconv"
	"time"
)

func (s Service) SendOtp(ctx context.Context, req authdto.SendOtpRequest) (authdto.SendOtpResponse, error) {
	const op = "auth-service.SendOtp"
	// get ot create user by phone number
	user, err := s.getOrCreateUser(ctx, req.PhoneNumber)
	if err != nil {
		return authdto.SendOtpResponse{}, err
	}
	// generate otp record
	_, err = s.createOrUpdateOtp(ctx, user.ID)
	if err != nil {
		return authdto.SendOtpResponse{}, err
	}
	// TODO - implement method for send otp with SMS
	return authdto.SendOtpResponse{}, nil
}

func (s Service) getOrCreateUser(ctx context.Context, phoneNumber string) (entity.User, error) {
	const op = "auth-service.getOrCreateUser"

	var user entity.User
	var getUerErr error

	if user, getUerErr = s.repository.GetUserByPhoneNumber(ctx, phoneNumber); getUerErr != nil {
		// if error is not of type richerror
		// or kind is not NotFoundErr  return unexpected error
		var richErr *richerror.RichError
		if !errors.As(getUerErr, &richErr) || richErr.GetKind() != richerror.KindNotFoundErr {
			return entity.User{}, richerror.New().
				SetOp(op).
				SetMsg("unexpected error").
				SetKind(richerror.KindUnexpectedErr).
				SetErr(getUerErr)
		}
		// in this section err == notfound
		// create user with phone number (basic data)
		u := entity.User{PhoneNumber: phoneNumber}
		var createUserErr error
		if user, createUserErr = s.repository.CreateUser(ctx, u); createUserErr != nil {
			// beacuse in previous step
			// we returned error as type richerror
			return entity.User{}, createUserErr
		}
	}
	return user, nil
}

func (s Service) createOrUpdateOtp(ctx context.Context, userId uint) (entity.Otp, error) {
	// generate otp data
	otpData := entity.Otp{
		UserID:    userId,
		Code:      strconv.Itoa(rand.Intn(90_000) + 10_000),
		ExpiresAt: time.Now().Add(time.Second * 120),
	}
	// declare vriebale
	var otp entity.Otp
	var createOtpErr error
	// insert otp record in database
	if otp, createOtpErr = s.repository.CreateOtp(ctx, otpData) ; createOtpErr != nil{
		// if error type is richerror
		// richerror kind is KindConflictErr update otp record
		var richErr *richerror.RichError
		if errors.As(createOtpErr, &richErr) && richErr.GetKind() == richerror.KindConflictErr {
			// call update method from repository
			if updateOtpErr := s.repository.UpdateOtp(ctx, otpData) ; updateOtpErr != nil {
				// first step we created otp if come to this line
				// thats means otp exist beacuse we dont detect error
				return entity.Otp{}, updateOtpErr
			}
			// operation is successful return empty response
			return entity.Otp{}, nil
		}
		// if error type is not richerror or kind is not KindConflictErr
		return entity.Otp{}, createOtpErr
	}
	
	return otp, nil
}
