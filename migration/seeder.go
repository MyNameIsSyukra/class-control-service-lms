package migration

import (
	"fmt"
	"time"

	database "LMSGo/config"
	entities "LMSGo/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeder()error{
	// Inisialisasi koneksi ke database
	db := database.SetUpDatabaseConnection()
	db.AutoMigrate(&entities.Kelas{}, &entities.Member{}, &entities.ItemPembelajaran{})
	// UUID tetap
	
	// Seed data kelas
	if err := SeedKelas(db); err != nil {
		fmt.Println("Gagal melakukan seeding kelas:", err)
	}
	// Seed data members
	if err := SeedMembers(db); err != nil {
		fmt.Println("Gagal melakukan seeding members:", err)
	}
	// Seed data week content
	if err := SeedWeekContent(db); err != nil {
		fmt.Println("Gagal melakukan seeding week content:", err)
	}

	fmt.Println("Seeder selesai dijalankan.")
	return nil
}

// create seeder class form data below	
func SeedKelas(db *gorm.DB) error {
	kelasData := []entities.Kelas{
		{
			Name:        "Mathematics 101",
			Tag:         "Class A",
			Description: "Master the fundamentals of mathematics",
			Teacher:     "John Doe",
			TeacherID:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		},
		{
			Name:        "Introduction to Physics",
			Tag:         "Class B",
			Description: "Learn the basic principles of physics",
			Teacher:     "Jane Smith",
			TeacherID:   uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		},
		{
			Name:        "Computer Science Basics",
			Tag:         "Class C",
			Description: "Enhance your programming skills through this class",
			Teacher:     "Alice Johnson",
			TeacherID:   uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		},
		{
			Name:        "English Literature",
			Tag:         "Class D",
			Description: "Explore classic and modern literature",
			Teacher:     "Bob Williams",
			TeacherID:   uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		},
		{
			Name:        "History of Art",
			Tag:         "Class E",
			Description: "A journey through the history of art and its movements",
			Teacher:     "Carla White",
			TeacherID:   uuid.MustParse("55555555-5555-5555-5555-555555555555"),
		},
		{
			Name:        "Biology 101",
			Tag:         "Class F",
			Description: "Understand the basics of biology and life sciences",
			Teacher:     "David Brown",
			TeacherID:   uuid.MustParse("66666666-6666-6666-6666-666666666666"),
		},
		{
			Name:        "Chemistry Fundamentals",
			Tag:         "Class G",
			Description: "Dive into the world of chemistry and its applications",
			Teacher:     "Emily Green",
			TeacherID:   uuid.MustParse("77777777-7777-7777-7777-777777777777"),
		},
		{
			Name:        "Introduction to Psychology",
			Tag:         "Class H",
			Description: "Discover the basics of human behavior and mind",
			Teacher:     "Frank Black",
			TeacherID:   uuid.MustParse("88888888-8888-8888-8888-888888888888"),
		},
		{
			Name:        "Economics 101",
			Tag:         "Class I",
			Description: "Learn the principles of economics and its impact on society",
			Teacher:     "Grace Blue",
			TeacherID:   uuid.MustParse("99999999-9999-9999-9999-999999999999"),
		},
	}

	for _, kelas := range kelasData {
		if err := db.Create(&kelas).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedMembers(db *gorm.DB) error {
	members := []entities.Member{
		{
			Username:      "Bambang",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), // ganti dengan UUID user asli
			Kelas_kelasID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), // ganti dengan UUID kelas asli
		},
		{
			Username:      "Bu Nanik",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			Kelas_kelasID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		},
		{
			Username:      "Siti",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
			Kelas_kelasID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		},
	}

	for _, member := range members {
		if err := db.Create(&member).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedWeekContent(db *gorm.DB) error {
	// UUID dummy untuk kelas (ubah dengan UUID kelas asli dari DB kamu)
	classID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// Buat data Week
	weeks := []entities.Week{
		{
			ID:            1,
			WeekNumber:    1,
			Kelas_idKelas: classID,
		},
		{
			ID:            2,
			WeekNumber:    2,
			Kelas_idKelas: classID,
		},
	}

	for _, week := range weeks {
		if err := db.Create(&week).Error; err != nil {
			return err
		}
	}

	// Buat data ItemPembelajaran (1:1 dengan Week)
	itemPembelajarans := []entities.ItemPembelajaran{
		{
			WeekID:           1,
			HeadingPertemuan: "Pengantar Matematika",
			BodyPertemuan:    "Materi dasar bilangan, operasi, dan persamaan.",
			UrlVideo:         "https://youtube.com/example-video-1",
			FileName:         "materi_minggu1.pdf",
			FileLink:         "https://drive.google.com/file/d/abc123",
		},
		{
			WeekID:           2,
			HeadingPertemuan: "Aljabar Dasar",
			BodyPertemuan:    "Konsep dasar variabel dan persamaan linear.",
			UrlVideo:         "https://youtube.com/example-video-2",
			FileName:         "materi_minggu2.pdf",
			FileLink:         "https://drive.google.com/file/d/def456",
		},
	}

	for _, item := range itemPembelajarans {
		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}

	// Buat data Assignment (1:1 dengan Week)
	assignments := []entities.Assignment{
		{
			Title:       "Tugas 1 - Pengantar",
			Description: "Kerjakan soal latihan dari PDF yang telah diberikan.",
			Deadline:    time.Now().Add(48 * time.Hour),
			FileName:    "tugas1.pdf",
			FileLink:    "https://drive.google.com/file/d/tugas1",
			WeekID:      1,
		},
		{
			Title:       "Tugas 2 - Aljabar",
			Description: "Buat ringkasan dan kerjakan latihan aljabar.",
			Deadline:    time.Now().Add(96 * time.Hour),
			FileName:    "tugas2.pdf",
			FileLink:    "https://drive.google.com/file/d/tugas2",
			WeekID:      2,
		},
	}

	for _, assignment := range assignments {
		if err := db.Create(&assignment).Error; err != nil {
			return err
		}
	}

	return nil
}
