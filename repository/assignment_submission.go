package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"

	"gorm.io/gorm"
)

type (
	AssignmentSubmissionRepository interface {
		CreateAssignmentSubmission(ctx context.Context, tx *gorm.DB, assignmentSubmissionReq dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error)
	}
	assignmentSubmissionRepository struct {
		db *gorm.DB
	}
)

func NewAssignmentSubmissionRepository(db *gorm.DB) *assignmentSubmissionRepository {
	return &assignmentSubmissionRepository{db}
}

func (repo *assignmentSubmissionRepository) CreateAssignmentSubmission(ctx context.Context, tx *gorm.DB, assignmentSubmissionReq dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error) {
	var assignmentSubmission entities.AssignmentSubmission
	assignmentSubmission.AssignmentID = assignmentSubmissionReq.AssignmentID
	assignmentSubmission.UserID = assignmentSubmissionReq.UserID
	assignmentSubmission.IDFile = assignmentSubmissionReq.IDFile

	if err := repo.db.Create(&assignmentSubmission).Error; err != nil {
		return nil, err
	}
	return &assignmentSubmission, nil
}