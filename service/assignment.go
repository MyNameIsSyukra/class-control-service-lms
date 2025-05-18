package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	"LMSGo/repository"
	"context"
)

type (
	AssignmentService interface {
		CreateAssignment(ctx context.Context, request dto.CreateAssignmentRequest) (*entities.Assignment, error)
		GetAssignmentByID(ctx context.Context, assignmentID int) (entities.Assignment, error)
	}
	assignmentService struct {
		assignmentRepo repository.AssignmentRepository
	}
)

func NewAssignmentService(assignmentRepo repository.AssignmentRepository) AssignmentService {
	return &assignmentService{assignmentRepo}
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