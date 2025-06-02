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
	db.AutoMigrate(&entities.AssignmentSubmission{})
	db.AutoMigrate(&entities.Week{})
	
	return nil
}

func Rollback(db *gorm.DB) error {
	db.Migrator().DropTable(&entities.Assignment{})
	db.Migrator().DropTable(&entities.ItemPembelajaran{})
	db.Migrator().DropTable(&entities.Kelas{})
	db.Migrator().DropTable(&entities.Member{})
	db.Migrator().DropTable(&entities.AssignmentSubmission{})
	db.Migrator().DropTable(&entities.Week{})

	return nil
}
