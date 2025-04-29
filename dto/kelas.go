package dto

type CreateKelasRequest struct {
	Name        string `json:"name" binding:"required"`
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