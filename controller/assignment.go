package controller

import (
	"LMSGo/dto"
	Assignment "LMSGo/service"
	"LMSGo/utils"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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
		// Validasi tipe file berdasarkan ekstensi
		allowedExtensions := map[string]bool{
			".doc":  true,
			".docx": true,
			".pdf":  true,
			".ppt":  true,
			".pptx": true,
			".xls":  true,
			".xlsx": true,
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}

		// Ambil ekstensi file
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !allowedExtensions[ext] {
			res := utils.FailedResponse("File type not allowed. Only .doc, .docx, .pdf, .ppt, .pptx, .xls, .xlsx, .jpg, .jpeg, .png files are permitted")
			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		// Validasi tambahan berdasarkan MIME type
		openedFile, err := fileHeader.Open()
		if err != nil {
			res := utils.FailedResponse("unable to open file")
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		defer openedFile.Close()

		// Baca 512 bytes pertama untuk deteksi MIME type
		buffer := make([]byte, 512)
		_, err = openedFile.Read(buffer)
		if err != nil {
			res := utils.FailedResponse("unable to read file content")
			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		// Reset file pointer ke awal
		openedFile.Seek(0, 0)

		// Deteksi MIME type
		mimeType := http.DetectContentType(buffer)
		
		allowedMimeTypes := map[string]bool{
			"application/pdf":                                                       true,
			"application/msword":                                                   true,
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
			"application/vnd.ms-powerpoint":                                        true,
			"application/vnd.openxmlformats-officedocument.presentationml.presentation": true,
			"application/vnd.ms-excel":                                             true,
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":   true,
			"image/jpeg":                                                           true,
			"image/jpg":                                                            true,
			"image/png":                                                            true,
			"application/octet-stream":                                             true, // Untuk beberapa file Office yang mungkin terdeteksi sebagai binary
		}

		// Khusus untuk file Office yang mungkin terdeteksi sebagai octet-stream
		if mimeType == "application/octet-stream" {
			// Periksa apakah ekstensi file adalah Office document
			officeExtensions := map[string]bool{
				".doc":  true,
				".docx": true,
				".ppt":  true,
				".pptx": true,
				".xls":  true,
				".xlsx": true,
			}
			if !officeExtensions[ext] {
				res := utils.FailedResponse("Invalid file type detected")
				ctx.JSON(http.StatusBadRequest, res)
				return
			}
		} else if !allowedMimeTypes[mimeType] {
			res := utils.FailedResponse(fmt.Sprintf("File MIME type not allowed: %s", mimeType))
			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		file = openedFile
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

	// Coba ambil file dari form, jika ada
	fileHeader, err := ctx.FormFile("file")
	var file io.Reader
	var fileName string
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
	claims, err := DecodeJWTToken(ctx)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(400, res)
		return
	}
	userID := claims.UserID
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