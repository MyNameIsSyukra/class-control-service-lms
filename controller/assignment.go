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
	AssignmentController interface {
		CreateAssignment(ctx *gin.Context)
		GetAssignmentByID(ctx *gin.Context)
	}
	assignmentController struct {
		assignmentService Assignment.AssignmentService
	}
)

func NewAssignmentController(assignmentService Assignment.AssignmentService) AssignmentController {
	return &assignmentController{assignmentService}
}

func (controller *assignmentController) CreateAssignment(ctx *gin.Context) {
	var req dto.CreateAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	assignment, err := controller.assignmentService.CreateAssignment(ctx.Request.Context(), req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(200, res)
}

func (controller *assignmentController) GetAssignmentByID(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	assignment, err := controller.assignmentService.GetAssignmentByID(ctx.Request.Context(), parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(200, res)
}