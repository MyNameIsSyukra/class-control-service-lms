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
	FileLink    string  `json:"file_link"`
}


type StudentGetAssignmentByIDResponse struct {
	WeekID 	int     `json:"week_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    time.Time `json:"deadline"`
	FileName    string  `json:"file_name"`
	FileLink    string  `json:"file_link"`
	Status entities.AssStatus `json:"status"`
	Score  int `json:"score"`
}