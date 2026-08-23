package router

func (r Router) registerProductRoute() {
	_ = r.engine.Group("/product")
}
