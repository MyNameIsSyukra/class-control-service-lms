package entities

type ItemPembelajaran struct {
	WeekID           int    `json:"id"` // same as Week.ID
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan    string `json:"bodyPertemuan"`
	UrlVideo         string `json:"urlVideo"`
	FileName         string `json:"fileName"`
	// FileID           string `json:"file_id"`
	FileId string `json:"fileId"`
	// one-to-one with Week
	Week *Week `gorm:"foreignKey:WeekID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// FileName string `json:"fileName"`
// FilePath string `json:"filePath"`