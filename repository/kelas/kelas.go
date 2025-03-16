package kelas

import (
	entities "LMSGo/entity"
)

type KelasDB struct {
	kelas []*entities.Kelas
}

func NewKelasDB() *KelasDB {
	return &KelasDB{
		kelas: []*entities.Kelas{
			{
				ID: 		"1",
				Name: 		"Kelas 1",
				Description: 	"Kelas 1",
				Teacher: 	"Teacher 1",
				TeacherID: 	1,
			},
			{
				ID: 		"2",
				Name: 		"Kelas 2",
				Description: 	"Kelas 2",
				Teacher: 	"Teacher 2",
				TeacherID: 	2,
			},
		},
	}
}

func (repo *KelasDB) GetAll() ([]*entities.Kelas, error) {
	return repo.kelas, nil
}


func (repo *KelasDB) GetById(id string) (*entities.Kelas, error) {
	for _, k := range repo.kelas {
		if k.ID == id {
			return k, nil
		}
	}
	return nil, nil
}

func (repo *KelasDB) Create(kelas *entities.Kelas) error {
	repo.kelas = append(repo.kelas, kelas)
	return nil
}