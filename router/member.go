package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)


func Member(server *gin.Engine, injector *do.Injector) {
	memberController := do.MustInvoke[controller.MemberController](injector)
	member := server.Group("/member/admin")
	{
		member.POST("", memberController.AddMemberToClass)
		member.DELETE("/:id", memberController.DeleteMember)
	}
	member = server.Group("/public")
	{
		member.GET("class/members/:classID", memberController.GetAllMembersByClassIDData)
		member.GET("/user/class/:userID", memberController.GetAllClassByUserID)
		member.GET("/assessment/upcoming/:userID", memberController.GetAllClassAndAssesmentByUserID)
	}
}