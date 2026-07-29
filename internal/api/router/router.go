package router

import (
	"shop/internal/api/handler/auth"
	"shop/internal/api/handler/health"
	authmiddleware"shop/internal/api/middleware/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	// handlers statements
	healthHandler health.Handler
	authHandler   auth.Handler
	authMiddleware authmiddleware.Middleware
}

func New(engine *gin.Engine, healthHandler health.Handler, authHandler auth.Handler, authMiddleware authmiddleware.Middleware) Router {
	return Router{
		engine: engine,
		// handlers statements
		healthHandler: healthHandler,
		authHandler: authHandler,
		authMiddleware: authMiddleware,
	}
}


func (r Router) Register () {
	r.registerHealthRoute()
	r.registerAuthRoute()
}