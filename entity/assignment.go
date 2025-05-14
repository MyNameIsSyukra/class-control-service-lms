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
	FileLink    string         `json:"path_file"`
	
	WeekID          int            `json:"id"` // same as Week.ID

	// one-to-one with Week
	Week *Week `gorm:"foreignKey:WeekID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"week"`

	Submissions []AssignmentSubmission `gorm:"foreignKey:AssignmentID;references:ID" json:"submissions"`
}
