package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h Handler) HealthCheck (ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{
		"message" : "Health-Check ✅",
	})
}