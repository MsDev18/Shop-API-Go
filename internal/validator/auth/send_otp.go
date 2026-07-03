package auth

import (
	"context"
	"errors"
	"regexp"
	authdto "shop/internal/dto/auth"
	"shop/internal/pkg/richerror"

	"github.com/go-ozzo/ozzo-validation/v4"
)

const (
	IR_MOBILE_REGEX = `^09[0-9]{9}$`
)

func (v Validator) SendOtp(ctx context.Context, req authdto.SendOtpRequest) (bool, error) {
	const op = "auth-validator.SendOtp"
	err := validation.ValidateStructWithContext(ctx, &req,
		// 1. validation phone number format
		// 2. check the database for uniqueness phone number
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(IR_MOBILE_REGEX)),
			validation.WithContext(v.validatePhoneNumber),
		),
	)

	// error handling ozzo-vlidation package
	if err != nil {
		// check type error is validation.Errors
		var validationErr validation.Errors
		if errors.As(err, &validationErr) {
			// create meta data for richerror package
			meta := make(map[string]any)
			for key, value := range validationErr {
				meta[key] = value
			}
			// return error
			return false, richerror.New().
				SetOp(op).
				SetMsg("input validation error").
				SetKind(richerror.KindBadRequestErr).
				SetMeta(meta)
		}
		return false, richerror.New().
			SetOp(op).
			SetMsg("unexpected errror").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	return true, nil
}

func (v Validator) validatePhoneNumber(ctx context.Context, value interface{}) error {
	const op = "auth-validator.validatePhoneNumber"
	// type assertion
	// convert interface{} type to string
	// we don't need validation for type assertion is ok
	// because previous step checked phone number
	phoneNumberStr, _ := value.(string)
	isUnique, err := v.repository.IsPhoneNumberUnique(ctx, phoneNumberStr)
	if err != nil {
		// in repository package created richerror
		// just wrap err
		return err
	}
	if isUnique {
		return nil
	}
	return richerror.New().
		SetOp(op).
		SetMsg("this phone number already exist").
		SetKind(richerror.KindConflictErr)
}
