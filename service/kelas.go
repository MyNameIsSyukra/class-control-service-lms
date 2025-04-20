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
		Delete(id string) error
	}

	kelasService struct {kelasRepo database.KelasRepository
}
)

func NewKelasService(kelasRepo database.KelasRepository) KelasService {
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

func (service *kelasService) Delete(id string) error {
	return service.kelasRepo.Delete(id)
}


