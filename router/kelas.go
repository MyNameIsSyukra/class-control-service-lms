package routes

import (
	"LMSGo/controller"
	"LMSGo/middleware"
	"LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func Kelas(route *gin.Engine, injector *do.Injector){
	kelasController := do.MustInvoke[controller.KelasController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("/kelas")
	{
		routes.GET("",middleware.Authenticate(jwtService) , kelasController.GetById)
	}
	routes = route.Group("/kelas/admin")
	{
		routes.POST("",middleware.Authenticate(jwtService) , kelasController.Create)
		routes.GET("",middleware.Authenticate(jwtService) , kelasController.GetAll)
		routes.PUT("",middleware.Authenticate(jwtService) , kelasController.Update)
		routes.DELETE("",middleware.Authenticate(jwtService) , kelasController.Delete)
	}
}