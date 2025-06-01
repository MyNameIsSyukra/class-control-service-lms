package dto

import (
	"github.com/google/uuid"
)

type WeekRequest struct {
	WeekNumber     int       `json:"week_number"`
	Kelas_idKelas  uuid.UUID `json:"class_id"`
}

type WeekResponse struct {
	WeekID           int    `json:"week_id"`
	WeekNumber       int    `json:"week_number"`
	KelasID 		uuid.UUID `json:"class_id"`
	ItemPembelajarans *ItemPembelajaranResponse `json:"item_pembelajaran"`
	Assignment        []AssignmentResponse        `json:"assignment"`
}
