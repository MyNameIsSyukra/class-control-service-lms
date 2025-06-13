package routes

import (
	"LMSGo/controller"
	"LMSGo/middleware"
	"LMSGo/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)


func Member(server *gin.Engine, injector *do.Injector) {
	memberController := do.MustInvoke[controller.MemberController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	member := server.Group("/member/admin")
	{
		member.POST("",middleware.Authenticate(jwtService) , memberController.AddMemberToClass)
		member.DELETE("",middleware.Authenticate(jwtService) , memberController.DeleteMember)
	}
	member = server.Group("/public")
	{
		member.GET("class/members",middleware.Authenticate(jwtService) , memberController.GetAllMembersByClassIDData)
		member.GET("/user/class",middleware.Authenticate(jwtService) , memberController.GetAllClassByUserID)
		member.GET("/assessment/upcoming", middleware.Authenticate(jwtService) ,memberController.GetAllClassAndAssesmentByUserID)
	}
}