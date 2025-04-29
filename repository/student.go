package repository

import (
	entities "LMSGo/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type(
	StudentRepository interface {
		AddStudentToClass(ctx context.Context, tx *gorm.DB, student *entities.Member) (*entities.Member, error) 
		GetAllMembers() ([]*entities.Member, error)
		GetMemberById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Member, error)
		// UpdateMember(ctx context.Context, tx *gorm.DB, id string, member *entities.Member) (*entities.Member, error)
		DeleteMember(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
		// GetMemberByClassID(ctx context.Context, tx *gorm.DB, classID string) (*entities.Member, error)
		// GetMemberByUserID(ctx context.Context, tx *gorm.DB, userID string) (*entities.Member, error)
		// GetMemberByClassIDAndUserID(ctx context.Context, tx *gorm.DB, classID string, userID string) (*entities.Member, error)
	}
	studentRepository struct {
		db *gorm.DB
	}
)

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

func (repo *studentRepository) AddStudentToClass(ctx context.Context, tx *gorm.DB, student *entities.Member) (*entities.Member, error) {
	if err := repo.db.Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *studentRepository) GetAllMembers() ([]*entities.Member, error) {
	var members []*entities.Member
	if err := repo.db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (repo *studentRepository) GetMemberById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Member, error) {
	var member entities.Member
	if err := repo.db.Where("id = ?", id).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (repo *studentRepository) DeleteMember(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.Member{}).Error; err != nil {
		return err
	}
	return nil
}