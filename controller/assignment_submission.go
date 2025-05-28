package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AssignmentSubmissionController interface {
		CreateAssignmentSubmission(ctx *gin.Context)
		GetAllStudentAssignmentSubmissionByAssignmentID(ctx *gin.Context)
		GetAssignmentSubmissionByID(ctx *gin.Context)
		UpdateStudentSubmissionScore(ctx *gin.Context)
		DeleteAssignmentSubmissionByID(ctx *gin.Context)
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
	status := ctx.Query("status")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submissions, err := controller.assignmentSubmissionService.GetAllStudentAssignmentSubmissionByAssignmentID(ctx.Request.Context(),status, parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submissions)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) UpdateStudentSubmissionScore(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	score := ctx.Query("score")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	parsedScore, err := strconv.Atoi(score)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submission, err := controller.assignmentSubmissionService.UpdateStudentSubmissionScore(ctx.Request.Context(), parsedScore, parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) GetAssignmentSubmissionByID(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submission, err := controller.assignmentSubmissionService.GetAssignmentSubmissionByID(ctx.Request.Context(), parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) DeleteAssignmentSubmissionByID(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	err = controller.assignmentSubmissionService.DeleteAssignmentSubmissionByID(ctx.Request.Context(), parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(nil)
	ctx.JSON(200, res)
}