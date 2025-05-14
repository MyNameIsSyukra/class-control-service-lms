package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"

	"github.com/gin-gonic/gin"
)

type (
	AssignmentController interface {
		CreateAssignment(ctx *gin.Context)
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
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(200, res)
}
