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
	Assignment        *AssignmentResponse        `json:"assignment"`
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
