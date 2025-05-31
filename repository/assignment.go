package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"

	"gorm.io/gorm"
)

type (
	AssignmentRepository interface {
		// GetAllAssignmentByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Assignment, error)
		GetAssignmentByID(ctx context.Context, tx *gorm.DB, assignmentID int) (entities.Assignment, error)
		CreateAssignment(ctx context.Context, tx *gorm.DB, assignmentReq dto.CreateAssignmentRequest) (*entities.Assignment, error)
		UpdateAssignment(ctx context.Context, tx *gorm.DB, assignmentID int, assignmentReq dto.ProrcessedUpdateAssignmentRequest) (*entities.Assignment, error)
		// DeleteAssignment(ctx context.Context, tx *gorm.DB, assignmentID int) error
		// UpdateAssignment(ctx context.Context, tx *gorm.DB, assignmentID int, assignmentReq dto.AssignmentRequest) (*entities.Assignment, error)
		// GetAssignmentByWeekID(ctx context.Context, tx *gorm.DB, weekID int) ([]*entities.Assignment, error)
	}
	assignmentRepository struct {
		db *gorm.DB
	}
)

func NewAssignmentRepository(db *gorm.DB) *assignmentRepository {
	return &assignmentRepository{db}
}

func (repo *assignmentRepository) CreateAssignment(ctx context.Context, tx *gorm.DB, assignmentReq dto.CreateAssignmentRequest) (*entities.Assignment, error) {
	var assignment entities.Assignment
	assignment.WeekID = assignmentReq.WeekID
	assignment.Title = assignmentReq.Title
	assignment.Description = assignmentReq.Description
	assignment.Deadline = assignmentReq.Deadline
	assignment.FileName = assignmentReq.FileName
	assignment.FileId = assignmentReq.FileId

	if err := repo.db.Create(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (repo *assignmentRepository) GetAssignmentByID(ctx context.Context, tx *gorm.DB, assignmentID int) (entities.Assignment, error) {
	var assignment entities.Assignment
	if err := repo.db.Where("id = ?", assignmentID).Preload("Week").First(&assignment).Error; err != nil {
		return entities.Assignment{}, err
	}
	return assignment, nil
}


// update
func (repo *assignmentRepository) UpdateAssignment(ctx context.Context, tx *gorm.DB, assignmentID int, assignmentReq dto.ProrcessedUpdateAssignmentRequest) (*entities.Assignment, error) {
	var assignment entities.Assignment
	if err := repo.db.Where("id = ?", assignmentID).First(&assignment).Error; err != nil {
		return nil, err
	}

	assignment.Title = assignmentReq.Title
	assignment.Description = assignmentReq.Description
	assignment.Deadline = assignmentReq.Deadline
	assignment.FileName = assignmentReq.FileName
	assignment.FileId = assignmentReq.FileId

	if err := repo.db.Save(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}