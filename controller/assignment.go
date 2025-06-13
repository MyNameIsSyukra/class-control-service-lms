package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AssignmentController interface {
		CreateAssignment(ctx *gin.Context)
		GetAssignmentByID(ctx *gin.Context)
		UpdateAssignment(ctx *gin.Context)
		DeleteAssignment(ctx *gin.Context)

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
	var req dto.AssignmentRequest
	// Bind form fields (tanpa file)
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// Form Validation
	if req.Deadline.Before(time.Now()){
		res := utils.FailedResponse("Deadline cannot be in the past")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	processedReq := dto.CreateAssignmentRequest{
		WeekID:      req.WeekID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
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
		// Validasi ekstensi file
		validatedFile := utils.ValidateFileUpload(ctx, fileHeader)
		if validatedFile == nil {
			// Jika validasi gagal, sudah ditangani di ValidateFileUpload
			return
		}
		
		file = validatedFile
		fileName = fileHeader.Filename
	} else {
		// File tidak ada, set nil (opsional)
		file = nil
		fileName = ""
	}

	processedReq.FileName = fileName
	fmt.Println("File Name:", processedReq.FileName)
	
	assignment, err := controller.assignmentService.CreateAssignment(ctx.Request.Context(), processedReq, file)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignment)
	ctx.JSON(http.StatusOK, res)
}

func (controller *assignmentController) UpdateAssignment(ctx *gin.Context) {
	var req dto.InitUpdateAssignmentRequest
	// Bind form fields (tanpa file)
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// form validation
	if req.Deadline.Before(time.Now()) {
		res := utils.FailedResponse("Deadline cannot be in the past")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	// Coba ambil file dari form, jika ada
	fileHeader, err := ctx.FormFile("file")
	var file io.Reader
	var fileName string
	if err == nil {
		// Validasi ekstensi file
		validatedFile := utils.ValidateFileUpload(ctx, fileHeader)
		if validatedFile == nil {
			// Jika validasi gagal, sudah ditangani di ValidateFileUpload
			return
		}
		file = validatedFile
		fileName = fileHeader.Filename
	} else {
		file = nil // File tidak ada, set nil (opsional)
		fileName = ""
	}

	processedReq := dto.ProrcessedUpdateAssignmentRequest{
		AssignmentID: req.AssignmentID,
		WeekID:       req.WeekID,
		Title:        req.Title,
		Description:  req.Description,
		Deadline:     req.Deadline,
		FileName:     fileName,
	}
	updated,err := controller.assignmentService.UpdateAssignment(ctx.Request.Context(), processedReq,file)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(updated)
	ctx.JSON(http.StatusOK, res)
}

func (controller *assignmentController) DeleteAssignment(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	err = controller.assignmentService.DeleteAssignment(ctx.Request.Context(), parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse("Assignment deleted successfully")
	ctx.JSON(200, res)
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
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}

	userID :=ctx.MustGet("uuid").(string)

	// userID := claims.UserID
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