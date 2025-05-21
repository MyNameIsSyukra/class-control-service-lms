package repository

import (
	"LMSGo/dto"
	entities "LMSGo/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type(
	StudentRepository interface {
		AddMemberToClass(ctx context.Context, tx *gorm.DB, student *entities.Member) (*entities.Member, error) 
		GetAllMembersByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Member, error)
		GetMemberById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Member, error)
		// UpdateMember(ctx context.Context, tx *gorm.DB, id string, member *entities.Member) (*entities.Member, error)
		DeleteMember(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
		GetAllClassByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entities.Kelas, error)
		GetAllClassAndAssesmentByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]dto.GetClassAndAssignmentResponse, error)
		GetMemberByClassIDAndUserID(ctx context.Context, tx *gorm.DB, classID uuid.UUID, userID uuid.UUID) (*entities.Member, error)
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

func (repo *studentRepository) AddMemberToClass(ctx context.Context, tx *gorm.DB, student *entities.Member) (*entities.Member, error) {
	// Check if the member already exists in the class
	// var flag bool = true
	data, err := repo.GetMemberByClassIDAndUserID(ctx, nil, student.Kelas_kelasID, student.User_userID)
	if err != nil {
		return nil, err
	}
	if data.Username != "" {
		return nil, errors.New("already In Class")
	}
	if err := repo.db.Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *studentRepository) GetAllMembersByClassID(ctx context.Context, tx *gorm.DB, classID uuid.UUID) ([]*entities.Member, error) {
	var members []*entities.Member
	if err := repo.db.Where("kelas_kelas_id = ?", classID).Find(&members).Error; err != nil {
		return []*entities.Member{}, err
	}
	return members, nil
}

func (repo *studentRepository) GetMemberById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Member, error) {
	var member entities.Member
	if err := repo.db.Where("user_user_id = ?", id).First(&member).Error; err != nil {
		return &entities.Member{}, err
	}
	return &member, nil
}

func (repo *studentRepository) GetMemberByClassIDAndUserID(ctx context.Context, tx *gorm.DB, classID uuid.UUID, userID uuid.UUID) (*entities.Member, error) {
	var member entities.Member
	if err := repo.db.Where("kelas_kelas_id = ? AND user_user_id = ?", classID, userID).Find(&member).Error; err != nil {
		return &entities.Member{}, err
	}
	return &member, nil
}

func (repo *studentRepository) DeleteMember(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.Member{}).Error; err != nil {
		return err
	}
	return nil
}

// get all class by user id
func (repo *studentRepository) GetAllClassByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entities.Kelas, error) {
	var members []entities.Member

	// Gunakan tx jika diberikan, kalau tidak pakai repo DB langsung
	db := repo.db
	if tx != nil {
		db = tx
	}

	// Ambil data member dengan preload relasi ke Kelas
	err := db.WithContext(ctx).
		Preload("Kelas").
		Where("user_user_id = ?", userID).
		Find(&members).Error
	if err != nil {
		return []entities.Kelas{}, err
	}

	// Ambil semua kelas dari member
	var kelasList []entities.Kelas
	for _, member := range members {
		kelasList = append(kelasList, member.Kelas)
	}

	return kelasList, nil
}

func (repo *studentRepository) GetAllClassAndAssesmentByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]dto.GetClassAndAssignmentResponse, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	urlAssessmentService := os.Getenv("ASSESSMENT_SERVICE_URL")

	listKelas, err := repo.GetAllClassByUserID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}
	var datas []dto.GetClassAndAssignmentResponse
	for _, kelas := range listKelas {
		var data dto.GetClassAndAssignmentResponse
		// Ambil data assesment dari kelas
		data.ClassID = kelas.ID
		data.ClassName = kelas.Name
		data.ClassTag = kelas.Tag
		data.ClassDesc = kelas.Description
		data.ClassTeacher = kelas.Teacher
		data.ClassTeacherID = kelas.TeacherID
		url := fmt.Sprintf("%s/service/assessment/class/%s/%s",urlAssessmentService, kelas.ID,userID)
		// Lakukan HTTP GET request ke URL
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		// Baca body response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// fmt.Print(body)

		// Unmarshal JSON ke struct
		var assessments []dto.AssessmentResponse
		err = json.Unmarshal(body, &assessments)
		if err != nil {
			assessments = nil
		}
		data.ClassAssessment = assessments
		datas = append(datas, data)
	}
	return datas, nil
}