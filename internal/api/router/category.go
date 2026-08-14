package router

import "shop/internal/entity"

func (r Router) registerCategoryRoute () {
	categoryG := r.engine.Group("/category")
	categoryG.POST("" , r.authMiddleware.AuthRequired(), r.authMiddleware.RoleRequired(entity.AdminRole) , r.categoryHandler.Create)
	categoryG.GET("" , r.categoryHandler.GetAll)
}