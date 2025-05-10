package migration

import (
	"fmt"
	"time"

	database "LMSGo/config"
	entities "LMSGo/entity"

	"github.com/google/uuid"
)

func Seeder()error{
	// Inisialisasi koneksi ke database
	db := database.SetUpDatabaseConnection()
	db.AutoMigrate(&entities.Kelas{}, &entities.Member{}, &entities.ItemPembelajaran{})
	// UUID tetap
	for i := 1; i <= 5; i++ {
		kelasUUID, _ := uuid.Parse(fmt.Sprintf("00000000-0000-0000-0000-00000000000%d", i))
		teacherUUID, _ := uuid.Parse(fmt.Sprintf("00000000-0000-0000-0000-0000000001%d", i))
		userUUID, _ := uuid.Parse(fmt.Sprintf("00000000-0000-0000-0000-0000000002%d", i))

	// Seeder: Kelas
	kelas := entities.Kelas{
		ID:          kelasUUID,
		Name:        "Struktur Data",
		Tag:         "SD123",
		Description: "Belajar dasar struktur data seperti array, linked list, dan tree",
		Teacher:     "Dosen Andi",
		TeacherID:   teacherUUID,
	}
	db.Create(&kelas)

	// Seeder: Member
	member := entities.Member{
		Username:      fmt.Sprintf("mahasiswa%d",i),
		Role:          entities.MemberRoleStudent,
		User_userID:   userUUID,
		Kelas_kelasID: kelasUUID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(&member)

	// Seeder: ItemPembelajaran
	item := entities.ItemPembelajaran{
		HeadingPertemuan:   fmt.Sprintf("Pertemuan %d: Array", i),
		BodyPertemuan:      "Penjelasan tentang array dan implementasinya dalam bahasa C",
		// FileName:           "array.pdf",
		// FilePath:           "/files/array.pdf",
		UrlVideo:           "https://youtu.be/strukturdata1",
		// Kelas_idKelas:      kelasUUID, 
	}
	db.Create(&item)
}

	fmt.Println("Seeder selesai dijalankan.")
	return nil
}