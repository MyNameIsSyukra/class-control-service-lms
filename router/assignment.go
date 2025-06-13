package routes

import (
	"LMSGo/controller"
	"LMSGo/middleware"
	"LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func Assignment(route *gin.Engine, injector *do.Injector) {
	assignmentController := do.MustInvoke[controller.AssignmentController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("teacher/kelas")
	{
		routes.POST("/assignment",middleware.Authenticate(jwtService), assignmentController.CreateAssignment)
		routes.GET("/assignment",middleware.Authenticate(jwtService) , assignmentController.GetAssignmentByID)
		routes.PUT("/assignment",middleware.Authenticate(jwtService) , assignmentController.UpdateAssignment)
		routes.DELETE("/assignment",middleware.Authenticate(jwtService) , assignmentController.DeleteAssignment)
	}
	studentRoute := route.Group("student/kelas")
	{
		studentRoute.GET("/assignment",middleware.Authenticate(jwtService) , assignmentController.GetAssignmentByIDStudentID)	
	}
}
