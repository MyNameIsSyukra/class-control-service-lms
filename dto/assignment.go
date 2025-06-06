package dto

import (
	entities "LMSGo/entity"
	"time"
)

type AssignmentRequest struct {
	WeekID      int     `form:"week_id" json:"week_id"`
	Title       string  `form:"title" json:"title"`
	Description string  `form:"description" json:"description"`
	Deadline    time.Time `form:"deadline" json:"deadline"`
}

type InitUpdateAssignmentRequest struct {
	AssignmentID int     `form:"assignment_id"`
	WeekID      int     `form:"week_id"`
	Title       string  `form:"title"`
	Description string  `form:"description"`
	Deadline    time.Time `form:"deadline"`
}
type ProrcessedUpdateAssignmentRequest struct {
	AssignmentID int     `form:"assignment_id"`
	WeekID      int     `form:"week_id"`
	Title       string  `form:"title"`
	Description string  `form:"description"`
	Deadline    time.Time `form:"deadline"`
	FileName    string  `form:"file_name"`
	FileId	string  `form:"file_id"`
}

type CreateAssignmentRequest struct {
	WeekID      int     `json:"week_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    string  `json:"file_name"`
	FileId	string  `json:"file_id"`
	// FileID      string  `json:"file_id"`
}

type AssignmentResponse struct {
	AssignmentID      int     `json:"assignment_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    *string  `json:"file_name"`
	FileId	*string  `json:"file_id"`
	FileUrl  *string  `json:"file_url"`
}


type StudentGetAssignmentByIDResponse struct {
	ID 		int     `json:"assignment_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    string  `json:"file_name"`
	FileLink	string  `json:"file_link_assignment"`
	StudentSubmissionLink *string `json:"file_link_submission"`
	StudentSubmissionFileName *string `json:"file_name_submission"`
	Status entities.AssStatus `json:"status"`
	Score  int `json:"score"`
}