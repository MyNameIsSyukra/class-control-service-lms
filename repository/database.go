package database

import (
	entities "LMSGo/entity"
)

type Repository interface {
	GetAll() ([]*entities.Kelas, error)
	GetById(id string)(*entities.Kelas, error)
	Create(class *entities.Kelas) error
	Update(id string, class *entities.Kelas) error
}