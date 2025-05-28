package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AssignmentController interface {
		CreateAssignment(ctx *gin.Context)
		GetAssignmentByID(ctx *gin.Context)

		// student
		GetAssignmentByIDStudentID(ctx *gin.Context)
	}
	assignmentController struct {
		assignmentService Assignment.AssignmentService
	}
)

func NewAssignmentController(assignmentService Assignment.AssignmentService) AssignmentController {
	return &assignmentController{assignmentService}
}

func (controller *assignmentController) CreateAssignment(ctx *gin.Context) {
	var req dto.CreateAssignmentRequest
	// Bind form fields (tanpa file)
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
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
	req.FileName = fileName
	fmt.Println("File Name:", req.FileName)
	assignment, err := controller.assignmentService.CreateAssignment(ctx.Request.Context(), req, file)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.SuccessResponse(assignment)
	ctx.JSON(http.StatusOK, res)
}


func (controller *assignmentController) GetAssignmentByID(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	assignment, err := controller.assignmentService.GetAssignmentByID(ctx.Request.Context(), parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(200, res)
}

// student

func (controller *assignmentController) GetAssignmentByIDStudentID(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	userID := ctx.Query("user_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	assignment, err := controller.assignmentService.GetAssignmentByIDStudentID(ctx.Request.Context(), parsedAssignmentID, parsedUserID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(200, res)
}