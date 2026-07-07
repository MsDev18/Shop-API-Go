package router

import (

	"github.com/gin-gonic/gin"
)

type RegisterRoutes interface {
	RegisterRoutes(e *gin.Engine)
}
type Router struct {
	engine *gin.Engine
	// handlers statements
	handlers []RegisterRoutes
}

func New(engine *gin.Engine ,handlers ...RegisterRoutes) Router {
	return Router{
		engine: engine,
		// handlers statements
		handlers: handlers,
	}
}

func (r Router) Register() {
	for _ , handler := range r.handlers {
		handler.RegisterRoutes(r.engine)
	}
}
