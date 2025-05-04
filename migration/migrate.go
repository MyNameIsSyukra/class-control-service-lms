package migration

import (
	entities "LMSGo/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.AutoMigrate(&entities.Kelas{})
	db.AutoMigrate(&entities.ItemPembelajaran{})
	db.AutoMigrate(&entities.Member{})
	
	return nil
}
