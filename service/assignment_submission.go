package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	"LMSGo/repository"
	"context"
)

type (
	AssignmentSubmissionService interface {
		CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error)
	}
	assignmentSubmissionService struct {
		assignmentSubmissionRepo repository.AssignmentSubmissionRepository
	}
)

func NewAssignmentSubmissionService(assignmentSubmissionRepo repository.AssignmentSubmissionRepository) AssignmentSubmissionService {
	return &assignmentSubmissionService{assignmentSubmissionRepo}
}
func (service *assignmentSubmissionService) CreateAssignmentSubmission(ctx context.Context, request dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error) {
	newAssignmentSubmission, err := service.assignmentSubmissionRepo.CreateAssignmentSubmission(ctx, nil, request)
	if err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	return newAssignmentSubmission, nil
}
