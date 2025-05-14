package dto

import (
	entities "LMSGo/entity"

	"github.com/google/uuid"
)

type WeekRequest struct {
	WeekNumber     int       `json:"week_number"`
	Kelas_idKelas  uuid.UUID `json:"class_id"`
}

type WeekResponseByID struct {
	WeekID           int    `json:"id"`
	WeekNumber       int    `json:"week_number"`
	KelasID 		uuid.UUID `json:"class_id"`
	ItemPembelajarans *entities.ItemPembelajaran `json:"item_pembelajaran,omitempty"`
	Assignment        *entities.Assignment        `json:"assignment,omitempty"`
}