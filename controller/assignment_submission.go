package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
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
	AssignmentSubmissionController interface {
		CreateAssignmentSubmission(ctx *gin.Context)
		GetAllStudentAssignmentSubmissionByAssignmentID(ctx *gin.Context)
		GetAssignmentSubmissionByID(ctx *gin.Context)
		UpdateStudentSubmissionScore(ctx *gin.Context)
		DeleteAssignmentSubmissionByID(ctx *gin.Context)
	}
	assignmentSubmissionController struct {
		assignmentSubmissionService Assignment.AssignmentSubmissionService
	}
)


func NewAssignmentSubmissionController(assignmentSubmissionService Assignment.AssignmentSubmissionService) AssignmentSubmissionController {
	return &assignmentSubmissionController{assignmentSubmissionService}
}

func (controller *assignmentSubmissionController) CreateAssignmentSubmission(ctx *gin.Context) {
	var req dto.InitAssignmentSubmissionRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	claims, err := DecodeJWTToken(ctx)
	if err != nil {
		res := utils.FailedResponse(fmt.Sprintf("Authentication failed: %s", err.Error()))
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	cleanUUIDStr := strings.Trim(claims.UserID, "[]\"")
	println("Cleaned UUID String:", cleanUUIDStr)
	userID, err := uuid.Parse(cleanUUIDStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid UUID format",
            "details": fmt.Sprintf("Cannot parse UUID: %s", cleanUUIDStr),
            "received": userID,
        })
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
		// Validasi tipe file berdasarkan ekstensi
		validatedFile := utils.ValidateFileUpload(ctx, fileHeader)
		if validatedFile == nil {
			// Jika file tidak valid, sudah ditangani di ValidateFileUpload
			return
		}		
		file = validatedFile
		fileName = fileHeader.Filename
	} else {
		// File tidak ada, set nil (opsional)
		file = nil
		fileName = ""
	}

	processedReq := dto.AssignmentSubmissionRequest{
		AssignmentID: req.AssignmentID,
		UserID:       userID,
		FileName:     fileName,
	}

	assignmentSubmission, err := controller.assignmentSubmissionService.CreateAssignmentSubmission(ctx.Request.Context(), processedReq, file)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assignmentSubmission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) GetAllStudentAssignmentSubmissionByAssignmentID(ctx *gin.Context) {
	assignmentID := ctx.Query("assignment_id")
	status := ctx.Query("status")
	parsedAssignmentID, err := strconv.Atoi(assignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submissions, err := controller.assignmentSubmissionService.GetAllStudentAssignmentSubmissionByAssignmentID(ctx.Request.Context(),status, parsedAssignmentID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submissions)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) UpdateStudentSubmissionScore(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	score := ctx.Query("score")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	parsedScore, err := strconv.Atoi(score)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submission, err := controller.assignmentSubmissionService.UpdateStudentSubmissionScore(ctx.Request.Context(), parsedScore, parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) GetAssignmentSubmissionByID(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	submission, err := controller.assignmentSubmissionService.GetAssignmentSubmissionByID(ctx.Request.Context(), parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submission)
	ctx.JSON(200, res)
}

func (controller *assignmentSubmissionController) DeleteAssignmentSubmissionByID(ctx *gin.Context) {
	assignmentSubmissionID := ctx.Query("assignment_submission_id")
	parsedAssignmentSubmissionID, err := uuid.Parse(assignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	err = controller.assignmentSubmissionService.DeleteAssignmentSubmissionByID(ctx.Request.Context(), parsedAssignmentSubmissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(nil)
	ctx.JSON(200, res)
}