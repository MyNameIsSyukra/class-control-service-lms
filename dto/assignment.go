package dto

import (
	entities "LMSGo/entity"
	"time"
)

type CreateAssignmentRequest struct {
	WeekID      int     `json:"week_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    string  `json:"file_name"`
	FileLink	string  `json:"file_link"`
	// FileID      string  `json:"file_id"`
}


type StudentGetAssignmentByIDResponse struct {
	WeekID 	int     `json:"week_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    string  `json:"file_name"`
	FileLink	string  `json:"file_link_assignment"`
	StudentSubmissionLink string `json:"file_link_submission"`
	// FileID    string  `json:"file_id"`
	Status entities.AssStatus `json:"status"`
	Score  int `json:"score"`
}