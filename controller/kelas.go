package controller

import (
	entities "LMSGo/entity"
)

type KelasController interface {
	GetAll() ([]*entities.Kelas, error)
	GetById(id string)(*entities.Kelas, error)
	Create(kelas *entities.Kelas) error	
}