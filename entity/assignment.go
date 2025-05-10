package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	WeekID      int            `json:"week_id"`
	// CreatorID   uuid.UUID      `gorm:"type:uuid" json:"creator_id"`
	// Creator     Member        `gorm:"foreignKey:CreatorID" json:"creator"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// relationship with Week
	Week Week `gorm:"foreignKey:WeekID" json:"week"`
	Submissions []Submission `gorm:"foreignKey:AssignmentID" json:"submissions"`
}