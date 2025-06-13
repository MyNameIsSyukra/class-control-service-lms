package routes

import (
	"LMSGo/controller"
	"LMSGo/middleware"
	"LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func AssignmentSubmission(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	assignmentSubmissionController := do.MustInvoke[controller.AssignmentSubmissionController](injector)

	routes := route.Group("student/kelas")
	{
		routes.POST("/assignment-submission", middleware.Authenticate(jwtService) ,assignmentSubmissionController.CreateAssignmentSubmission)
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/assignment-submission",middleware.Authenticate(jwtService), assignmentSubmissionController.GetAllStudentAssignmentSubmissionByAssignmentID)
		routes.PUT("/assignment-submission", middleware.Authenticate(jwtService),assignmentSubmissionController.UpdateStudentSubmissionScore)
		routes.DELETE("/assignment-submission",middleware.Authenticate(jwtService), assignmentSubmissionController.DeleteAssignmentSubmissionByID)
		routes.GET("/assignment-submission/student",middleware.Authenticate(jwtService), assignmentSubmissionController.GetAssignmentSubmissionByID)
	}
}