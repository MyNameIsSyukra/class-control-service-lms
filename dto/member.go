package dto

import (
	entities "LMSGo/entity"

	"github.com/google/uuid"
)

type AddMemberRequest struct {
	Username      string  `json:"username" binding:"required"`
	User_userID   uuid.UUID `json:"user_user_id" binding:"required"`
	Role entities.MemberRole `json:"role" binding:"required"`
	Kelas_kelasID uuid.UUID  `json:"kelas_kelas_id" binding:"required"`
}

type GetMemberResponse struct {
	Username      string  `json:"username"`
	User_userID   uuid.UUID `json:"user_user_id"`
	Role          entities.MemberRole `json:"role"`
	Kelas_kelasID uuid.UUID  `json:"kelas_kelas_id"`
	PhotoUrl  string  `json:"photo_url"`
}
