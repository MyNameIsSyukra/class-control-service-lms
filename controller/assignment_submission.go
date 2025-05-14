package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	AssignmentSubmissionController interface {
		CreateAssignmentSubmission(ctx *gin.Context)
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

	assignmentID := ctx.Param("assignment_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	req.AssignmentID = parsedAssignmentID
	assignmentSubmission, err := controller.assignmentSubmissionService.CreateAssignmentSubmission(ctx.Request.Context(), req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(assignmentSubmission)
	ctx.JSON(200, res)
}