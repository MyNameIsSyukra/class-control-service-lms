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

type ItemPembelajaranRequest struct {
	KelasIDStr          string `form:"kelas_id" json:"kelas_id" binding:"required"`
	WeekNumber       int    `form:"week_number" json:"week_number" binding:"required"`
	HeadingPertemuan string `form:"headingPertemuan" json:"headingPertemuan" binding:"required"`
	BodyPertemuan    string `form:"bodyPertemuan" json:"bodyPertemuan" binding:"required"`
	UrlVideo         string `form:"urlVideo" json:"urlVideo"`
}

type CreateItemPembelajaranRequest struct {
    KelasID          uuid.UUID `json:"kelas_id"`
    WeekNumber       int       `json:"week_number"`
    HeadingPertemuan string    `json:"headingPertemuan"`
    BodyPertemuan    string    `json:"bodyPertemuan"`
    UrlVideo         string    `json:"urlVideo"`
    FileName         string    `json:"fileName"`
    FileLink         string    `json:"file_link"`
}

type InitialUpdateItemPembelajaranRequest struct {
	WeekID 		int    `form:"week_id"`
	HeadingPertemuan string `form:"headingPertemuan"`
	BodyPertemuan    string `form:"bodyPertemuan"`
	UrlVideo         string `form:"urlVideo"`
}

type UpdateItemPembelajaranRequest struct {
	WeekID           int    `json:"week_id"` // same as Week.ID	
	HeadingPertemuan string `json:"headingPertemuan"`
	BodyPertemuan    string `json:"bodyPertemuan"`
	UrlVideo         string `json:"urlVideo"`
	FileName         string `json:"fileName"`
	FileLink         string `json:"file_link"`
}
