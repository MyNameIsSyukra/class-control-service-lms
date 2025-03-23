package kelas

import (
	entities "LMSGo/entity"
	database "LMSGo/repository"
)

type (
	KelasService interface {
		GetAll() ([]*entities.Kelas, error)
		GetById(id string) (*entities.Kelas, error)
		Create(kelas *entities.Kelas) error
		Update(id string, kelas *entities.Kelas) error
	}

	kelasService struct {
	kelasRepo database.Repository
}
)

func NewKelasService(kelasRepo database.Repository) KelasService {
	return &kelasService{kelasRepo}
}

func (service *kelasService) GetAll() ([]*entities.Kelas, error) {
	return service.kelasRepo.GetAll()
}

func (service *kelasService) GetById(id string) (*entities.Kelas, error) {
	return service.kelasRepo.GetById(id)
}

func (service *kelasService) Create(kelas *entities.Kelas) error {
	return service.kelasRepo.Create(kelas)
}

func (service *kelasService) Update(id string, kelas *entities.Kelas) error {
	return service.kelasRepo.Update(id, kelas)
}




