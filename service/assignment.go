package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	"LMSGo/repository"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	AssignmentService interface {
		CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest,file io.Reader) (*entities.Assignment, error)
		GetAssignmentByID(ctx context.Context, assignmentID int) (entities.Assignment, error)

		// student
		GetAssignmentByIDStudentID(ctx context.Context, assignmentID int, userID uuid.UUID) (dto.StudentGetAssignmentByIDResponse, error)
	}
	assignmentService struct {
		assignmentRepo repository.AssignmentRepository
		assignmentSubmissionRepo repository.AssignmentSubmissionRepository
	}
)

func NewAssignmentService(assignmentRepo repository.AssignmentRepository, assignmentSubmissionRepo repository.AssignmentSubmissionRepository) AssignmentService {
	return &assignmentService{assignmentRepo,assignmentSubmissionRepo}
}

func (service *assignmentService) uploadFile(file io.Reader, fileName string) (string, error) {
	type FileUploadResponse struct {
		Url string `json:"url"`
	}
	
	fmt.Printf("Processing file upload\n")
	
	// Create multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	
	// Create form file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	
	// Copy file content to form
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}
	
	// Close writer to finalize multipart data
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %w", err)
	}
	
	// Prepare HTTP request
	url := os.Getenv("CONTENT_URL") + "/item-pembelajaran/"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	
	// Set proper content type with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	fmt.Printf("Sending file upload request\n")
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	
	fmt.Printf("Received response with status: %d\n", resp.StatusCode)
	
	// Check response status
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("file upload failed with status %d: %s", resp.StatusCode, string(respBody))
	}
	
	// Parse response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	
	var uploadResp FileUploadResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return "", fmt.Errorf("failed to parse upload response: %w", err)
	}
	
	fmt.Printf("File uploaded successfully: %s\n", uploadResp.Url)
	return uploadResp.Url, nil
}

func (service *assignmentService) CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest, file io.Reader) (*entities.Assignment, error) {
	fmt.Printf("Starting assignment creation\n")
	
	if file != nil {
		fileURL, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return nil, err
		}
		request.FileLink = fileURL
	} else {
		fmt.Printf("No file provided\n")
		request.FileLink = ""
		request.FileName = ""
	}
	
	fmt.Printf("Creating assignment in database\n")
	
	// Create assignment in database
	newAssignment, err := service.assignmentRepo.CreateAssignment(ctx, nil, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create assignment: %w", err)
	}
	
	fmt.Printf("Assignment created successfully\n")
	return newAssignment, nil
}

func (service *assignmentService) GetAssignmentByID(ctx context.Context, assignmentID int) (entities.Assignment, error) {
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return entities.Assignment{}, err
	}
	return assignment, nil
}

// student
func (service *assignmentService) GetAssignmentByIDStudentID(ctx context.Context, assignmentID int, userID uuid.UUID) (dto.StudentGetAssignmentByIDResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return dto.StudentGetAssignmentByIDResponse{}, err
	}
	
	// check if student has submitted the assignment
	assSubmission, err := service.assignmentSubmissionRepo.CheckStudentSubmssionByAssIdUserID(ctx, nil, assignmentID, userID)
	if err != nil {
		return dto.StudentGetAssignmentByIDResponse{}, err
	}
	link := os.Getenv("CONTENT_URL") + "/student-assignment/user/" + assSubmission.IDFile
	resp := dto.StudentGetAssignmentByIDResponse{
		WeekID:      assignment.WeekID,
		Title:       assignment.Title,
		Description: assignment.Description,
		FileName:    assignment.FileName,
		FileLink:    assignment.FileLink,
		StudentSubmissionLink: link,
		Status:      assSubmission.Status,
		Score:       assSubmission.Score,
	}
	return resp, nil
}