package auth

import (
	"context"
	"database/sql"
	"errors"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
)

func (r Repository) IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error) {
	const op = "auth-repository.IsPhoneNumberUnique"

	query := `SELECT * FROM user WHERE phone_number = ?`
	row := r.connection.DB.QueryRowContext(ctx, query, phoneNumber)

	var u entity.User
	err := row.Scan(&u.ID, &u.Name, &u.Avatar, &u.PhoneNumber, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, richerror.New().
			SetOp(op).
			SetMsg("unexpected error").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	return false, nil
}
