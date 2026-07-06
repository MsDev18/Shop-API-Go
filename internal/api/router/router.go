package router

import (
	"shop/internal/api/handler/auth"
	"shop/internal/api/handler/health"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	// handlers statements
	healthHandler health.Handler
	authHandler auth.Handler
}

func New(engine *gin.Engine , healthHandler health.Handler , authHandler auth.Handler) Router {
	return Router{
		engine: engine,
		// handlers statements
		healthHandler: healthHandler,
		authHandler: authHandler,
	}
}

func (r Router) Register() {
	r.registerHealthRoute()
}
