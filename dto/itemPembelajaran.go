package dto

import "github.com/google/uuid"

// type ItemPembelajaranRequest struct {
// 	WeekID           int    `json:"id"` // same as Week.ID
// 	HeadingPertemuan string `json:"headingPertemuan"`
// 	BodyPertemuan    string `json:"bodyPertemuan"`
// 	UrlVideo         string `json:"urlVideo"`
// 	FileName         string `json:"fileName"`
// 	FileID         string `json:"file_id"`
// }

type CreateItemPembelajaranRequest struct {
	KelasID          uuid.UUID `json:"kelas_id" binding:"required"`
	WeekNumber       int    `json:"week_number" binding:"required"`
	HeadingPertemuan string `json:"headingPertemuan" binding:"required"`
	BodyPertemuan    string `json:"bodyPertemuan" binding:"required"`
	UrlVideo         string `json:"urlVideo"`
	FileName         string `json:"fileName"`
	FileLink         string `json:"file_link"`
}
type UpdateItemPembelajaranRequest struct {
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan    string `json:"bodyPertemuan"`
	UrlVideo         string `json:"urlVideo"`
	FileName         string `json:"fileName"`
	FileLink         string `json:"file_link"`
}
