package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateFileUpload(ctx *gin.Context, fileHeader *multipart.FileHeader)multipart.File {
	const maxFileSize = 5 * 1024 * 1024 // 5MB dalam bytes
	if fileHeader.Size > maxFileSize {
		res := FailedResponse("File size exceeds limit. Maximum file size is 5MB")
		ctx.JSON(http.StatusBadRequest, res)
		return nil
	}
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
		res := FailedResponse("File type not allowed. Only .doc, .docx, .pdf, .ppt, .pptx, .xls, .xlsx, .jpg, .jpeg, .png files are permitted")
		ctx.JSON(http.StatusBadRequest, res)
		return nil
	}
	// Validasi tambahan berdasarkan MIME type
	openedFile, err := fileHeader.Open()
	if err != nil {
		res := FailedResponse("unable to open file")
		ctx.JSON(http.StatusBadRequest, res)
		return nil
	}
	defer openedFile.Close()
	// Baca 512 bytes pertama untuk deteksi MIME type
	buffer := make([]byte, 512)
	_, err = openedFile.Read(buffer)
	if err != nil {
		res := FailedResponse("unable to read file content")
		ctx.JSON(http.StatusBadRequest, res)
		return nil
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
			res := FailedResponse("Invalid file type detected")
			ctx.JSON(http.StatusBadRequest, res)
			return nil
		}
	} else if !allowedMimeTypes[mimeType] {
		res := FailedResponse(fmt.Sprintf("File MIME type not allowed: %s", mimeType))
		ctx.JSON(http.StatusBadRequest, res)
		return nil
	}
	return openedFile
}