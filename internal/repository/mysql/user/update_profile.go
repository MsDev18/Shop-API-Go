package user

import (
	"context"
	"database/sql"
	"errors"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
)

func (r Repository) UpdateProfile(ctx context.Context, user entity.User) error {
	const op = "user-repository.UpdateProfile"

	const query = `UPDATE user SET name = ?, avatar = ?, password = ? WHERE id = ?`

	_, err := r.connection.DB.ExecContext(ctx, query, user.Name, user.Avatar, user.Password, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return richerror.New().
				SetOp(op).
				SetMsg("not found user").
				SetKind(richerror.KindNotFoundErr).
				SetErr(err)
		}
		return richerror.New().
			SetOp(op).
			SetMsg("unexpected error in update user record").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	return nil
}
