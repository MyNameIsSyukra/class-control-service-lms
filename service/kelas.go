package kelas

import (
	entities "LMSGo/entity"
	database "LMSGo/repository"
)

type classUseCase struct {
	DB database.ClassDB
}

func NewClassUseCase(DB database.ClassDB) *classUseCase {
	return &classUseCase{
		DB }
}

func (uc *classUseCase) GetAll() ([]*entities.Kelas, error) {
	return uc.DB.GetAll()
}

func (uc *classUseCase) GetById(id string) (*entities.Kelas, error) {
	return uc.DB.GetById(id)
}

func (uc *classUseCase) Create(kelas *entities.Kelas) error {
	return uc.DB.Create(kelas)
}



