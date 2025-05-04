package dto

import "github.com/google/uuid"

type AddMemberRequest struct {
	Username      string  `json:"username" binding:"required"`
	User_userID   uuid.UUID `json:"user_user_id" binding:"required"`
	Kelas_kelasID uuid.UUID  `json:"kelas_kelas_id" binding:"required"`
}
