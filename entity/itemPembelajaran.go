package entities

import (
	"gorm.io/gorm"
)

type ItemPembelajaran struct{
	gorm.Model
	IdItemPembelajaran int
	HeadingPertemuan string
	BodyPertemuan string
	FileName string
	FilePath string
	UrlVideo string
	Kelas_idKelas int
}

