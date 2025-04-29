package controller

import (
	"LMSGo/dto"
	kelas "LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	MemberController interface {
		AddMemberToClass(ctx *gin.Context)
		GetAllMembers(ctx *gin.Context)
		// GetMemberById(ctx *gin.Context)q
		// UpdateMember(ctx *gin.Context)
		DeleteMember(ctx *gin.Context)
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
		ctx.JSON(500, gin.H{"error": "Failed to add member to class"})
		return
	}
	ctx.JSON(200, member)
}

func (controller *memberController) GetAllMembers(ctx *gin.Context) {
	members, err := controller.memberService.GetAllMembers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get members"})
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

