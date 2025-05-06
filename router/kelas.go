package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func Kelas(route *gin.Engine, injector *do.Injector){
	kelasController := do.MustInvoke[controller.KelasController](injector)

	routes := route.Group("/kelas/admin")
	{
		routes.POST("", kelasController.Create)
		routes.GET("", kelasController.GetAll)
		routes.PUT("/:id", kelasController.Update)
		routes.DELETE("/:id", kelasController.Delete)
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/:id", kelasController.GetById)
	}
}