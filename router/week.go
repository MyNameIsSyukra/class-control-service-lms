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
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/weekly-section/:id", weekController.GetWeekByID)
		routes.GET("/weekly-section/class/:class_id", weekController.GetAllWeekByClassID)
	}
}