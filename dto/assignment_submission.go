package dto

import (
	entities "LMSGo/entity"
	"time"

	"github.com/google/uuid"
)
type AssignmentSubmissionRequest struct {
	AssignmentID int       `json:"assignment_id"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	IDFile       string    `json:"id_file"`
	// Score        int       `json:"score"`
}

type GetAssSubmissionStudentResponse struct {
	ID 		 *uuid.UUID `gorm:"type:uuid" json:"id_submission"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`	
	Username string `json:"username"`
	Status entities.AssStatus `json:"status"`
	LinkFile *string `json:"link_file"`
	Score int `json:"score"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}	
