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
		GetAllMembersByClassID(ctx *gin.Context)
		// GetMemberById(ctx *gin.Context)q
		// UpdateMember(ctx *gin.Context)
		DeleteMember(ctx *gin.Context)
		GetAllClassAndAssesmentByUserID(ctx *gin.Context)
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
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	member, err := controller.memberService.AddMemberToClass(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(200, member)
}

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

func (controller *memberController) DeleteMember(ctx *gin.Context) {
	id := ctx.Param("id")
	// Assuming id is a UUID, you might want to parse it here
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := controller.memberService.DeleteMember(ctx.Request.Context(), parsedID); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete member"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Member deleted successfully"})
}

func (controller *memberController) GetAllClassAndAssesmentByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	// Assuming userID is a UUID, you might want to parse it here
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid User ID format"})
		return
	}
	classes, err := controller.memberService.GetAllClassAndAssesmentByUserID(ctx.Request.Context(), parsedUserID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get classes and assessments"})
		return
	}
	ctx.JSON(200, classes)
}
