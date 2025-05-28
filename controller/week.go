package controller

import (
	"LMSGo/dto"
	"LMSGo/service"
	"LMSGo/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	WeekController interface {
		CreateWeeklySection(ctx *gin.Context)
		UpdateWeeklySection(ctx *gin.Context)
		GetAllWeekByClassID(ctx *gin.Context)
		GetWeekByID(ctx *gin.Context)
		DeleteWeeklySection(ctx *gin.Context)
	}
	weekController struct {
		weekService service.WeekService
	}
)

func NewWeekController(weekService service.WeekService) WeekController {
	return &weekController{weekService}
}

type FormUUID struct {
    uuid.UUID
}
func (u *FormUUID) UnmarshalText(text []byte) error {
    str := string(text)
    // Remove brackets dan quotes jika ada
    str = strings.Trim(str, "[]\"")
    parsed, err := uuid.Parse(str)
    if err != nil {
        return err
    }
    u.UUID = parsed
    return nil
}

func (controller *weekController) CreateWeeklySection(ctx *gin.Context) {
	var req dto.ItemPembelajaranRequest
	// Bind form fields (tanpa file)
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	cleanUUIDStr := strings.Trim(req.KelasIDStr, "[]\"")
	kelasID, err := uuid.Parse(cleanUUIDStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid UUID format",
            "details": fmt.Sprintf("Cannot parse UUID: %s", cleanUUIDStr),
            "received": req.KelasIDStr,
        })
        return
    }
	processedReq := dto.CreateItemPembelajaranRequest{
        KelasID:          kelasID,
        WeekNumber:       req.WeekNumber,
        HeadingPertemuan: req.HeadingPertemuan,
        BodyPertemuan:    req.BodyPertemuan,
        UrlVideo:         req.UrlVideo,
    }

	var file io.Reader
	var fileName string
    fileCount := 0
    if ctx.Request.MultipartForm != nil && ctx.Request.MultipartForm.File != nil {
        for _, files := range ctx.Request.MultipartForm.File {
            fileCount += len(files)
        }
    }
    
    // LIMIT: Hanya boleh 1 file
    if fileCount > 1 {
		res := utils.FailedResponse("Only one file upload is allowed")
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
	// Coba ambil file dari form, jika ada
	fileHeader, err := ctx.FormFile("file")
	if err == nil {
		openedFile, err := fileHeader.Open()
		if err != nil {
			res := utils.FailedResponse("unable to open file")
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		defer openedFile.Close()
		file = openedFile
		fileName = fileHeader.Filename
	} else {
		// File tidak ada, set nil (opsional)
		file = nil
		fileName = ""
	}
	processedReq.FileName = fileName
	item, err := controller.weekService.CreateWeeklySection(ctx.Request.Context(),processedReq,file)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(item)
	ctx.JSON(200, res)
}

func (controller *weekController) UpdateWeeklySection(ctx *gin.Context) {
	KelasID := ctx.Query("kelas_id")
	parsedKelasID, err := uuid.Parse(KelasID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	WeekNumber := ctx.Query("week_number")
	parsedWeekNumber, err := strconv.Atoi(WeekNumber)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	var req dto.UpdateItemPembelajaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	item, err := controller.weekService.UpdateWeeklySection(ctx.Request.Context(),parsedKelasID,parsedWeekNumber, req)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(item)
	ctx.JSON(200, res)
}

func (controller *weekController) GetAllWeekByClassID(ctx *gin.Context) {
	classID := ctx.Query("class_id")
	parsedClassID, err := uuid.Parse(classID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	weeks, err := controller.weekService.GetAllWeekByClassID(ctx.Request.Context(), parsedClassID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(weeks)
	ctx.JSON(200, res)
}

func (controller *weekController) GetWeekByID(ctx *gin.Context) {
	weekID := ctx.Query("id")
	weekIDInt, err := strconv.Atoi(weekID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	week, err := controller.weekService.GetWeekByID(ctx.Request.Context(), weekIDInt)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}

	res := utils.SuccessResponse(week)
	ctx.JSON(200, res)
}

func (controller *weekController) DeleteWeeklySection(ctx *gin.Context) {
	weekID := ctx.Query("id")
	parsedWeekID, err := strconv.Atoi(weekID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	err = controller.weekService.DeleteWeeklySection(ctx.Request.Context(), parsedWeekID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(500, res)
		return
	}
	res := utils.SuccessResponse(nil)
	ctx.JSON(200, res)	
}


