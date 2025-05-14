package controller

import (
	"LMSGo/dto"
	"LMSGo/service"
	"LMSGo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	WeekController interface {
		CreateWeeklySection(ctx *gin.Context)
		GetAllWeekByClassID(ctx *gin.Context)
		GetWeekByID(ctx *gin.Context)
	}
	weekController struct {
		weekService service.WeekService
	}
)

func NewWeekController(weekService service.WeekService) WeekController {
	return &weekController{weekService}
}

func (controller *weekController) CreateWeeklySection(ctx *gin.Context) {
	var req dto.CreateItemPembelajaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	item, err := controller.weekService.CreateWeeklySection(ctx.Request.Context(), req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(item)
	ctx.JSON(200, res)
}

func (controller *weekController) GetAllWeekByClassID(ctx *gin.Context) {
	classID := ctx.Param("class_id")
	parsedClassID, err := uuid.Parse(classID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	weeks, err := controller.weekService.GetAllWeekByClassID(ctx.Request.Context(), parsedClassID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(weeks)
	ctx.JSON(200, res)
}

func (controller *weekController) GetWeekByID(ctx *gin.Context) {
	weekID := ctx.Param("id")
	weekIDInt, err := strconv.Atoi(weekID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	week, err := controller.weekService.GetWeekByID(ctx.Request.Context(), weekIDInt)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}

	res := utils.SuccessResponse(week)
	ctx.JSON(200, res)
}
