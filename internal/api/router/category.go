package router

func (r Router) registerCategoryRoute () {
	categoryG := r.engine.Group("/category")
	categoryG.GET("/check" , r.categoryHandler.Check)
}