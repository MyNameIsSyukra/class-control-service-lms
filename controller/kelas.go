package controller

import (
	entities "LMSGo/entity"
	kelas "LMSGo/service"
)

type (
	KelasController interface {
	GetAll() ([]*entities.Kelas, error)
	GetById(id string)(*entities.Kelas, error)
	Create(kelas *entities.Kelas) error	
	Update(id string, kelas *entities.Kelas) error
	Delete(id string) error
}
	kelasController struct {
		kelasService kelas.KelasService 
	}
)

func NewKelasController(kelasService kelas.KelasService) KelasController {
	return &kelasController{kelasService}
}

func (service *kelasController) Create(kelas *entities.Kelas) error {
	return service.kelasService.Create(kelas)
}

func (service *kelasController) GetAll() ([]*entities.Kelas, error) {
	return service.kelasService.GetAll()
}

func (service *kelasController) GetById(id string) (*entities.Kelas, error) {
	return service.kelasService.GetById(id)
}

func (service *kelasController) Update(id string, kelas *entities.Kelas) error {
	return service.kelasService.Update(id, kelas)
}

func (service *kelasController) Delete(id string) error {
	return service.kelasService.Delete(id)
}
