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
	db.AutoMigrate(&entities.Kelas{}, &entities.Member{}, &entities.ItemPembelajaran{}, &entities.Assignment{}, &entities.AssignmentSubmission{}, &entities.Week{})
	// UUID tetap
	
	// Generate static data for consistency across services
	fmt.Println("Generating static shared data...")
	sharedClasses, sharedUsers := GenerateStaticData()

	// Seed Class Control service
	SeedClassControlData(db, sharedClasses, sharedUsers)

	// Print static UUIDs for reference
	PrintStaticUUIDs()

	fmt.Println("\n========== CLASS CONTROL SEEDING COMPLETED ==========")
	fmt.Println("Data seeded successfully with static UUID references!")
	fmt.Println("Use the same static UUIDs in Assessment service seeder for consistency.")
	return nil
}
var (
	// Class IDs (will be used as ClassID in Assessment service)
	ClassWebProgID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")
	ClassDatabaseID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440002")
	ClassAlgorithmID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440003")

	// Teacher IDs
	TeacherAhmadID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440101")
	TeacherSitiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440102")
	TeacherBudiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440103")

	// Student IDs (will be used as UserID in Assessment service)
	StudentAliceID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440201")
	StudentBobID      = uuid.MustParse("550e8400-e29b-41d4-a716-446655440202")
	StudentCharlieID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440203")
	StudentDianaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440204")
	StudentEdwardID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440205")
	StudentFionaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440206")
	StudentGeorgeID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440207")
	StudentHannahID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440208")
	StudentIvanID     = uuid.MustParse("550e8400-e29b-41d4-a716-446655440209")
	StudentJuliaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440210")
	StudentKevinID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440211")
	StudentLindaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440212")
	StudentMichaelID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440213")
	StudentNancyID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440214")
	StudentOscarID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440215")

	// Random YouTube URLs for video content
	randomYouTubeURLs = []string{
		"https://youtu.be/k2_2H3qT9q0?feature=shared",
		"https://youtu.be/djF9-SHIgQg?feature=shared",
		"https://youtu.be/B2IldXHBDA0?feature=shared",
	}
)

// ========== SHARED DATA STRUCTURE ==========
type SharedClassData struct {
	ClassID     uuid.UUID
	Name        string
	Tag         string
	Description string
	Teacher     string
	TeacherID   uuid.UUID
}

type SharedUserData struct {
	UserID   uuid.UUID
	Username string
	Role     entities.MemberRole
	ClassID  uuid.UUID
}

// Helper function to get random YouTube URL
func getRandomYouTubeURL(weekNum int) string {
	// Use week number to cycle through the URLs
	index := (weekNum - 1) % len(randomYouTubeURLs)
	return randomYouTubeURLs[index]
}

func GenerateStaticData() ([]SharedClassData, []SharedUserData) {
	// Generate consistent class data with static UUIDs
	classes := []SharedClassData{
   {
   	ClassID:     ClassWebProgID,
   	Name:        "English Grammar",
   	Tag:         "EG",
   	Description: "Mata kuliah tata bahasa Inggris untuk pemahaman struktur kalimat",
   	Teacher:     "Dr. Ahmad Santoso",
   	TeacherID:   TeacherAhmadID,
   },
   {
   	ClassID:     ClassDatabaseID,
   	Name:        "English Conversation",
   	Tag:         "EC",
   	Description: "Mata kuliah percakapan bahasa Inggris untuk komunikasi sehari-hari",
   	Teacher:     "Prof. Siti Nurhaliza",
   	TeacherID:   TeacherSitiID,
   },
   {
   	ClassID:     ClassAlgorithmID,
   	Name:        "English Literature",
   	Tag:         "EL",
   	Description: "Mata kuliah sastra Inggris untuk pemahaman karya sastra klasik dan modern",
   	Teacher:     "Dr. Budi Raharjo",
   	TeacherID:   TeacherBudiID,
   },
}
	// Generate user data with static UUIDs
	var users []SharedUserData

	// Add teachers
	users = append(users, []SharedUserData{
		{UserID: TeacherAhmadID, Username: "Dr. Ahmad Santoso", Role: entities.MemberRoleTeacher, ClassID: ClassWebProgID},
		{UserID: TeacherSitiID, Username: "Prof. Siti Nurhaliza", Role: entities.MemberRoleTeacher, ClassID: ClassDatabaseID},
		{UserID: TeacherBudiID, Username: "Dr. Budi Raharjo", Role: entities.MemberRoleTeacher, ClassID: ClassAlgorithmID},
	}...)

	// Add students with static UUIDs
	studentData := []struct {
		ID       uuid.UUID
		Username string
		ClassID  uuid.UUID
	}{
		// Web Programming Students
		{StudentAliceID, "Alice Johnson", ClassWebProgID},
		{StudentBobID, "Bob Smith", ClassWebProgID},
		{StudentCharlieID, "Charlie Brown", ClassWebProgID},
		{StudentDianaID, "Diana Prince", ClassWebProgID},
		{StudentEdwardID, "Edward Norton", ClassWebProgID},

		// Database Students
		{StudentFionaID, "Fiona Green", ClassDatabaseID},
		{StudentGeorgeID, "George Wilson", ClassDatabaseID},
		{StudentHannahID, "Hannah Davis", ClassDatabaseID},
		{StudentIvanID, "Ivan Petrov", ClassDatabaseID},
		{StudentJuliaID, "Julia Roberts", ClassDatabaseID},

		// Algorithm Students
		{StudentKevinID, "Kevin Hart", ClassAlgorithmID},
		{StudentLindaID, "Linda Carter", ClassAlgorithmID},
		{StudentMichaelID, "Michael Jordan", ClassAlgorithmID},
		{StudentNancyID, "Nancy Drew", ClassAlgorithmID},
		{StudentOscarID, "Oscar Wilde", ClassAlgorithmID},
	}

	for _, student := range studentData {
		users = append(users, SharedUserData{
			UserID:   student.ID,
			Username: student.Username,
			Role:     entities.MemberRoleStudent,
			ClassID:  student.ClassID,
		})
	}

	return classes, users
}

