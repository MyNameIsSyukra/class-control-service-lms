package entities

type ItemPembelajaran struct {
	WeekID           int    `json:"id"` // same as Week.ID
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan    string `json:"bodyPertemuan"`
	UrlVideo         string `json:"urlVideo"`
	FileName         string `json:"fileName"`
	FileLink         string `json:"linkFile"`

	// one-to-one with Week
	Week *Week `gorm:"foreignKey:WeekID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"week"`
}

// FileName string `json:"fileName"`
// FilePath string `json:"filePath"`