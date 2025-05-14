package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	Kelas(server, injector)
	Member(server, injector)	
	Service(server, injector)
	WeekSection(server, injector)	
	Assignment(server, injector)
	AssignmentSubmission(server, injector)
}