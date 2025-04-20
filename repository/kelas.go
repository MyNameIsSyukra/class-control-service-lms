package repository

import (
	entities "LMSGo/entity"

	"gorm.io/gorm"
)

type(
	KelasRepository interface {
			GetAll() ([]*entities.Kelas, error)
			GetById(id string)(*entities.Kelas, error)
			Create(class *entities.Kelas) error
			Update(id string, class *entities.Kelas) error
			Delete(id string) error
		}
	kelasRepository struct {
		db *gorm.DB
	}
) 

func NewKelasRepository(db *gorm.DB) *kelasRepository {
	return &kelasRepository{db}
}
func (repo *kelasRepository) GetAll() ([]*entities.Kelas, error) {
	var kelas []*entities.Kelas
	if err := repo.db.Find(&kelas).Error; err != nil {
		return nil, err
	}
	return kelas, nil
}


func (repo *kelasRepository) GetById(id string) (*entities.Kelas, error) {
	var kelas entities.Kelas
	if err := repo.db.Where("id = ?", id).Find(&kelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (repo *kelasRepository) Create(kelas *entities.Kelas) error {
	if err := repo.db.Create(kelas).Error; err != nil {
		return err
	}
	return nil
}

func (repo *kelasRepository) Update(id string, kelas *entities.Kelas) error {
	if err := repo.db.Where("id = ?", id).Updates(kelas).Error; err != nil {
		return err
	}
	return nil
}

func (repo *kelasRepository) Delete(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.Kelas{}).Error; err != nil {
		return err
	}
	return nil
}