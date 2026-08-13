package category

import (
	"context"
	"database/sql"
	"errors"
	"shop/internal/entity"
	"shop/internal/pkg/richerror"
)

func (r Repository) IsUniqueSlug(ctx context.Context, slug string) (bool, error) {
	const op = "category-repository.IsUniqueSlug"

	const query = `SELECT * FROM category WHERE slug = ? AND deleted_at IS NULL`

	row := r.connection.DB.QueryRowContext(ctx, query, slug)

	var c entity.Category
	if err := row.Scan(&c.ID, &c.ParentID, &c.Title, &c.Slug, &c.Image, &c.DeletedAt, &c.CreatedAt, &c.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, richerror.New().
			SetOp(op).
			SetMsg("unexpected error in repository").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	return false, nil
}
