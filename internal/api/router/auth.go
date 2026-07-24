package router

func (r Router) registerAuthRoute () {
	authG := r.engine.Group("/auth")
	authG.POST("/send-otp" , r.authHandler.SendOtp)
	authG.POST("/check-otp" , r.authHandler.CheckOtp)
}