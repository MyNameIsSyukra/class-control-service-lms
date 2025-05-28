package controller

import (
	"LMSGo/dto"
	kelas "LMSGo/service"
	response "LMSGo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	KelasController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
	kelasController struct {
		kelasService kelas.KelasService 
	}
)

func NewKelasController(kelasService kelas.KelasService) KelasController {
	return &kelasController{kelasService}
}

func (service *kelasController) Create(ctx *gin.Context) {
	var req dto.CreateKelasRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.FailedResponse("Invalid request")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	create, err := service.kelasService.Create(ctx.Request.Context(), &req)
	if err != nil {
		res := response.FailedResponse("Failed to create class")
		ctx.JSON(500, res)
		return
	}
	res := response.SuccessResponse(create)	
	ctx.JSON(200, res)
}

func (service *kelasController) GetAll(ctx *gin.Context) {
	var pagination dto.PaginationRequest
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		res := response.FailedResponse("Invalid body parameters")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	kelas, err := service.kelasService.GetAllKelasWithPagination(ctx.Request.Context(),pagination)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get classes"})
		return
	}
	res := response.SuccessResponse(kelas)
	ctx.JSON(200, res)
}

func (service *kelasController) GetById(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := response.FailedResponse("Invalid ID")
		ctx.JSON(400, res)
		return
	}
	kelas, err := service.kelasService.GetById(ctx.Request.Context(), id)
	if err != nil {
		res := response.FailedResponse("Failed to get class")
		ctx.JSON(404, res)
		return
	}
	if kelas.ID == uuid.Nil {
		res := response.FailedResponse("Class not found")
		ctx.JSON(404, res)
		return
	}
	res := response.SuccessResponse(kelas)
	ctx.JSON(200, res)
}

func (service *kelasController) Update(ctx *gin.Context) {
	id,err :=  uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := response.FailedResponse("Invalid ID")
		ctx.JSON(400, res)
		return
	}
	var req dto.KelasUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.FailedResponse("Invalid request body")
		ctx.JSON(400, res)
		return
	}

	kelas, err := service.kelasService.Update(ctx.Request.Context(), id, &req)
	if err != nil {
		res := response.FailedResponse("Failed to update class")
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, kelas)
}

func (service *kelasController) Delete(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := response.FailedResponse("Invalid ID")
		ctx.JSON(400, res)
		return
	}
	err = service.kelasService.Delete(ctx.Request.Context(), id)
	if err != nil {
		res := response.FailedResponse("Failed to delete class")
		ctx.JSON(500, res)
		return
	}
	res := response.SuccessResponse(nil)
	ctx.JSON(200, res)
}
