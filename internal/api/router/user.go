package router

func (r Router) registerUserRoute () {
	userG := r.engine.Group("/user")

	userG.GET("/profile" ,r.authMiddleware.AuthRequired() ,r.userHandler.Profile)
}