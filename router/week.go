package routes

import (
	"LMSGo/controller"
	"LMSGo/middleware"
	"LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func WeekSection(route *gin.Engine, injector *do.Injector){
	weekController := do.MustInvoke[controller.WeekController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	
	routes := route.Group("teacher/kelas")
	{
		routes.POST("/weekly-section",middleware.Authenticate(jwtService) , weekController.CreateWeeklySection)
		routes.PUT("/weekly-section", middleware.Authenticate(jwtService) ,weekController.UpdateWeeklySection)
		routes.DELETE("/weekly-section",middleware.Authenticate(jwtService) , weekController.DeleteWeeklySection)
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/weekly-section", middleware.Authenticate(jwtService) ,weekController.GetWeekByID)
		routes.GET("/weekly-section/class",middleware.Authenticate(jwtService), weekController.GetAllWeekByClassID)
	}
}