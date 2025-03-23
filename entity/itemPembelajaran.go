package entities

import (
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
	Kelas_idKelas int `json:"kelas_idKelas"`
}

