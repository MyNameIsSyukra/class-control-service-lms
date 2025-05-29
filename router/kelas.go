package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func Kelas(route *gin.Engine, injector *do.Injector){
	kelasController := do.MustInvoke[controller.KelasController](injector)

	routes := route.Group("/kelas")
	{
		routes.GET("", kelasController.GetById)
	}
	routes = route.Group("/kelas/admin")
	{
		routes.POST("", kelasController.Create)
		routes.GET("", kelasController.GetAll)
		routes.PUT("", kelasController.Update)
		routes.DELETE("", kelasController.Delete)
	}
}