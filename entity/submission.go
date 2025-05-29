package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssignmentSubmission struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	AssignmentID int            `json:"assignment_id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	IDFile       string         `json:"id_file"`
	FileName	string         `json:"file_name"`
	Score        int            `json:"score"`
	Status  	AssStatus      `json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Assignment *Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"assignment,omitempty"`
}


type AssStatus string


const (
	StatusLate AssStatus = "Late"
	StatusSubmitted  AssStatus = "submitted"
	StatusTodo 	 AssStatus = "todo"
)


// type AssignmentSubmission struct {
// 	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	AssignmentID int      `json:"assignment_id"`
// 	UserID    uuid.UUID      `gorm:"type:uuid" json:"user_id"`
// 	IDFile 	string         `json:"id_file"`
// 	Score 	 int           `json:"score"`
// 	// Member      Member        `gorm:"foreignKey:MemberID" json:"member"`
// 	CreatedAt   time.Time      `json:"created_at"`
// 	UpdatedAt   time.Time      `json:"updated_at"`
// 	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

// 	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID" json:"kelas"`
// }  