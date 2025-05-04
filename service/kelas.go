package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"context"
)

type (
	KelasService interface {
		Create(ctx context.Context,kelas *dto.CreateKelasRequest) (*entities.Kelas, error)
		GetAllKelasWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.KelasPaginationResponse, error)
		GetById(ctx context.Context,id string) (*entities.Kelas, error)
		Update(ctx context.Context,id string, kelas *dto.CreateKelasUpdateRequest) (*entities.Kelas,error)
		Delete(ctx context.Context,id string) error
	}

	kelasService struct {kelasRepo database.KelasRepository
}
)

func NewKelasService(kelasRepo database.KelasRepository) KelasService {
	return &kelasService{kelasRepo}
}

func (service *kelasService) GetAllKelasWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.KelasPaginationResponse, error) {
	dataWithPaginate,err := service.kelasRepo.GetAll(ctx, nil, req)
	if err != nil {
		return dto.KelasPaginationResponse{}, err
	}
	var datas []dto.KelasResponse
	for _, kelas := range dataWithPaginate.Kelas {
		datas = append(datas, dto.KelasResponse{
			ID:          kelas.ID,
			Name:        kelas.Name,
			Description: kelas.Description,
			Teacher:     kelas.Teacher,
			TeacherID:   kelas.TeacherID,
		})
	}
	return dto.KelasPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.PaginationResponse.Page,
			PerPage: dataWithPaginate.PaginationResponse.PerPage,
			MaxPage: dataWithPaginate.PaginationResponse.MaxPage,
			Count:   dataWithPaginate.PaginationResponse.Count,
		},
	}, nil
}
















































func (service *kelasService) GetById(ctx context.Context, id string) (*entities.Kelas, error) {
	class, err := service.kelasRepo.GetById(ctx,nil, id)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (service *kelasService) Create(ctx context.Context,kelas *dto.CreateKelasRequest)(*entities.Kelas, error) {
	kelasEntity := &entities.Kelas{
		Name:        kelas.Name,
		Tag:         kelas.Tag,
		Description: kelas.Description,
		Teacher:     kelas.Teacher,
		TeacherID:   kelas.TeacherID,
	}
	class ,err := service.kelasRepo.Create(ctx, nil, kelasEntity);
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (service *kelasService) Update(ctx context.Context,id string, kelas *dto.CreateKelasUpdateRequest) (*entities.Kelas,error) {
	clas,err := service.kelasRepo.GetById(ctx, nil, id)
	if clas == nil {
		return nil,err
	}
	classEntity := &entities.Kelas{
		Name:        kelas.Name,
		Description: kelas.Description,
		Teacher:     kelas.Teacher,
		TeacherID:   kelas.TeacherID,
	}

	data, err := service.kelasRepo.Update(ctx, nil, id, classEntity); 
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *kelasService) Delete(ctx context.Context,id string) error {
	class, err := service.kelasRepo.GetById(ctx, nil, id)
	if class == nil {
		return err
	}
	err = service.kelasRepo.Delete(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}


