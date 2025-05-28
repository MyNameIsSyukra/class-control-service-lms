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
			ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), // UUID dummy, ganti dengan UUID asli
			Name:        "Mathematics 101",
			Tag:         "Class A",
			Description: "Master the fundamentals of mathematics",
			Teacher:     "John Doe",
			TeacherID:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		},
		{
			ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), // UUID dummy, ganti dengan UUID asli
			Name:        "Introduction to Physics",
			Tag:         "Class B",
			Description: "Learn the basic principles of physics",
			Teacher:     "Jane Smith",
			TeacherID:   uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		},
		{
			ID: uuid.MustParse("33333333-3333-3333-3333-333333333333"), // UUID dummy, ganti dengan UUID asli
			Name:        "Computer Science Basics",
			Tag:         "Class C",
			Description: "Enhance your programming skills through this class",
			Teacher:     "Alice Johnson",
			TeacherID:   uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		},
		{
			ID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), // UUID dummy, ganti dengan UUID asli
			Name:        "English Literature",
			Tag:         "Class D",
			Description: "Explore classic and modern literature",
			Teacher:     "Bob Williams",
			TeacherID:   uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		},
		{
			ID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), // UUID dummy, ganti dengan UUID asli
			Name:        "History of Art",
			Tag:         "Class E",
			Description: "A journey through the history of art and its movements",
			Teacher:     "Carla White",
			TeacherID:   uuid.MustParse("55555555-5555-5555-5555-555555555555"),
		},
		{
			ID: uuid.MustParse("66666666-6666-6666-6666-666666666666"), // UUID dummy, ganti dengan UUID asli
			Name:        "Biology 101",
			Tag:         "Class F",
			Description: "Understand the basics of biology and life sciences",
			Teacher:     "David Brown",
			TeacherID:   uuid.MustParse("66666666-6666-6666-6666-666666666666"),
		},
		{
			ID: uuid.MustParse("77777777-7777-7777-7777-777777777777"), // UUID dummy, ganti dengan UUID asli
			Name:        "Chemistry Fundamentals",
			Tag:         "Class G",
			Description: "Dive into the world of chemistry and its applications",
			Teacher:     "Emily Green",
			TeacherID:   uuid.MustParse("77777777-7777-7777-7777-777777777777"),
		},
		{
			ID: uuid.MustParse("88888888-8888-8888-8888-888888888888"), // UUID dummy, ganti dengan UUID asli
			Name:        "Introduction to Psychology",
			Tag:         "Class H",
			Description: "Discover the basics of human behavior and mind",
			Teacher:     "Frank Black",
			TeacherID:   uuid.MustParse("88888888-8888-8888-8888-888888888888"),
		},
		{
			ID: uuid.MustParse("99999999-9999-9999-9999-999999999999"), // UUID dummy, ganti dengan UUID asli
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
		{
			Username:      "Pak Joko",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd"),
			Kelas_kelasID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		},
		{
			Username:      "Andi",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee"),
			Kelas_kelasID: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		},
		{
			Username:      "Ibu Sari",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff"),
			Kelas_kelasID: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		},
		{
			Username:      "Rudi",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("11111111-2222-3333-4444-555555555555"),
			Kelas_kelasID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		},
		{
			Username:      "Ibu Tini",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("22222222-3333-4444-5555-666666666666"),
			Kelas_kelasID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		},
		{
			Username:      "Dewi",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("33333333-4444-5555-6666-777777777777"),
			Kelas_kelasID: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		},
		{
			Username:      "Pak Agus",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("44444444-5555-6666-7777-888888888888"),
			Kelas_kelasID: uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		},
		{
			Username:      "Lina",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("55555555-6666-7777-8888-999999999999"),
			Kelas_kelasID: uuid.MustParse("55555555-5555-5555-5555-555555555555"),
		},
		{
			Username:      "Pak Budi",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("66666666-7777-8888-9999-000000000000"),
			Kelas_kelasID: uuid.MustParse("66666666-6666-6666-6666-666666666666"),
		},
		{
			Username:      "Rina",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("77777777-8888-9999-0000-111111111111"),
			Kelas_kelasID: uuid.MustParse("77777777-7777-7777-7777-777777777777"),
		},
		{
			Username:      "Pak Jaya",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("88888888-9999-0000-1111-222222222222"),
			Kelas_kelasID: uuid.MustParse("77777777-7777-7777-7777-777777777777"),
		},
		{
			Username:      "Tina",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("99999999-0000-1111-2222-333333333333"),
			Kelas_kelasID: uuid.MustParse("88888888-8888-8888-8888-888888888888"),
		},
		{
			Username:      "Pak Hasan",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("00000000-1111-2222-3333-444444444444"),
			Kelas_kelasID: uuid.MustParse("88888888-8888-8888-8888-888888888888"),
		},
		{
			Username:      "Sari",
			Role:          entities.MemberRoleStudent,
			User_userID:   uuid.MustParse("11111111-2222-3333-4444-555555555555"),
			Kelas_kelasID: uuid.MustParse("99999999-9999-9999-9999-999999999999"),
		},
		{
			Username:      "Pak Danu",
			Role:          entities.MemberRoleTeacher,
			User_userID:   uuid.MustParse("22222222-3333-4444-5555-666666666666"),
			Kelas_kelasID: uuid.MustParse("99999999-9999-9999-9999-999999999999"),
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
	for i := 1; i <= 9; i++ {
		classID := uuid.MustParse(fmt.Sprintf("%d%d%d%d%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d%d%d%d%d%d%d%d%d",i, i, i, i, i, i, i, i,i, i, i, i,i, i, i, i,i, i, i, i,i, i, i, i, i, i, i, i, i, i, i, i)) // Ganti dengan UUID kelas yang sesuai
		weeks := []entities.Week{
			{
				WeekNumber:    1,
				Kelas_idKelas: classID,
			},
			{
				WeekNumber:    2,
				Kelas_idKelas: classID,
			},
		}
		for _, week := range weeks {
			if err := db.Create(&week).Error; err != nil {
				return err
			}
			// Seed item pembelajaran untuk setiap minggu
			itemPembelajaran := entities.ItemPembelajaran{
				WeekID: week.ID,
				HeadingPertemuan: fmt.Sprintf("Pertemuan %d", week.WeekNumber),
				BodyPertemuan: fmt.Sprintf("Materi untuk pertemuan %d", week.WeekNumber),
				UrlVideo: fmt.Sprintf("https://example.com/video_pertemuan_%d.mp4", week.WeekNumber),
				FileName: fmt.Sprintf("file_pertemuan_%d.pdf", week.WeekNumber),
				FileLink: fmt.Sprintf("https://example.com/file_pertemuan_%d.pdf", week.WeekNumber),
			}
			if err := db.Create(&itemPembelajaran).Error; err != nil {
				return err
			}
			// Seed assignment untuk setiap minggu
			assignment := entities.Assignment{
				Title: fmt.Sprintf("Tugas Minggu %d", week.WeekNumber),
				Description: fmt.Sprintf("Deskripsi tugas untuk minggu %d", week.WeekNumber),
				Deadline: time.Now().AddDate(0, 0, 7), // Deadline 7 hari dari sekarang
				FileName: fmt.Sprintf("tugas_minggu_%d.pdf", week.WeekNumber),
				FileLink: fmt.Sprintf("https://example.com/tugas_minggu_%d.pdf", week.WeekNumber),
				WeekID: week.ID,
			}
			if err := db.Create(&assignment).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