// ========== SEEDER FUNCTIONS ==========
func SeedClassControlData(db *gorm.DB, sharedClasses []SharedClassData, sharedUsers []SharedUserData) {
	fmt.Println("Seeding Class Control data...")

	// Seed Kelas using shared data
	for _, classData := range sharedClasses {
		class := entities.Kelas{
			ID:          classData.ClassID,
			Name:        classData.Name,
			Tag:         classData.Tag,
			Description: classData.Description,
			Teacher:     classData.Teacher,
			TeacherID:   classData.TeacherID,
		}
		db.Create(&class)
	}

	// Seed Members using shared data
	for _, userData := range sharedUsers {
		member := entities.Member{
			Username:      userData.Username,
			Role:          userData.Role,
			User_userID:   userData.UserID,
			Kelas_kelasID: userData.ClassID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		db.Create(&member)
	}

	// Seed Weeks and related data
	for _, classData := range sharedClasses {
		// 68525b244df1e68bde926058 tugas 1 kelas 1
		// 68525b614df1e68bde92605b tugas 2 kelas 1
		// 68525b7f4df1e68bde926060 tugas 1 kelas 2
		// 68525b8c4df1e68bde926063 tugas 2 kelas 2
		// 68525ba44df1e68bde926068 tugas 1 kelas 3
		// 68525bb34df1e68bde92606b tugas 2 kelas 3
		// 68525bc74df1e68bde926070 tugas 1 kelas 4
		// 68525bd34df1e68bde926073 tugas 2 kelas 4

		// file id for submission
		// 68429cdb04d68646d09139b8 submission file id
		// 68433640d85fc55aafc70d2d
		var assignmentFileID string
		for weekNum := 1; weekNum <= 4; weekNum++ {
		week := entities.Week{
			WeekNumber:    weekNum,
			Kelas_idKelas: classData.ClassID,
		}
		db.Create(&week)

		// Seed ItemPembelajaran with random YouTube URLs
		itemPembelajaran := entities.ItemPembelajaran{
		    WeekID:           week.ID,
		    HeadingPertemuan: fmt.Sprintf("Pertemuan %d - %s", weekNum, classData.Name),
		    BodyPertemuan:    fmt.Sprintf("Materi pembelajaran minggu ke-%d untuk mata kuliah %s", weekNum, classData.Name),
		    UrlVideo:         getRandomYouTubeURL(weekNum), // Using random YouTube URLs
		    FileName:         fmt.Sprintf("materi_%s_week_%d.pdf", classData.Tag, weekNum),
		    FileId:           "6852508f4df1e68bde926047", // Google Drive file ID format
		}
		db.Create(&itemPembelajaran)

		// Seed Assignment (every 2 weeks)
		if weekNum%2 == 0 {
			// Determine assignment FileID based on class tag and week
			switch classData.Tag {
			case "PWEB": // Kelas 1
				if weekNum == 2 {
					assignmentFileID = "68525b244df1e68bde926058" // tugas 1 kelas 1
				} else if weekNum == 4 {
					assignmentFileID = "68525b614df1e68bde92605b" // tugas 2 kelas 1
				}
			case "BD": // Kelas 2
				if weekNum == 2 {
					assignmentFileID = "68525b7f4df1e68bde926060" // tugas 1 kelas 2
				} else if weekNum == 4 {
					assignmentFileID = "68525b8c4df1e68bde926063" // tugas 2 kelas 2
				}
			case "ASD": // Kelas 3
				if weekNum == 2 {
					assignmentFileID = "68525ba44df1e68bde926068" // tugas 1 kelas 3
				} else if weekNum == 4 {
					assignmentFileID = "68525bb34df1e68bde92606b" // tugas 2 kelas 3
				}
			}

		    assignment := entities.Assignment{
		        Title:       fmt.Sprintf("Tugas %s - Minggu %d", classData.Name, weekNum),
		        Description: fmt.Sprintf("Tugas praktikum untuk minggu ke-%d mata kuliah %s", weekNum, classData.Name),
		        Deadline:    time.Now().AddDate(0, 0, 7),
		        FileName:    fmt.Sprintf("tugas_%s_week_%d.pdf", classData.Tag, weekNum),
		        FileId:      assignmentFileID, // Using specific FileID from comment
		        WeekID:      week.ID,
		    }
		    db.Create(&assignment)

			// Seed AssignmentSubmissions using shared user data
			studentsInClass := make([]SharedUserData, 0)
			for _, user := range sharedUsers {
				if user.ClassID == classData.ClassID && user.Role == entities.MemberRoleStudent {
					studentsInClass = append(studentsInClass, user)
				}
			}

			// Create submissions for first 3 students in class
			for j := 0; j < 3 && j < len(studentsInClass); j++ {
				submission := entities.AssignmentSubmission{
					AssignmentID: int(assignment.ID),
					UserID:       studentsInClass[j].UserID,
					IDFile:       "68429cdb04d68646d09139b8",
					FileName:     fmt.Sprintf("submission_%s.pdf", studentsInClass[j].Username),
					Score:        85 + j*5,
					Status:       entities.StatusSubmitted,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}
				db.Create(&submission)
			}
			}
		}
	}

	fmt.Println("Class Control data seeded successfully!")
}

// ========== SUMMARY ==========
func PrintStaticUUIDs() {
	fmt.Println("\n========== STATIC UUID REFERENCES ==========")
	fmt.Println("CLASS IDs (untuk Assessment.ClassID):")
	fmt.Printf("  - Pemrograman Web: %s\n", ClassWebProgID)
	fmt.Printf("  - Basis Data: %s\n", ClassDatabaseID)
	fmt.Printf("  - Algoritma: %s\n", ClassAlgorithmID)

	fmt.Println("\nTEACHER IDs:")
	fmt.Printf("  - Dr. Ahmad Santoso: %s\n", TeacherAhmadID)
	fmt.Printf("  - Prof. Siti Nurhaliza: %s\n", TeacherSitiID)
	fmt.Printf("  - Dr. Budi Raharjo: %s\n", TeacherBudiID)

	fmt.Println("\nSTUDENT IDs (untuk Assessment.Submission.UserID):")
	fmt.Printf("  - Alice Johnson: %s\n", StudentAliceID)
	fmt.Printf("  - Bob Smith: %s\n", StudentBobID)
	fmt.Printf("  - Charlie Brown: %s\n", StudentCharlieID)
	fmt.Printf("  - Diana Prince: %s\n", StudentDianaID)
	fmt.Printf("  - Edward Norton: %s\n", StudentEdwardID)
	fmt.Printf("  - Fiona Green: %s\n", StudentFionaID)
	fmt.Printf("  - George Wilson: %s\n", StudentGeorgeID)
	fmt.Printf("  - Hannah Davis: %s\n", StudentHannahID)
	fmt.Printf("  - Ivan Petrov: %s\n", StudentIvanID)
	fmt.Printf("  - Julia Roberts: %s\n", StudentJuliaID)
	fmt.Printf("  - Kevin Hart: %s\n", StudentKevinID)
	fmt.Printf("  - Linda Carter: %s\n", StudentLindaID)
	fmt.Printf("  - Michael Jordan: %s\n", StudentMichaelID)
	fmt.Printf("  - Nancy Drew: %s\n", StudentNancyID)
	fmt.Printf("  - Oscar Wilde: %s\n", StudentOscarID)

	fmt.Println("\nRANDOM YOUTUBE URLs USED:")
	for i, url := range randomYouTubeURLs {
		fmt.Printf("  - URL %d: %s\n", i+1, url)
	}
}