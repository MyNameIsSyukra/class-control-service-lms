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
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	AssignmentSubmissionService interface {
		CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest, file io.Reader) (*entities.AssignmentSubmission, error)

		// teacher
		GetAllStudentAssignmentSubmissionByAssignmentID(ctx context.Context,status string, assignmentID int) ([]dto.GetAssSubmissionStudentResponse, error)
		UpdateStudentSubmissionScore(ctx context.Context, score int, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error)
		GetAssignmentSubmissionByID(ctx context.Context, assignmentSubmissionID uuid.UUID) (dto.GetAssSubmissionStudentResponse, error)
		DeleteAssignmentSubmissionByID(ctx context.Context, assignmentSubmissionID uuid.UUID) error
	}
	assignmentSubmissionService struct {
		assignmentSubmissionRepo repository.AssignmentSubmissionRepository
		memberRepo 			repository.StudentRepository
		assignmentRepo 		repository.AssignmentRepository
	}
)

func NewAssignmentSubmissionService(assignmentSubmissionRepo repository.AssignmentSubmissionRepository, memberRepo repository.StudentRepository, assigmentRepo repository.AssignmentRepository) AssignmentSubmissionService {
	return &assignmentSubmissionService{assignmentSubmissionRepo, memberRepo, assigmentRepo}
}
func (service *assignmentSubmissionService) CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest, file io.Reader) (*entities.AssignmentSubmission, error) {
	// check if the user has submitted the assignment
	assSubmission, err := service.assignmentSubmissionRepo.CheckStudentSubmssionByAssIdUserID(ctx, nil, request.AssignmentID, request.UserID)
	if err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	if assSubmission.Status == entities.StatusSubmitted || assSubmission.Status == entities.StatusLate {
		return &entities.AssignmentSubmission{}, fmt.Errorf("you have submitted this assignment")
	}
	// check if student is member of the class
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx,nil,request.AssignmentID)
	if err != nil{
		return &entities.AssignmentSubmission{}, err
	}
	if assignment.Deadline.Before(time.Now()) {
		return &entities.AssignmentSubmission{}, fmt.Errorf("assignment deadline has passed")
	}
	
	member, err := service.memberRepo.GetMemberByClassIDAndUserID(ctx,nil,assignment.Week.Kelas_idKelas, request.UserID)
	if err != nil {
		return nil, fmt.Errorf("you are not class member")
	}
	if member.User_userID == uuid.Nil {
		return nil, fmt.Errorf("you are not class member")
	}
	if member.Role == entities.MemberRoleTeacher{
		return nil, fmt.Errorf("you are a teacher")
	}

	// check if the file is empty
	if file == nil {
		return &entities.AssignmentSubmission{}, fmt.Errorf("file is required")
	}
	IDFile, err := service.uploadFile(file, request.FileName, request.UserID.String())
	if err != nil {
		return &entities.AssignmentSubmission{}, fmt.Errorf("failed to upload file: %w", err)
	}
	request.IDFile = IDFile
	newAssignmentSubmission, err := service.assignmentSubmissionRepo.CreateAssignmentSubmission(ctx, nil, request)
	if err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	return newAssignmentSubmission, nil
}


// teacher
func (service *assignmentSubmissionService) UpdateStudentSubmissionScore(ctx context.Context, score int, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error) {
	assignmentSubmission, err := service.assignmentSubmissionRepo.UpdateStudentSubmissionScore(ctx, nil, score, assignmentSubmissionID)
	if err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	return assignmentSubmission, nil
}

func (service *assignmentSubmissionService) GetAssignmentSubmissionByID(ctx context.Context, assignmentSubmissionID uuid.UUID) (dto.GetAssSubmissionStudentResponse, error) {
	assignmentSubmission, err := service.assignmentSubmissionRepo.GetAssignmentSubmissionByID(ctx, nil, assignmentSubmissionID)
	if err != nil {
		return dto.GetAssSubmissionStudentResponse{}, err
	}
	if assignmentSubmission.ID == uuid.Nil {
		return dto.GetAssSubmissionStudentResponse{}, fmt.Errorf("assignment submission not found")
	}
	student, err := service.memberRepo.GetMemberById(ctx, nil, assignmentSubmission.UserID)
	if err != nil {
		return dto.GetAssSubmissionStudentResponse{}, err
	}
	params := url.Values{}
	params.Add("id", assignmentSubmission.IDFile)
	link := os.Getenv("GATEWAY_URL") + "/teacher/student-assignment/" + "?" + params.Encode()	
	photoUrl := os.Getenv("GATEWAY_URL") + "/storage/user_profile_pictures" + student.User_userID.String() + ".jpg"
	return dto.GetAssSubmissionStudentResponse{
		ID:         &assignmentSubmission.ID,
		User_userID: student.User_userID,
		Username:   student.Username,
		PhotoUrl:   &photoUrl,
		Status:     assignmentSubmission.Status,
		Score:      assignmentSubmission.Score,
		LinkFile:   &link,
		CreatedAt: &assignmentSubmission.CreatedAt,
		UpdatedAt: &assignmentSubmission.UpdatedAt,
	}, nil
}


