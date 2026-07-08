package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responder struct {
	ctx *gin.Context
}

type body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func New(ctx *gin.Context) *Responder {
	return &Responder{
		ctx: ctx,
	}
}

func (r *Responder) Send(code int, message string, data any) {
	r.ctx.JSON(code, body{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// helper
func (r *Responder) OK(message string, data any) {
	r.Send(http.StatusOK, message, data)
}
// helper
func (r *Responder) Created(message string, data any) {
	r.Send(http.StatusCreated, message, data)
}
