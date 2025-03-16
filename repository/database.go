package database

import (
	entities "LMSGo/entity"
)

type ClassDB interface {
	GetAll() ([]*entities.Kelas, error)
	GetById(id string)(*entities.Kelas, error)
	Create(class *entities.Kelas) error
}