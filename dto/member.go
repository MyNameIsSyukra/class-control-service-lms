package dto

import "github.com/google/uuid"

type AddMemberRequest struct {
	Username      string  `json:"username"`
	User_userID   uuid.UUID `json:"user_user_id"`
	Kelas_kelasID uuid.UUID  `json:"kelas_kelas_id"`
}
