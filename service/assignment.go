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
		CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest,file io.Reader) (*dto.AssignmentResponse, error)
		GetAssignmentByID(ctx context.Context, assignmentID int) (*dto.AssignmentResponse, error)
		UpdateAssignment(ctx context.Context, request dto.ProrcessedUpdateAssignmentRequest, file io.Reader) (*dto.AssignmentResponse,error)
		DeleteAssignment(ctx context.Context, assignmentID int) error

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
		Id string `json:"id"`
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
	
	fmt.Printf("File uploaded successfully: %s\n", uploadResp.Id)
	return uploadResp.Id, nil
}

func (service *assignmentService) CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest, file io.Reader) (*dto.AssignmentResponse, error) {
	fmt.Printf("Starting assignment creation\n")
	
	if file != nil {
		fileId, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return nil, err
		}
		request.FileId = fileId
	} else {
		fmt.Printf("No file provided\n")
		request.FileId = ""
		request.FileName = ""
	}
	
	fmt.Printf("Creating assignment in database\n")
	
	// Create assignment in database
	newAssignment, err := service.assignmentRepo.CreateAssignment(ctx, nil, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create assignment: %w", err)
	}
	fileUrl := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/?id=" + newAssignment.FileId
	if newAssignment.FileId == "" {
		fileUrl = ""
	}
	fmt.Printf("Assignment created successfully\n")
	return &dto.AssignmentResponse{
		AssignmentID:      int(newAssignment.ID),
		Title:       newAssignment.Title,
		Description: newAssignment.Description,
		Deadline:    newAssignment.Deadline,
		FileName:    &newAssignment.FileName,
		FileId:      &newAssignment.FileId,
		FileUrl:     &fileUrl,
	}, nil
}

func (service *assignmentService) UpdateAssignment(ctx context.Context, request dto.ProrcessedUpdateAssignmentRequest, file io.Reader) (*dto.AssignmentResponse,error) {
	fmt.Printf("Starting assignment update\n")
	// check if assignment exists
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, request.AssignmentID)
	if err != nil {
		return nil,fmt.Errorf("assignment not found: %w", err)
	}
	if file != nil {
		fileId := assignment.FileId
		if fileId != "" {
			params := url.Values{}
			params.Add("id", fileId)
			delurl := os.Getenv("CONTENT_URL") + "/item-pembelajaran/?" + params.Encode()
			delreq, err := http.NewRequest(http.MethodDelete, delurl, nil)
			if err != nil {
				return nil,fmt.Errorf("failed to create delete request: %w", err)
			}
			client := &http.Client{}
			resp, err := client.Do(delreq)
			if err != nil {
				return nil,fmt.Errorf("failed to delete old file: %w", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusNoContent {
				respBody, _ := io.ReadAll(resp.Body)
				return nil,fmt.Errorf("failed to delete old file with status %d: %s", resp.StatusCode, string(respBody))
			}
		}
		fileId, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return nil,fmt.Errorf("failed to upload new file: %w", err)
		}
		request.FileId = fileId 
	} else {
		fmt.Printf("No file provided for update\n")
		request.FileId = assignment.FileId
		request.FileName = assignment.FileName
	}
	var updates dto.ProrcessedUpdateAssignmentRequest
	updates.AssignmentID = request.AssignmentID
	updates.WeekID = request.WeekID
	if request.Title != "" {
		updates.Title = request.Title
	} else {
		updates.Title = assignment.Title
	}
	if request.Description != "" {
		updates.Description = request.Description
	} else {
		updates.Description = assignment.Description
	}
	if !request.Deadline.IsZero() {
		updates.Deadline = request.Deadline
	} else {
		updates.Deadline = assignment.Deadline
	}
	if request.FileName != "" {
		updates.FileName = request.FileName
	} else {
		updates.FileName = assignment.FileName
	}
	if request.FileId != "" {
		updates.FileId = request.FileId
	} else {
		updates.FileId = assignment.FileId
	}

	// Update assignment in database
	updated, err := service.assignmentRepo.UpdateAssignment(ctx, nil, request.AssignmentID, updates)
	if err != nil {
		return nil,fmt.Errorf("failed to update assignment: %w", err)
	}
	fmt.Printf("Assignment updated successfully\n")
	fileUrl := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/?id=" + updated.FileId
	return &dto.AssignmentResponse{
		AssignmentID:      int(updated.ID),
		Title:       updated.Title,
		Description: updated.Description,
		Deadline:    updated.Deadline,
		FileName:    &updated.FileName,
		FileId:      &updated.FileId,
		FileUrl:     &fileUrl,
	}, nil
	
}

func (service *assignmentService) GetAssignmentByID(ctx context.Context, assignmentID int) (*dto.AssignmentResponse, error) {
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return nil, err
	}
	fileUrl := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/?id=" + assignment.FileId
	return &dto.AssignmentResponse{
		AssignmentID:      int(assignment.ID),
		Title:       assignment.Title,
		Description: assignment.Description,
		Deadline:    assignment.Deadline,
		FileName:    &assignment.FileName,
		FileId:      &assignment.FileId,
		FileUrl:     &fileUrl,
	}, nil
}


func (service *assignmentService) DeleteAssignment(ctx context.Context, assignmentID int) error {
	// check if assignment exists
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return fmt.Errorf("assignment not found: %w", err)
	}
	// delete file if exists
	if assignment.FileId != "" {
		params := url.Values{}
		params.Add("id", assignment.FileId)
		delurl := os.Getenv("CONTENT_URL") + "/item-pembelajaran/?" + params.Encode()
		delreq, err := http.NewRequest(http.MethodDelete, delurl, nil)
		if err != nil {
			return fmt.Errorf("failed to create delete request: %w", err)
		}
		client := &http.Client{}
		resp, err := client.Do(delreq)
		if err != nil {
			return fmt.Errorf("failed to delete file: %w", err)
		}
		defer resp.Body.Close()
	}
	err = service.assignmentRepo.DeleteAssignment(ctx, nil, assignmentID)
	if err != nil {
		return fmt.Errorf("failed to delete assignment: %w", err)
	}
	return nil
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
	fileLink := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/?id=" + assignment.FileId
	resp.ID = int(assignment.ID)
	resp.Title = assignment.Title
	resp.Description = assignment.Description
	resp.Deadline = assignment.Deadline
	resp.FileName = assignment.FileName
	resp.FileLink = fileLink
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
		link := os.Getenv("GATEWAY_URL") + "/student-assignment/user?" + params.Encode()
		resp.StudentSubmissionLink = &link
		resp.StudentSubmissionFileName = &assSubmission.FileName
	} else {
		resp.Score = 0
		resp.Status = entities.StatusTodo
	}
	return resp, nil
}