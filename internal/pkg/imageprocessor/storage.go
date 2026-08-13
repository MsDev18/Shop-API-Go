package imageprocessor

import (
	"context"
	"io"
)

type Storage interface {
	Save(ctx context.Context , filename string, data io.Reader) (string , error)
}