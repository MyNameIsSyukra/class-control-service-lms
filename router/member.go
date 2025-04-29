package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)


func Member(server *gin.Engine, injector *do.Injector) {
	memberController := do.MustInvoke[controller.MemberController](injector)
	member := server.Group("/member")
	{
		member.POST("/", memberController.AddMemberToClass)
		member.GET("/", memberController.GetAllMembers)
		member.DELETE("/:id", memberController.DeleteMember)
	}
}