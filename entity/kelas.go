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
	TeacherID   int    `json:"teacher_id"`

	// one to many relationship with member
	Members []Member `gorm:"foreignKey:Kelas_kelasID;references:ID"`
}

