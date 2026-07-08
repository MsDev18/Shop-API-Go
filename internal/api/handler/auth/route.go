package auth

import "github.com/gin-gonic/gin"

func (h Handler) RegisterRoutes (e *gin.Engine) {
	authGroup := e.Group("/auth")
	authGroup.POST("/send-otp" , h.SendOtp)
}