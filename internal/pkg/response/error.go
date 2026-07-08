package response

import (
	"errors"
	"shop/internal/pkg/mapper"
	"shop/internal/pkg/richerror"
)

func (r *Responder) WithRichErr(err error) {
	var richErr *richerror.RichError
	if errors.As(err, &richErr) {
		statusCode := mapper.KindToHttpStatusCode(richErr.GetKind())
		r.Send(statusCode, richErr.GetMessage(), nil)
		return
	}
}

func (r *Responder) WithErr(code int, message string) {
	r.Send(code, message, nil)
}
