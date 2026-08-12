package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Check (ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{
		"message" : "test",
	})
}