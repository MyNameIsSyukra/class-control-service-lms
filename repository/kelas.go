package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type(
	KelasRepository interface {
			Create(ctx context.Context, tx *gorm.DB,class *entities.Kelas) (*entities.Kelas,error)
			GetAll(ctx context.Context,
				tx *gorm.DB,
				req dto.PaginationRequest,) (dto.GetAllKelasRepoResponse, error)
			GetById(ctx context.Context, tx *gorm.DB,id string)(*entities.Kelas, error)
			Update(ctx context.Context, tx *gorm.DB,id string, class *entities.Kelas) (*entities.Kelas,error)
			Delete(ctx context.Context, tx *gorm.DB,id string) error
			UpdateClassTeacherID(ctx context.Context, tx *gorm.DB, teacherId uuid.UUID, classID uuid.UUID, teacherName string) (*entities.Kelas,error)
		}
	kelasRepository struct {
		db *gorm.DB
	}
) 

func NewKelasRepository(db *gorm.DB) *kelasRepository {
	return &kelasRepository{db}
}

func (repo *kelasRepository) GetAll(ctx context.Context,
	tx *gorm.DB, req dto.PaginationRequest,) (dto.GetAllKelasRepoResponse, error) {
		if tx == nil {
			tx = repo.db
		}
	var kelas []entities.Kelas
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entities.Kelas{})
	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllKelasRepoResponse{}, err
	}
	if err := query.Scopes(Paginate(req)).Find(&kelas).Error; err != nil {
		return dto.GetAllKelasRepoResponse{}, err
	}
	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllKelasRepoResponse{
		Kelas: kelas,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	},err

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

func (repo *kelasRepository) UpdateClassTeacherID(ctx context.Context, tx *gorm.DB, teacherId uuid.UUID, classID uuid.UUID, teacherName string) (*entities.Kelas,error) {
	if err := repo.db.Where("id = ?", classID).Updates(&entities.Kelas{TeacherID: teacherId, Teacher: teacherName}).Error; err != nil {
		return nil, err
	}
	return &entities.Kelas{TeacherID: teacherId, ID: classID}, nil	
}