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
	ID 		 *uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	Username string `json:"username"`
	Role  entities.MemberRole `json:"role"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`	
	Status entities.AssStatus `json:"status"`
	Score int `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}