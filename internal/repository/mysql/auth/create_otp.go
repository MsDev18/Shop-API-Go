package auth

import (
	"context"
	"errors"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"

	"github.com/go-sql-driver/mysql"
)

func (r Repository) CreateOtp(ctx context.Context, otp entity.Otp) (entity.Otp, error) {
	const op = "auth-repository.CreateOtp"
	
	query := `INSERT INTO otp(user_id , code , expires_at) VALUES (?,?,?)`

	result, err := r.connection.DB.ExecContext(ctx, query, otp.UserID, otp.Code, otp.ExpiresAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return entity.Otp{}, richerror.New().
				SetOp(op).
				SetMsg("otp record with this user_id already exist").
				SetKind(richerror.KindConflictErr).
				SetErr(err)
		}
		return entity.Otp{}, richerror.New().
			SetOp(op).
			SetMsg("unexpected error").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	// set otp id 
	otpId, _ := result.LastInsertId()
	otp.ID = uint(otpId)

	return otp, nil
}
