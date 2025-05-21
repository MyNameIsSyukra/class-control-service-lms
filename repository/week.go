package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	WeekRepository interface {
		GetAllWeekByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Week, error)
		GetWeekByID(ctx context.Context, tx *gorm.DB, weekID int) (entities.Week, error)

		// teacher
		CreateWeeklySection(ctx context.Context, tx *gorm.DB, weekReq dto.WeekRequest) (*entities.Week, error)
		CreateItemPembelajaran(ctx context.Context, tx *gorm.DB, item *entities.ItemPembelajaran)(*entities.ItemPembelajaran, error)
		DeleteWeeklySection(ctx context.Context, tx *gorm.DB, weekID int) error
		// CreateWeek(ctx context.Context, tx *gorm.DB, week *entities.Week) (*entities.Week, error)
		// DeleteWeek(ctx context.Context, tx *gorm.DB, weekID int) error
	}
	weekRepository struct {
		db *gorm.DB
	}
)

func NewWeekRepository(db *gorm.DB) *weekRepository {
	return &weekRepository{db}
}

func (repo *weekRepository) GetAllWeekByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Week, error) {
	var weeks []*entities.Week
	if err := repo.db.Where("kelas_id_kelas = ?", classID).Preload("ItemPembelajaran").Preload("Assignment").Find(&weeks).Error; err != nil {
		return []*entities.Week{}, err
	}
	return weeks, nil
}

func (repo *weekRepository) GetWeekByID(ctx context.Context, tx *gorm.DB, weekID int) (entities.Week, error) {
	var week entities.Week
	if err := repo.db.Where("id = ?", weekID).Preload("ItemPembelajaran").Preload("Assignment").First(&week).Error; err != nil {
		return entities.Week{}, err
	}
	return week, nil
}

func (repo *weekRepository) CreateWeeklySection(ctx context.Context, tx *gorm.DB, weekReq dto.WeekRequest) (*entities.Week, error){
	var week entities.Week
	week.WeekNumber = weekReq.WeekNumber
	week.Kelas_idKelas = weekReq.Kelas_idKelas

	err := repo.db.Create(&week).Error
	if err != nil {
		return nil, err
	}
	return &week, nil
} 

func (repo *weekRepository) CreateItemPembelajaran(ctx context.Context, tx *gorm.DB, item *entities.ItemPembelajaran) (*entities.ItemPembelajaran, error) {
	if err := repo.db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// deleteWeeklySection(ctx context.Context, tx *gorm.DB, weekID int) error {
func (repo *weekRepository) DeleteWeeklySection(ctx context.Context, tx *gorm.DB, weekID int) error {
	if err := repo.db.Where("id = ?", weekID).Delete(&entities.Week{}).Error; err != nil {
		return err
	}
	return nil
}
		