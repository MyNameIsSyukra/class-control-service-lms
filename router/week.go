package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func WeekSection(route *gin.Engine, injector *do.Injector){
	weekController := do.MustInvoke[controller.WeekController](injector)

	routes := route.Group("teacher/kelas")
	{
		routes.POST("/weekly-section", weekController.CreateWeeklySection)
		routes.PUT("/weekly-section", weekController.UpdateWeeklySection)
		routes.DELETE("/weekly-section", weekController.DeleteWeeklySection)
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/weekly-section", weekController.GetWeekByID)
		routes.GET("/weekly-section/class", weekController.GetAllWeekByClassID)
	}
}