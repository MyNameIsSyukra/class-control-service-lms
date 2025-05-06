package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemPembelajaran struct{
	gorm.Model
	IdItemPembelajaran int `json:"idItemPembelajaran"`
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan string `json:"bodyPertemuan"`
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
	UrlVideo string `json:"urlVideo"`
	Kelas_idKelas uuid.UUID `json:"kelas_idKelas"`

	// Relationship with Kelas
	Kelas Kelas `gorm:"foreignKey:Kelas_idKelas" json:"-"`
}

