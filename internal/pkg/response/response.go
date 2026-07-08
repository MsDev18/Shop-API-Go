package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responder struct {
	ctx *gin.Context
}

type body struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    any            `json:"data,omitempty"`
	Errors  map[string]any `json:"errors,omitempty"`
}

func New(ctx *gin.Context) *Responder {
	return &Responder{
		ctx: ctx,
	}
}

func (r *Responder) send(code int, message string, data any, errors map[string]any) {
	r.ctx.JSON(code, body{
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}

// helper
func (r *Responder) OK(message string, data any) {
	r.send(http.StatusOK, message, data, nil)
}

// helper
func (r *Responder) Created(message string, data any) {
	r.send(http.StatusCreated, message, data , nil)
}
