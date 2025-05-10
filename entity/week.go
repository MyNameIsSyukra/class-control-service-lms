package entities

import "github.com/google/uuid"

type Week struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	WeekNumber int    `json:"week_number"`
	Tag        string `json:"tag"`
	Kelas_idKelas    uuid.UUID    `json:"class_id"`

	// many to one with class
	Kelas Kelas `gorm:"foreignKey:Kelas_idKelas" json:"kelas"`
	// one to many with assignment
	Assignments []Assignment `gorm:"foreignKey:WeekID" json:"assignments"`
	// one too many with item pembelajaran
	ItemPembelajaran []ItemPembelajaran `gorm:"foreignKey:WeekID" json:"item_pembelajaran"`
}