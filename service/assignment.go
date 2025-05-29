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
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	AssignmentService interface {
		CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest,file io.Reader) (*entities.Assignment, error)
		GetAssignmentByID(ctx context.Context, assignmentID int) (entities.Assignment, error)
		UpdateAssignment(ctx context.Context, request dto.ProrcessedUpdateAssignmentRequest, file io.Reader) error

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

func (service *assignmentService) UpdateAssignment(ctx context.Context, request dto.ProrcessedUpdateAssignmentRequest, file io.Reader) error {
	fmt.Printf("Starting assignment update\n")
	// check if assignment exists
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, request.AssignmentID)
	if err != nil {
		return fmt.Errorf("assignment not found: %w", err)
	}
	if file != nil {
		delurl := assignment.FileLink
		if delurl != "" {
			delreq, err := http.NewRequest(http.MethodDelete, delurl, nil)
			if err != nil {
				return fmt.Errorf("failed to create delete request: %w", err)
			}
			client := &http.Client{}
			resp, err := client.Do(delreq)
			if err != nil {
				return fmt.Errorf("failed to delete old file: %w", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusNoContent {
				respBody, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("failed to delete old file with status %d: %s", resp.StatusCode, string(respBody))
			}
		}
		fileURL, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return fmt.Errorf("failed to upload new file: %w", err)
		}
		request.FileLink = fileURL
	} else {
		fmt.Printf("No file provided for update\n")
		request.FileLink = assignment.FileLink
		request.FileName = assignment.FileName
	}
	// Update assignment in database
	_, err = service.assignmentRepo.UpdateAssignment(ctx, nil, request.AssignmentID, request)
	if err != nil {
		return fmt.Errorf("failed to update assignment: %w", err)
	}
	fmt.Printf("Assignment updated successfully\n")
	return nil
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
	var resp dto.StudentGetAssignmentByIDResponse
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return dto.StudentGetAssignmentByIDResponse{}, err
	}
	resp.WeekID = assignment.WeekID
	resp.Title = assignment.Title
	resp.Description = assignment.Description
	resp.FileName = assignment.FileName
	resp.FileLink = assignment.FileLink
	// check if student has submitted the assignment
	assSubmission, err := service.assignmentSubmissionRepo.CheckStudentSubmssionByAssIdUserID(ctx, nil, assignmentID, userID)
	if err != nil {
		return dto.StudentGetAssignmentByIDResponse{}, err
	}
	if assSubmission.ID == uuid.Nil {
		resp.Score = 0
		resp.Status = entities.StatusTodo
		resp.StudentSubmissionLink = nil
	}
	if assSubmission.Status == entities.StatusSubmitted || assSubmission.Status == entities.StatusLate {
		resp.Score = assSubmission.Score
		resp.Status = assSubmission.Status
		params := url.Values{}
		params.Add("id", assSubmission.IDFile)
		link := os.Getenv("CONTENT_URL") + "/student-assignment/user/?" + params.Encode()
		resp.StudentSubmissionLink = &link
	} else {
		resp.Score = 0
		resp.Status = entities.StatusTodo
	}
	return resp, nil
}