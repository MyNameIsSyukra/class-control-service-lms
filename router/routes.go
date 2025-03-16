package routes

import (
	"LMSGo/controller"
	entities "LMSGo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
) 

func SetupRouter(classUC controller.KelasController) *gin.Engine {
	router := gin.Default()

	router.GET("/class", func(c *gin.Context) {
		class, err := classUC.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, class)
	})

	router.GET("/class/:id", func(c *gin.Context) {
		id := c.Param("id")
		class, err := classUC.GetById(id)
		if err != nil {
			
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, class)
	})

	router.POST("/class", func(c *gin.Context) {
		var class entities.Kelas
		c.BindJSON(&class)
		err := classUC.Create(&class)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, class)
	})

	return router
}
