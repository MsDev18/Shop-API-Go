package router

import (
	"shop/internal/api/handler/health"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router      *gin.Engine
	Health health.Handler
}

func New(router *gin.Engine) Router {
	return Router{
		Health: health.New(),
		router: router,
	}
}

func (r Router) Register () {
	r.router.GET("/health-check" , r.Health.HealthCheck)
}