package migration

import (
	entities "LMSGo/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.AutoMigrate(&entities.Assignment{})
	db.AutoMigrate(&entities.ItemPembelajaran{})
	db.AutoMigrate(&entities.Kelas{})
	db.AutoMigrate(&entities.Member{})
	db.AutoMigrate(&entities.Submission{})
	db.AutoMigrate(&entities.Week{})
	
	return nil
}
