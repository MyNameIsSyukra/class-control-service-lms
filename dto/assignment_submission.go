package dto

import "github.com/google/uuid"

type AssignmentSubmissionRequest struct {
	AssignmentID int       `json:"assignment_id"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	IDFile       string    `json:"id_file"`
	Score        int       `json:"score"`
}