package repository

import (
	entities "LMSGo/entity"
	"context"

	"gorm.io/gorm"
)

type(
	KelasRepository interface {
			Create(ctx context.Context, tx *gorm.DB,class *entities.Kelas) (*entities.Kelas,error)
			GetAll() ([]*entities.Kelas, error)
			GetById(ctx context.Context, tx *gorm.DB,id string)(*entities.Kelas, error)
			Update(ctx context.Context, tx *gorm.DB,id string, class *entities.Kelas) (*entities.Kelas,error)
			Delete(ctx context.Context, tx *gorm.DB,id string) error
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


func (repo *kelasRepository) GetById(ctx context.Context, tx *gorm.DB,id string)(*entities.Kelas, error) {
	var kelas entities.Kelas
	if err := repo.db.Where("id = ?", id).Find(&kelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (repo *kelasRepository) Create(ctx context.Context, tx *gorm.DB,class *entities.Kelas) (*entities.Kelas, error) {
	if err := repo.db.Create(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (repo *kelasRepository) Update(ctx context.Context, tx *gorm.DB,id string, class *entities.Kelas) (*entities.Kelas,error) {
	if err := repo.db.Where("id = ?", id).Updates(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (repo *kelasRepository) Delete(ctx context.Context, tx *gorm.DB,id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.Kelas{}).Error; err != nil {
		return err
	}
	return nil
}