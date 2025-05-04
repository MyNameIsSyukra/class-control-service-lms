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
	TeacherID   int    `json:"teacher_id" binding:"required"`
}

type CreateKelasUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Teacher     string `json:"teacher" binding:"required"`
	TeacherID   int    `json:"teacher_id" binding:"required"`
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
	TeacherID   int    `json:"teacher_id"`
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
}



type GetClassAndAssignmentResponse struct {
	ClassID      uuid.UUID `json:"class_id"`
	ClassName    string    `json:"class_name"`
	ClassTag     string    `json:"class_tag"`
	ClassDesc    string    `json:"class_desc"`
	ClassTeacher string    `json:"class_teacher"`
	ClassTeacherID int       `json:"class_teacher_id"`
	ClassAssessment []AssessmentResponse `json:"class_assignments"`
}