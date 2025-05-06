package dto

import (
	entities "LMSGo/entity"
	"time"

	"github.com/google/uuid"
)

type CreateKelasRequest struct {
	Name        string `json:"name" binding:"required"`
	Tag         string `json:"tag" binding:"required"`
	Description string `json:"description" binding:"required"`
	Teacher     string `json:"teacher" binding:"required"`
	TeacherID   uuid.UUID    `json:"teacher_id" binding:"required"`
}

type KelasUpdateRequest struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
	TeacherID   uuid.UUID    `json:"teacher_id"`
}

type GetAllKelasRepoResponse struct {
	Kelas              []entities.Kelas            `json:"kelas"`
	PaginationResponse PaginationResponse `json:"pagination"`
}

type KelasResponse struct {
	ID          uuid.UUID    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
	TeacherID   uuid.UUID    `json:"teacher_id"`
}


type KelasPaginationResponse struct {
	Data []KelasResponse `json:"data"`
	PaginationResponse PaginationResponse `json:"pagination"`
}

type AssessmentResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	ClassID   uuid.UUID `gorm:"type:uuid" json:"class_id"`
	SubmissionStatus ExamStatus  `json:"submission_status"`
}



type GetClassAndAssignmentResponse struct {
	ClassID      uuid.UUID `json:"class_id"`
	ClassName    string    `json:"class_name"`
	ClassTag     string    `json:"class_tag"`
	ClassDesc    string    `json:"class_desc"`
	ClassTeacher string    `json:"class_teacher"`
	ClassTeacherID uuid.UUID       `json:"class_teacher_id"`
	ClassAssessment []AssessmentResponse `json:"class_assignments"`
}

type ExamStatus string


const (
	StatusInProgress ExamStatus = "in_progress"
	StatusSubmitted  ExamStatus = "submitted"
	StatusTodo 	 ExamStatus = "todo"
)