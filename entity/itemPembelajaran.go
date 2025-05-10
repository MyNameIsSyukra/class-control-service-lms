package entities

import (
	"gorm.io/gorm"
)

type ItemPembelajaran struct{
	gorm.Model
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan string `json:"bodyPertemuan"`
	UrlVideo string `json:"urlVideo"`
	WeekID int `json:"week_id"`
	FileName string `json:"fileName"`
	LinkFile string `json:"linkFile"`
	
	// relationship with Week
	Week Week `gorm:"foreignKey:WeekID" json:"week"`

	// Kelas_idKelas uuid.UUID `json:"kelas_idKelas"`
	// Relationship with Kelas
	// Kelas Kelas `gorm:"foreignKey:Kelas_idKelas" json:"-"`
}


// FileName string `json:"fileName"`
// FilePath string `json:"filePath"`