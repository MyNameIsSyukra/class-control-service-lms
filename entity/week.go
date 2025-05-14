package entities

import "github.com/google/uuid"

type Week struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	WeekNumber     int       `json:"week_number"`
	Kelas_idKelas  uuid.UUID `json:"class_id"`

	// many-to-one with Kelas
	Kelas Kelas `gorm:"foreignKey:Kelas_idKelas;references:ID" json:"kelas,omitempty"`

	// one-to-one with Assignment
	Assignment Assignment `gorm:"foreignKey:WeekID;references:ID" json:"assignment,omitempty"`

	// one-to-one with ItemPembelajaran
	ItemPembelajaran ItemPembelajaran `gorm:"foreignKey:WeekID;references:ID" json:"item_pembelajaran,omitempty"`
}
