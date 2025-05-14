package repository

import (
	"gorm.io/gorm"
)

type (
	AssignmentRepository interface {
		// GetAllAssignmentByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Assignment, error)
		// GetAssignmentByID(ctx context.Context, tx *gorm.DB, assignmentID int) (*entities.Assignment, error)
		// CreateAssignment(ctx context.Context, tx *gorm.DB, assignmentReq dto.AssignmentRequest) (*entities.Assignment, error)
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
