package routes

import (
	"LMSGo/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)


func Service(server *gin.Engine, injector *do.Injector) {
	memberController := do.MustInvoke[controller.MemberController](injector)
	// kelasController := do.MustInvoke[controller.KelasController](injector)	

	service := server.Group("/service")
	{
		service.GET("/class/:classID", memberController.GetAllMembersByClassID)
	}	
}