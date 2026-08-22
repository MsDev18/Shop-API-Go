package router

func (r Router) registerAddressRoute() {
	addressG := r.engine.Group("/address")
	
	addressG.POST("", r.authMiddleware.AuthRequired(), r.addressHandler.Create)
	addressG.GET("" , r.authMiddleware.AuthRequired(), r.addressHandler.GetAll)
	addressG.GET("/:id" , r.authMiddleware.AuthRequired() , r.addressHandler.GetOne)
}