func (service *assignmentSubmissionService) GetAllStudentAssignmentSubmissionByAssignmentID(ctx context.Context,status string, assignmentID int) ([]dto.GetAssSubmissionStudentResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return []dto.GetAssSubmissionStudentResponse{}, err
	}

	members, err := service.memberRepo.GetAllMembersByClassID(ctx, nil, assignment.Week.Kelas_idKelas)
	if err != nil {
		return []dto.GetAssSubmissionStudentResponse{}, err
	}

	assignmentSubmissions, err := service.assignmentSubmissionRepo.GetAllSubmissionByAssignmentID(ctx, nil, assignmentID)
	if err != nil {
		return []dto.GetAssSubmissionStudentResponse{}, err
	}

	submissionMap := make(map[uuid.UUID]*entities.AssignmentSubmission)
	submissionMapping := make(map[uuid.UUID]*entities.AssignmentSubmission)
	
	for _, submission := range assignmentSubmissions {
		submissionMap[submission.UserID] = submission
	}

	for _, member := range members {
		if member.Role == "teacher" {
			continue
		}
		if submissionMap[member.User_userID] == nil {
			submissionMapping[member.User_userID] = &entities.AssignmentSubmission{
				ID: uuid.Nil,
				UserID: member.User_userID,
				AssignmentID: assignmentID,
				IDFile: "",
				Score: 0,
				FileName: "",
				Status: entities.StatusTodo,
				CreatedAt: member.CreatedAt,
				UpdatedAt: member.UpdatedAt,
			}
		}else {
			submissionMapping[member.User_userID] = &entities.AssignmentSubmission{
				ID	:submissionMap[member.User_userID].ID,
				UserID: member.User_userID,
				AssignmentID: assignmentID,
				FileName: submissionMap[member.User_userID].FileName,
				IDFile: submissionMap[member.User_userID].IDFile,
				Score: submissionMap[member.User_userID].Score,
				Status: submissionMap[member.User_userID].Status,
				CreatedAt: member.CreatedAt,
				UpdatedAt: member.UpdatedAt,
			}
		}
	}
	
	var res []dto.GetAssSubmissionStudentResponse
	for _, member := range members {
		if member.Role == "teacher" {
			continue
		}
		result := dto.GetAssSubmissionStudentResponse{}
		mem := submissionMapping[member.User_userID]
		photoUrl := os.Getenv("GATEWAY_URL") + "/storage/user_profile_pictures" + member.User_userID.String()+".jpg"
		if mem.ID == uuid.Nil {
			result.ID = nil
			result.User_userID = member.User_userID
			result.PhotoUrl = &photoUrl
			result.Username = member.Username
			result.Status = entities.AssStatus("todo")
			result.LinkFile = nil
			result.Filename = nil
			result.Score = 0
			result.CreatedAt = nil
			result.UpdatedAt = nil
		}else {
			params := url.Values{}
			params.Add("id", mem.IDFile)
			link := os.Getenv("GATEWAY_URL") + "/teacher/student-assignment/?" + params.Encode()
			result.PhotoUrl = &photoUrl
			result.ID = &mem.ID
			result.User_userID = member.User_userID
			result.Username = member.Username
			result.Status = mem.Status
			result.LinkFile = &link
			result.Filename = &mem.FileName
			result.Score = mem.Score
			result.CreatedAt = &mem.CreatedAt
			result.UpdatedAt = &mem.UpdatedAt
		}
		if mem.Status == entities.AssStatus(status){
			res = append(res,result)
		}else if status == "" {
		res = append(res, result)
		}
	}

	return res, nil
}

// delete assignment submission by id
func (service *assignmentSubmissionService) DeleteAssignmentSubmissionByID(ctx context.Context, assignmentSubmissionID uuid.UUID) error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	
	// delete file from storage
	
	assignmentSubmission, err := service.assignmentSubmissionRepo.GetAssignmentSubmissionByID(ctx, nil, assignmentSubmissionID)
	if err != nil {
		return err
	}
	if assignmentSubmission.ID == uuid.Nil {
		return fmt.Errorf("assignment submission not found")
	}
	params := url.Values{}
	params.Add("id", assignmentSubmission.IDFile)
	link := os.Getenv("CONTENT_URL") + "/student-assignment/?" + params.Encode()
	req, err := http.NewRequest(http.MethodDelete, link, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	err = service.assignmentSubmissionRepo.DeleteAssignmentSubmission(ctx, nil, assignmentSubmissionID)
	if err != nil {
		return err
	}
	return nil
}
		


func (service *assignmentSubmissionService) uploadFile(file io.Reader, fileName string, userID string) (string, error) {
	type FileUploadResponse struct {
		Id string `json:"id"`
	}
	
	// fmt.Printf("Processing file upload\n")
	
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
	params := url.Values{}
	params.Add("userID", userID)
	url := os.Getenv("CONTENT_URL") + "/student-assignment/?" + params.Encode()
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	
	// Set proper content type with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	// fmt.Printf("Sending file upload request\n")
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	
	// fmt.Printf("Received response with status: %d\n", resp.StatusCode)
	
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