package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	// ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username  string    `json:"username"`
	Role  MemberRole `json:"role"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`
	Kelas_kelasID uuid.UUID `gorm:"type:uuid" json:"kelas_kelas_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Kelas Kelas `gorm:"foreignKey:Kelas_kelasID" json:"-"`
}

const (
	MemberRoleAdmin = "admin"
	MemberRoleStudent = "student"
	MemberRoleTeacher = "teacher"
)

type MemberRole string