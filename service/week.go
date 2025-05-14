package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"context"

	"github.com/google/uuid"
)

type (
	WeekService interface {
		GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) ([]*entities.Week, error)
		GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponseByID, error)
		CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest) (*entities.ItemPembelajaran, error)
		// CreateWeeklySection(ctx context.Context, weekReq dto.WeekRequest) (*entities.Week, error)
		// DeleteWeeklySection(ctx context.Context, weekID int) error
	}
	weekService struct {
		weekRepo database.WeekRepository
	}
)

func NewWeekService(weekRepo database.WeekRepository) WeekService {
	return &weekService{weekRepo}
}

func (service *weekService) GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) ([]*entities.Week, error) {
	weeks, err := service.weekRepo.GetAllWeekByClassID(ctx, nil, classID)
	if err != nil {
		return nil, err
	}
	return weeks, nil
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
		Assignment:        &week.Assignment,
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
