package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func Assignment(route *gin.Engine, injector *do.Injector) {
	assignmentController := do.MustInvoke[controller.AssignmentController](injector)

	routes := route.Group("teacher/kelas")
	{
		routes.POST("/assignment", assignmentController.CreateAssignment)
		routes.GET("/assignment/", assignmentController.GetAssignmentByID)
	}
	studentRoute := route.Group("student/kelas")
	{
		studentRoute.GET("/assignment/", assignmentController.GetAssignmentByIDStudentID)	
	}
}
