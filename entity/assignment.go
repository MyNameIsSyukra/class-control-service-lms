package entities

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	FileName    string         `json:"file_name"`
	FileId	string         `json:"fileId"`
	WeekID          int            `json:"WeekdID"` // same as Week.ID
	

	// one-to-one with Week
	Week *Week `gorm:"foreignKey:WeekID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	Submissions []AssignmentSubmission `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"submissions,omitempty"`
}
