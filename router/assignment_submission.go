package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
) 

func AssignmentSubmission(route *gin.Engine, injector *do.Injector) {
	assignmentSubmissionController := do.MustInvoke[controller.AssignmentSubmissionController](injector)

	routes := route.Group("student/kelas")
	{
		routes.POST("/assignment-submission", assignmentSubmissionController.CreateAssignmentSubmission)
	}
	routes = route.Group("/kelas")
	{
		routes.GET("/assignment-submission/", assignmentSubmissionController.GetAllStudentAssignmentSubmissionByAssignmentID)
		routes.PUT("/assignment-submission", assignmentSubmissionController.UpdateStudentSubmissionScore)
		routes.GET("/assignment-submission/student", assignmentSubmissionController.GetAssignmentSubmissionByID)
	}
}