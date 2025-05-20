package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AssignmentSubmissionRepository interface {
		// student
		CreateAssignmentSubmission(ctx context.Context, tx *gorm.DB, assignmentSubmissionReq dto.AssignmentSubmissionRequest) (*entities.AssignmentSubmission, error)
		CheckStudentSubmssionByAssIdUserID(ctx context.Context, tx *gorm.DB,assignmentId int, userID uuid.UUID) (entities.AssStatus,int, error)

		// teacher
		GetAllSubmissionByAssignmentID(ctx context.Context, tx *gorm.DB, assignmentID int) ([]*entities.AssignmentSubmission, error)
		UpdateStudentSubmissionScore(ctx context.Context, tx *gorm.DB, score int, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error)
		GetAssignmentSubmissionByID(ctx context.Context, tx *gorm.DB, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error)
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
		Status: 	 assignmentSubmission.Status,
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

func (repo *assignmentSubmissionRepository) UpdateStudentSubmissionScore(ctx context.Context, tx *gorm.DB, score int, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error) {
	var assignmentSubmission entities.AssignmentSubmission
	if err := repo.db.Where("id = ?", assignmentSubmissionID).First(&assignmentSubmission).Error; err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	assignmentSubmission.Score = score
	if err := repo.db.Save(&assignmentSubmission).Error; err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	res := entities.AssignmentSubmission{
		ID:           assignmentSubmission.ID,
		AssignmentID: assignmentSubmission.AssignmentID,
		UserID:       assignmentSubmission.UserID,
		IDFile:       assignmentSubmission.IDFile,
		Status: 	 assignmentSubmission.Status,
		Score:        assignmentSubmission.Score,
		CreatedAt:     assignmentSubmission.CreatedAt,
		UpdatedAt:     assignmentSubmission.UpdatedAt,
		Assignment: nil,
	}
	return &res, nil
}

func (repo *assignmentSubmissionRepository) CheckStudentSubmssionByAssIdUserID(ctx context.Context, tx *gorm.DB,assignmentId int, userID uuid.UUID) (entities.AssStatus,int, error) {
	var assignmentSubmissions entities.AssignmentSubmission
	if err := repo.db.Where("assignment_id = ? AND user_id = ?",assignmentId, userID).Find(&assignmentSubmissions).Error; err != nil {
		return entities.StatusTodo, 0, err
	}
	if assignmentSubmissions.Status == "" {
		return entities.StatusTodo, 0, nil
	}
	return assignmentSubmissions.Status,assignmentSubmissions.Score, nil
}

// get subbmssion by id 
func (repo *assignmentSubmissionRepository) GetAssignmentSubmissionByID(ctx context.Context, tx *gorm.DB, assignmentSubmissionID uuid.UUID) (*entities.AssignmentSubmission, error) {
	var assignmentSubmission entities.AssignmentSubmission
	if err := repo.db.Where("id = ?", assignmentSubmissionID).First(&assignmentSubmission).Error; err != nil {
		return &entities.AssignmentSubmission{}, err
	}
	res := entities.AssignmentSubmission{
		ID:           assignmentSubmission.ID,
		AssignmentID: assignmentSubmission.AssignmentID,
		UserID:       assignmentSubmission.UserID,
		IDFile:       assignmentSubmission.IDFile,
		Status: 	 assignmentSubmission.Status,
		Score:        assignmentSubmission.Score,
		CreatedAt:     assignmentSubmission.CreatedAt,
		UpdatedAt:     assignmentSubmission.UpdatedAt,
		Assignment: nil,
	}
	return &res, nil
}
