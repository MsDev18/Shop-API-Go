package auth

import (
	"context"
	"errors"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
)

var OtpNotFoundErr = errors.New("otp not found")

func (r Repository) UpdateOtp(ctx context.Context, otp entity.Otp) error {
	const op = "auth-repository.UpdateOtp"

	const query = `UPDATE otp SET code = ? , expires_at = ? WHERE user_id = ?`

	result, err := r.connection.DB.ExecContext(ctx, query, otp.Code, otp.ExpiresAt, otp.UserID)
	if err != nil {
		return richerror.New().
			SetOp(op).
			SetMsg("unexpected error").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New().
			SetOp(op).
			SetMsg("unexpected error").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	if rowsAffected == 0 {
		return richerror.New().
			SetOp(op).
			SetMsg("otp not found").
			SetKind(richerror.KindNotFoundErr).
			SetErr(OtpNotFoundErr)
	}

	return nil
}
