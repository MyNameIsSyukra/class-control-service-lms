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

type ClassIDResponse struct {
	ID uuid.UUID `json:"id_kelas"`
	Name string    `json:"name"`
	Tag string     `json:"tag"`
	Description string `json:"description"`
	Teacher string `json:"teacher"`
	TeacherID uuid.UUID `json:"teacher_id"`

	Week []WeekResponse `json:"week"`
}

type WeekResponse struct {
	WeekID           int    `json:"id"`
	WeekNumber       int    `json:"week_number"`
	KelasID 		uuid.UUID `json:"class_id"`
	ItemPembelajarans *entities.ItemPembelajaran `json:"item_pembelajaran,omitempty"`
	Assignment        *entities.Assignment        `json:"assignment,omitempty"`
}
