package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	AssignmentSubmissionRepository interface {
		CreateAssignmentSubmission(ctx context.Context, tx *gorm.DB, assignmentSubmissionReq dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error)
		
		// teacher
		GetAllSubmissionByAssignmentID(ctx context.Context, tx *gorm.DB, assignmentID int) ([]*entities.AssignmentSubmission, error)
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
	assignmentSubmission.Status = entities.StatusSubmitted
	// check assessment deadline
	var assesment entities.Assignment
	if err := repo.db.Where("id = ?", assignmentSubmission.AssignmentID).First(&assesment).Error; err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	
	// check if late submission
	if assesment.Deadline.Before(time.Now()) {
		assignmentSubmission.Status = entities.StatusLate
	} 

	if err := repo.db.Create(&assignmentSubmission).Error; err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	res := entities.AssignmentSubmission{
		ID:           assignmentSubmission.ID,
		AssignmentID: assignmentSubmission.AssignmentID,
		UserID:       assignmentSubmission.UserID,
		IDFile:       assignmentSubmission.IDFile,
		CreatedAt:     assignmentSubmission.CreatedAt,
		UpdatedAt:     assignmentSubmission.UpdatedAt,
		Assignment: nil,
	}
	return &res, nil
}

func (repo *assignmentSubmissionRepository) GetAllSubmissionByAssignmentID(ctx context.Context, tx *gorm.DB, assignmentID int) ([]*entities.AssignmentSubmission, error) {
	var assignmentSubmissions []*entities.AssignmentSubmission
	if err := repo.db.Where("assignment_id = ?", assignmentID).Find(&assignmentSubmissions).Error; err != nil {
		return []*entities.AssignmentSubmission{}, err
	}
	return assignmentSubmissions, nil
}

