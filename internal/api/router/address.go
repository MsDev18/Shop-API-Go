package router

func (r Router) registerAddressRoute() {
	r.engine.POST("", r.authMiddleware.AuthRequired(), r.addressHandler.Create)
}
