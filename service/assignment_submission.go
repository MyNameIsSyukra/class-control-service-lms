package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	"LMSGo/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	AssignmentSubmissionService interface {
		CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error)

		// teacher
		GetAllStudentAssignmentSubmissionByAssignmentID(ctx context.Context, assignmentID int) ([]dto.GetAssSubmissionStudentResponse, error)
		UpdateStudentSubmissionScore(ctx context.Context, score int, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error)
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
func (service *assignmentSubmissionService) CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error) {
	// check if the user has submitted the assignment
	assStatus, _, err := service.assignmentSubmissionRepo.CheckStudentSubmssionByAssIdUserID(ctx, nil, request.AssignmentID, request.UserID)
	if err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	if assStatus == entities.StatusSubmitted || assStatus == entities.StatusLate {
		return &entities.AssignmentSubmission{}, fmt.Errorf("you have submitted this assignment")
	}
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

func (service *assignmentSubmissionService) GetAllStudentAssignmentSubmissionByAssignmentID(ctx context.Context, assignmentID int) ([]dto.GetAssSubmissionStudentResponse, error) {
	assignment, err := service.assignmentRepo.GetAssignmentByID(ctx, nil, assignmentID)
	if err != nil {
		return nil, err
	}
	assignmentSubmissions, err := service.assignmentSubmissionRepo.GetAllSubmissionByAssignmentID(ctx, nil, assignmentID)
	if err != nil {
		return nil, err
	}
	
	members, err := service.memberRepo.GetAllMembersByClassID(ctx, nil, assignment.Week.Kelas_idKelas)
	if err != nil {
		return nil, err
	}

	submissionMap := make(map[uuid.UUID]*entities.AssignmentSubmission)
	for _, submission := range assignmentSubmissions {
		submissionMap[submission.UserID] = submission
	}
	
	if len(assignmentSubmissions) == 0 {
		for _, member := range members {
			if member.Role == "teacher" {
				continue
			}
			submissionMap[member.User_userID] = &entities.AssignmentSubmission{
				ID: uuid.Nil,
				UserID: member.User_userID,
				AssignmentID: assignmentID,
				Status: entities.StatusTodo,
				Score: 0,
				CreatedAt: member.CreatedAt,
				UpdatedAt: member.UpdatedAt,
			}
		}
	}
		
	var result []dto.GetAssSubmissionStudentResponse
	for _, member := range members {
		if member.Role == "teacher" {
			continue
		}
		// fmt.Println("member", member)
		mem := submissionMap[member.User_userID]
		fmt.Println("mem", mem.Status)
		if mem.Status == "" {
			mem.Status = entities.StatusTodo
		}
		dto := dto.GetAssSubmissionStudentResponse{
			ID:         &mem.ID,
			Username:   member.Username,
			Role: 	  member.Role,
			User_userID: member.User_userID,
			Status: 	mem.Status,
			Score:      mem.Score,
			CreatedAt: mem.CreatedAt,
			UpdatedAt: mem.UpdatedAt,
		}
		result = append(result, dto)
	}
	


	return result, nil
}
