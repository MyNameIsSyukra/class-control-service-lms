package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	WeekService interface {
		GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) (dto.ClassIDResponse, error)
		GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponseByID, error)

		// teacher
		CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest) (*entities.ItemPembelajaran, error)
		DeleteWeeklySection(ctx context.Context, weekID int) error
		UpdateWeeklySection(ctx context.Context,classId uuid.UUID, week_number int, req dto.UpdateItemPembelajaranRequest) (*entities.ItemPembelajaran, error)
		// CreateWeeklySection(ctx context.Context, weekReq dto.WeekRequest) (*entities.Week, error)
		// DeleteWeeklySection(ctx context.Context, weekID int) error
	}
	weekService struct {
		weekRepo database.WeekRepository
		kelasRepo database.KelasRepository
	}
)

func NewWeekService(weekRepo database.WeekRepository, kelasRepo database.KelasRepository) WeekService {
	return &weekService{weekRepo , kelasRepo}
}

func (service *weekService) GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) (dto.ClassIDResponse, error) {
	class, err := service.kelasRepo.GetById(ctx, nil, classID)
	if err != nil {
		return dto.ClassIDResponse{}, err
	}
	if class.ID == uuid.Nil {
		return dto.ClassIDResponse{}, fmt.Errorf("class with ID %s not found", classID)
	}
	weeks, err := service.weekRepo.GetAllWeekByClassID(ctx, nil, classID)
	if err != nil {
		return dto.ClassIDResponse{
			ID: class.ID,
			Name: class.Name,
			Tag: class.Tag,
			Description: class.Description,
			Teacher: class.Teacher,
			TeacherID: class.TeacherID,
			Week : nil,
		}, err
	}
	if len(weeks) == 0 {
		return dto.ClassIDResponse{
			ID: class.ID,
			Name: class.Name,
			Tag: class.Tag,
			Description: class.Description,
			Teacher: class.Teacher,
			TeacherID: class.TeacherID,
			Week : nil,
		}, nil
	}
		
	var weekResponse []dto.WeekResponse
	for _, week := range weeks {
		var weekRes dto.WeekResponse
		if week.Assignment.Title == "" {
			fmt.Println("Assignment is empty")
			weekRes.Assignment = nil
			weekRes.ItemPembelajarans = &week.ItemPembelajaran
		}else if week.ItemPembelajaran.HeadingPertemuan == "" {
			weekRes.ItemPembelajarans = nil
			weekRes.Assignment = &week.Assignment
		}else {
			weekRes.ItemPembelajarans = &week.ItemPembelajaran
			weekRes.Assignment = &week.Assignment
		}
		weekRes.WeekID = week.ID
		weekRes.WeekNumber = week.WeekNumber
		weekRes.KelasID = week.Kelas_idKelas
		weekResponse = append(weekResponse, weekRes)
	}		


	return dto.ClassIDResponse{
		ID: class.ID,
		Name: class.Name,
		Tag: class.Tag,
		Description: class.Description,
		Teacher: class.Teacher,
		TeacherID: class.TeacherID,
		Week: weekResponse,
	}, nil
}

func (service *weekService) GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponseByID, error) {
	week, err := service.weekRepo.GetWeekByID(ctx, nil, weekID)
	if err != nil {
		return dto.WeekResponseByID{}, err
	}
	resp := dto.WeekResponseByID{
		WeekID:           week.ID,
		WeekNumber:       week.WeekNumber,
		KelasID:          week.Kelas_idKelas,
		ItemPembelajarans: &week.ItemPembelajaran,
		Assignment:       &week.Assignment,
	}
	if week.Assignment.Title == "" {
		resp.Assignment = nil
	}
	if week.ItemPembelajaran.HeadingPertemuan == "" {
		resp.ItemPembelajarans = nil
	}
	return resp, nil
}

func (service *weekService) CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest) (*entities.ItemPembelajaran, error) {
	class , err := service.kelasRepo.GetById(ctx, nil, request.KelasID)
	if err != nil {
		return nil, fmt.Errorf("class with ID %s not found", request.KelasID)
	}
	if class.ID == uuid.Nil {
		return nil, fmt.Errorf("class with ID %s not found", request.KelasID)
	}
	newWeekRequest := dto.WeekRequest{
		WeekNumber:    request.WeekNumber,
		Kelas_idKelas: request.KelasID,
	}
	newWeek, err := service.weekRepo.CreateWeeklySection(ctx, nil, newWeekRequest)
	if err != nil {
		return nil, err
	}
	newItem := &entities.ItemPembelajaran{
		WeekID: newWeek.ID,
		HeadingPertemuan: request.HeadingPertemuan,
		BodyPertemuan: request.BodyPertemuan,
		UrlVideo: request.UrlVideo,
		FileName: request.FileName,
		FileLink: request.FileLink,
	}
	newItem, err = service.weekRepo.CreateItemPembelajaran(ctx, nil, newItem)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

// Update WeeklySection is not implemented in the original code, so we will not implement it here.
func (service *weekService) UpdateWeeklySection(ctx context.Context,classId uuid.UUID, week_number int, req dto.UpdateItemPembelajaranRequest) (*entities.ItemPembelajaran, error) {
	class, err := service.kelasRepo.GetById(ctx, nil, classId)
	if err != nil {
		return nil, fmt.Errorf("class with ID %s not found", classId)
	}
	if class.ID == uuid.Nil {
		return nil, fmt.Errorf("class with ID %s not found", classId)
	}
	
	item := &entities.ItemPembelajaran{
		WeekID: week_number,
		HeadingPertemuan: req.HeadingPertemuan,
		BodyPertemuan: req.BodyPertemuan,
		UrlVideo: req.UrlVideo,
		FileName: req.FileName,
		FileLink: req.FileLink,
	}
	
	item, err = service.weekRepo.UpdateItemPembelajaran(ctx, nil, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
	

func (service *weekService) DeleteWeeklySection(ctx context.Context, weekID int) error {
	err := service.weekRepo.DeleteWeeklySection(ctx, nil, weekID)
	if err != nil {
		return err
	}
	return nil
}