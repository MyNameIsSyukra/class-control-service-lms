package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	"LMSGo/repository"
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	AssignmentService interface {
		CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest) (*entities.Assignment, error)
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

func (service *assignmentService) CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest) (*entities.Assignment, error) {	
	newAssignment, err := service.assignmentRepo.CreateAssignment(ctx, nil, request)
	if err != nil {
		return &entities.Assignment{}, err
	}
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