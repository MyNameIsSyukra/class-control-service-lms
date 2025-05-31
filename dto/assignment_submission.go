package dto

import (
	entities "LMSGo/entity"
	"time"

	"github.com/google/uuid"
)
type InitAssignmentSubmissionRequest struct {
	AssignmentID int       `form:"assignment_id"`
	UserID       string `gorm:"type:uuid" form:"user_id"`
	IDFile       string    `form:"id_file"`
}
type AssignmentSubmissionRequest struct {
	AssignmentID int       `json:"assignment_id"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	IDFile       string    `json:"id_file"`
	FileName     string    `json:"file_name"`
}

type GetAssSubmissionStudentResponse struct {
	ID 		 *uuid.UUID `gorm:"type:uuid" json:"id_submission"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`	
	Username string `json:"username"`
	PhotoUrl *string `json:"photo_url"`
	Status entities.AssStatus `json:"status"`
	LinkFile *string `json:"link_file"`
	Filename *string `json:"filename"`
	Score int `json:"score"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}	
