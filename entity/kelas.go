package entities

import (
	"github.com/google/uuid"
)

type Kelas struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
	TeacherID   int    `json:"teacher_id"`
}

