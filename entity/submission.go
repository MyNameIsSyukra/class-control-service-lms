package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Submission struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	AssignmentID uuid.UUID      `gorm:"type:uuid" json:"assignment_id"`
	UserID    uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	IDFile 	string         `json:"id_file"`
	Score 	 int           `json:"score"`
	// Member      Member        `gorm:"foreignKey:MemberID" json:"member"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Assignment Assignment `gorm:"foreignKey:AssignmentID" json:"assignment"`
}  