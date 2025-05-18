package controller

import (
	"LMSGo/dto"
	kelas "LMSGo/service"
	"LMSGo/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	MemberController interface {
		AddMemberToClass(ctx *gin.Context)
		GetAllMembersByClassIDData(ctx *gin.Context)
		// GetMemberById(ctx *gin.Context)q
		// UpdateMember(ctx *gin.Context)
		DeleteMember(ctx *gin.Context)
		GetAllClassAndAssesmentByUserID(ctx *gin.Context)
		GetAllClassByUserID(ctx *gin.Context)
		
		// Lintas Service
		GetAllMembersByClassID(ctx *gin.Context)
	}
	memberController struct {
		memberService kelas.MemberService
	}
)

func NewMemberController(memberService kelas.MemberService) MemberController {
	return &memberController{memberService}
}


func (controller *memberController) AddMemberToClass(ctx *gin.Context) {
	var req dto.AddMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	member, err := controller.memberService.AddMemberToClass(ctx.Request.Context(), &req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(member)
	ctx.JSON(200, res)
}


func (controller *memberController) DeleteMember(ctx *gin.Context) {
	id := ctx.Query("id")
	// Assuming id is a UUID, you might want to parse it here
	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	if err := controller.memberService.DeleteMember(ctx.Request.Context(), parsedID); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, gin.H{"message": "Member deleted successfully"})
}

func (controller *memberController) GetAllClassAndAssesmentByUserID(ctx *gin.Context) {
	userID := ctx.Query("userID")
	// Assuming userID is a UUID, you might want to parse it here
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	classes, err := controller.memberService.GetAllClassAndAssesmentByUserID(ctx.Request.Context(), parsedUserID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(classes)
	ctx.JSON(200, res)
}

func (controller *memberController) GetAllClassByUserID(ctx *gin.Context) {
	userID := ctx.Query("userID")
	// Assuming userID is a UUID, you might want to parse it here
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	classes, err := controller.memberService.GetAllClassByUserID(ctx.Request.Context(), parsedUserID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	res := utils.SuccessResponse(classes)
	ctx.JSON(200, res)
}

func (controller *memberController) GetAllMembersByClassIDData(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Query("classID"))
	if err != nil {
		res := utils.FailedResponse("Invalid class ID format")
		ctx.JSON(400, res)
		return
	}
		
	members, err := controller.memberService.GetAllMembersByClassID(ctx.Request.Context(), classID)
	if err != nil {
		res := utils.FailedResponse("Failed to get members")
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(members)
	ctx.JSON(200, res)
}


// lintas Servicelintas Servicelintas Servicelintas Servicelintas Servicelintas Servicelintas Servicelintas Servicelintas Service
func (controller *memberController) GetAllMembersByClassID(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Param("classID"))
	if err != nil {
		res := utils.FailedResponse("Invalid class ID format")
		ctx.JSON(400, res)
		return
	}
		
	members, err := controller.memberService.GetAllMembersByClassID(ctx.Request.Context(), classID)
	if err != nil {
		res := utils.FailedResponse("Failed to get members")
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, members)
}