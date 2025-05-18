package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	AssignmentSubmissionController interface {
		CreateAssignmentSubmission(ctx *gin.Context)
		GetAllStudentAssignmentSubmissionByAssignmentID(ctx *gin.Context)
	}
	assignmentSubmissionController struct {
		assignmentSubmissionService Assignment.AssignmentSubmissionService
	}
)

func NewAssignmentSubmissionController(assignmentSubmissionService Assignment.AssignmentSubmissionService) AssignmentSubmissionController {
	return &assignmentSubmissionController{assignmentSubmissionService}
}

func (controller *assignmentSubmissionController) CreateAssignmentSubmission(ctx *gin.Context) {
	var req dto.AssignmentSubmissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	assignmentSubmission, err := controller.assignmentSubmissionService.CreateAssignmentSubmission(ctx.Request.Context(), req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignmentSubmission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) GetAllStudentAssignmentSubmissionByAssignmentID(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submissions, err := controller.assignmentSubmissionService.GetAllStudentAssignmentSubmissionByAssignmentID(ctx.Request.Context(), parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submissions)
	ctx.JSON(200, res)
}

