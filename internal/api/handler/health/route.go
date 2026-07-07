package health

import "github.com/gin-gonic/gin"

func (h Handler) SetRoutes (e *gin.Engine) {
	e.GET("/health-check" , h.HealthCheck)
}