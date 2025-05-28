package entities

import (
	"github.com/google/uuid"
)

type Kelas struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string `json:"name"`
	Tag 	  string `json:"tag"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
	TeacherID   uuid.UUID    `json:"teacher_id"`

	// one to many relationship with member
	Members []Member `gorm:"foreignKey:Kelas_kelasID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members,omitempty"`

	// one to many relationship with week
	Weeks []Week `gorm:"foreignKey:Kelas_idKelas;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"weeks,omitempty"`
	// unused for now
	// one to many relationship with item pembelajaran
	// ItemPembelajaran []ItemPembelajaran `gorm:"foreignKey:Kelas_idKelas;references:ID"`
}

