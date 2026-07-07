package auth

import "github.com/gin-gonic/gin"

func (h Handler) SetRoutes (e *gin.Engine) {
	authGroup := e.Group("/auth")
	authGroup.GET("/login")
	authGroup.GET("/register")
}