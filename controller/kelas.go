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
}
	kelasController struct {
		kelasService kelas.KelasService 
	}
)

func NewKelasController(kelasService kelas.KelasService) KelasController {
	return &kelasController{kelasService}
}

func (uc *kelasController) GetAll() ([]*entities.Kelas, error) {
	return uc.kelasService.GetAll()
}

func (uc *kelasController) GetById(id string) (*entities.Kelas, error) {
	return uc.kelasService.GetById(id)
}

func (uc *kelasController) Create(kelas *entities.Kelas) error {
	return uc.kelasService.Create(kelas)
}

func (uc *kelasController) Update(id string, kelas *entities.Kelas) error {
	return uc.kelasService.Update(id, kelas)
}