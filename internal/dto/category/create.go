package category

import "mime/multipart"

type CreateRequest struct {
	ParentID *uint
	Title string
	Slug string
	Image *multipart.FileHeader
}