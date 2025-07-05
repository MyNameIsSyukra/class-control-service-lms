// package seed

// import (
// 	"fmt"
// 	"time"

// 	database "LMSGo/config"
// 	entities "LMSGo/entity"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// func Seeder() error {
// 	// Inisialisasi koneksi ke database
// 	db := database.SetUpDatabaseConnection()
// 	db.AutoMigrate(&entities.Kelas{}, &entities.Member{}, &entities.ItemPembelajaran{}, &entities.Assignment{}, &entities.AssignmentSubmission{}, &entities.Week{})

// 	// Generate static data for consistency across services
// 	fmt.Println("Generating static shared data...")
// 	sharedClasses, sharedUsers := GenerateStaticData()

// 	// Seed Class Control service
// 	SeedClassControlData(db, sharedClasses, sharedUsers)

// 	// Print static UUIDs for reference
// 	PrintStaticUUIDs()

// 	fmt.Println("\n========== CLASS CONTROL SEEDING COMPLETED ==========")
// 	fmt.Println("Data seeded successfully with static UUID references!")
// 	fmt.Println("Use the same static UUIDs in Assessment service seeder for consistency.")
// 	return nil
// }

// var (
// 	// Class IDs (will be used as ClassID in Assessment service)
// 	ClassWebProgID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")
// 	ClassDatabaseID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440002")
// 	ClassAlgorithmID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440003")

// 	// Teacher IDs (matching Python seeder)
// 	TeacherAhmadID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440101")
// 	TeacherSitiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440102")
// 	TeacherBudiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440103")

// 	// Student IDs (matching Python seeder - first 15 named students)
// 	StudentAliceID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440201")
// 	StudentBobID      = uuid.MustParse("550e8400-e29b-41d4-a716-446655440202")
// 	StudentCharlieID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440203")
// 	StudentDianaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440204")
// 	StudentEdwardID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440205")
// 	StudentFionaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440206")
// 	StudentGeorgeID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440207")
// 	StudentHannahID   = uuid.MustParse("550e8400-e29b-41d4-a716-446655440208")
// 	StudentIvanID     = uuid.MustParse("550e8400-e29b-41d4-a716-446655440209")
// 	StudentJuliaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440210")
// 	StudentKevinID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440211")
// 	StudentLindaID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440212")
// 	StudentMichaelID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440213")
// 	StudentNancyID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440214")
// 	StudentOscarID    = uuid.MustParse("550e8400-e29b-41d4-a716-446655440215")

// 	// Random YouTube URLs for video content
// 	randomYouTubeURLs = []string{
// 		"https://youtu.be/k2_2H3qT9q0?feature=shared",
// 		"https://youtu.be/djF9-SHIgQg?feature=shared",
// 		"https://youtu.be/B2IldXHBDA0?feature=shared",
// 	}
// )

// // ========== SHARED DATA STRUCTURE ==========
// type SharedClassData struct {
// 	ClassID     uuid.UUID
// 	Name        string
// 	Tag         string
// 	Description string
// 	Teacher     string
// 	TeacherID   uuid.UUID
// }

// type SharedUserData struct {
// 	UserID   uuid.UUID
// 	Username string
// 	Role     entities.MemberRole
// 	ClassID  uuid.UUID
// }

// // Helper function to get random YouTube URL
// func getRandomYouTubeURL(weekNum int) string {
// 	index := (weekNum - 1) % len(randomYouTubeURLs)
// 	return randomYouTubeURLs[index]
// }

// // Generate all student UUIDs (including the 1000 generated students)
// func generateAllStudentUUIDs() []uuid.UUID {
// 	var studentUUIDs []uuid.UUID

// 	// Add the first 15 named students
// 	namedStudents := []uuid.UUID{
// 		StudentAliceID, StudentBobID, StudentCharlieID, StudentDianaID, StudentEdwardID,
// 		StudentFionaID, StudentGeorgeID, StudentHannahID, StudentIvanID, StudentJuliaID,
// 		StudentKevinID, StudentLindaID, StudentMichaelID, StudentNancyID, StudentOscarID,
// 	}
// 	studentUUIDs = append(studentUUIDs, namedStudents...)

// 	// Add the remaining 985 students (216-1215 from Python seeder)
// 	for i := 216; i <= 1215; i++ {
// 		uuidStr := fmt.Sprintf("550e8400-e29b-41d4-a716-44665544%04d", i)
// 		studentID := uuid.MustParse(uuidStr)
// 		studentUUIDs = append(studentUUIDs, studentID)
// 	}

// 	return studentUUIDs
// }

// func GenerateStaticData() ([]SharedClassData, []SharedUserData) {
// 	// Generate consistent class data with static UUIDs
// 	classes := []SharedClassData{
// 		{
// 			ClassID:     ClassWebProgID,
// 			Name:        "English Grammar",
// 			Tag:         "EG",
// 			Description: "Mata kuliah tata bahasa Inggris untuk pemahaman struktur kalimat",
// 			Teacher:     "Ahmad Susanto",
// 			TeacherID:   TeacherAhmadID,
// 		},
// 		{
// 			ClassID:     ClassDatabaseID,
// 			Name:        "English Conversation",
// 			Tag:         "EC",
// 			Description: "Mata kuliah percakapan bahasa Inggris untuk komunikasi sehari-hari",
// 			Teacher:     "Siti Nurhaliza",
// 			TeacherID:   TeacherSitiID,
// 		},
// 		{
// 			ClassID:     ClassAlgorithmID,
// 			Name:        "English Literature",
// 			Tag:         "EL",
// 			Description: "Mata kuliah sastra Inggris untuk pemahaman karya sastra klasik dan modern",
// 			Teacher:     "Budi Santoso",
// 			TeacherID:   TeacherBudiID,
// 		},
// 	}

// 	// Generate user data with static UUIDs
// 	var users []SharedUserData

// 	// Add teachers (matching Python seeder names)
// 	users = append(users, []SharedUserData{
// 		{UserID: TeacherAhmadID, Username: "Ahmad Susanto", Role: entities.MemberRoleTeacher, ClassID: ClassWebProgID},
// 		{UserID: TeacherSitiID, Username: "Siti Nurhaliza", Role: entities.MemberRoleTeacher, ClassID: ClassDatabaseID},
// 		{UserID: TeacherBudiID, Username: "Budi Santoso", Role: entities.MemberRoleTeacher, ClassID: ClassAlgorithmID},
// 	}...)

// 	// Get all student UUIDs
// 	allStudentUUIDs := generateAllStudentUUIDs()

// 	// Named students data (matching Python seeder exactly)
// 	namedStudents := []struct {
// 		ID       uuid.UUID
// 		Username string
// 	}{
// 		{StudentAliceID, "Alice Johnson"},
// 		{StudentBobID, "Bob Smith"},
// 		{StudentCharlieID, "Charlie Brown"},
// 		{StudentDianaID, "Diana Prince"},
// 		{StudentEdwardID, "Edward Norton"},
// 		{StudentFionaID, "Fiona Green"},
// 		{StudentGeorgeID, "George Washington"},
// 		{StudentHannahID, "Hannah Montana"},
// 		{StudentIvanID, "Ivan Petrov"},
// 		{StudentJuliaID, "Julia Roberts"},
// 		{StudentKevinID, "Kevin Hart"},
// 		{StudentLindaID, "Linda Hamilton"},
// 		{StudentMichaelID, "Michael Jordan"},
// 		{StudentNancyID, "Nancy Drew"},
// 		{StudentOscarID, "Oscar Wilde"},
// 	}

// 	// Distribute students across classes more evenly
// 	studentsPerClass := len(allStudentUUIDs) / len(classes)
// 	remainder := len(allStudentUUIDs) % len(classes)

// 	studentIndex := 0
// 	for classIndex, class := range classes {
// 		// Calculate how many students this class should get
// 		numStudentsForClass := studentsPerClass
// 		if classIndex < remainder {
// 			numStudentsForClass++
// 		}

// 		// Add students to this class
// 		for i := 0; i < numStudentsForClass && studentIndex < len(allStudentUUIDs); i++ {
// 			studentID := allStudentUUIDs[studentIndex]
// 			var username string

// 			// Use named student username if it's one of the first 15
// 			if studentIndex < len(namedStudents) {
// 				username = namedStudents[studentIndex].Username
// 			} else {
// 				// Generate username for numbered students (matching Python pattern)
// 				studentNum := studentIndex + 216 - 15 // Adjust for the offset
// 				username = fmt.Sprintf("Student%d Test", studentNum)
// 			}

// 			users = append(users, SharedUserData{
// 				UserID:   studentID,
// 				Username: username,
// 				Role:     entities.MemberRoleStudent,
// 				ClassID:  class.ClassID,
// 			})

// 			studentIndex++
// 		}
// 	}

// 	return classes, users
// }

// // ========== SEEDER FUNCTIONS ==========
// func SeedClassControlData(db *gorm.DB, sharedClasses []SharedClassData, sharedUsers []SharedUserData) {
// 	fmt.Println("Seeding Class Control data...")

// 	// Seed Kelas using shared data
// 	for _, classData := range sharedClasses {
// 		class := entities.Kelas{
// 			ID:          classData.ClassID,
// 			Name:        classData.Name,
// 			Tag:         classData.Tag,
// 			Description: classData.Description,
// 			Teacher:     classData.Teacher,
// 			TeacherID:   classData.TeacherID,
// 		}
// 		db.Create(&class)
// 	}

// 	// Seed Members using shared data
// 	fmt.Printf("Seeding %d members...\n", len(sharedUsers))
// 	for i, userData := range sharedUsers {
// 		member := entities.Member{
// 			Username:      userData.Username,
// 			Role:          userData.Role,
// 			User_userID:   userData.UserID,
// 			Kelas_kelasID: userData.ClassID,
// 			CreatedAt:     time.Now(),
// 			UpdatedAt:     time.Now(),
// 		}
// 		db.Create(&member)

// 		// Progress indicator for large datasets
// 		if (i+1)%100 == 0 {
// 			fmt.Printf("Seeded %d/%d members...\n", i+1, len(sharedUsers))
// 		}
// 	}

// 	// Seed Weeks and related data
// 	for _, classData := range sharedClasses {
// 		var assignmentFileID string
// 		for weekNum := 1; weekNum <= 4; weekNum++ {
// 			week := entities.Week{
// 				WeekNumber:    weekNum,
// 				Kelas_idKelas: classData.ClassID,
// 			}
// 			db.Create(&week)

// 			// Seed ItemPembelajaran with random YouTube URLs
// 			itemPembelajaran := entities.ItemPembelajaran{
// 				WeekID:           week.ID,
// 				HeadingPertemuan: fmt.Sprintf("Pertemuan %d - %s", weekNum, classData.Name),
// 				BodyPertemuan:    fmt.Sprintf("Materi pembelajaran minggu ke-%d untuk mata kuliah %s", weekNum, classData.Name),
// 				UrlVideo:         getRandomYouTubeURL(weekNum),
// 				FileName:         fmt.Sprintf("materi_%s_week_%d.pdf", classData.Tag, weekNum),
// 				FileId:           "6852508f4df1e68bde926047",
// 			}
// 			db.Create(&itemPembelajaran)

// 			// Seed Assignment (every 2 weeks)
// 			if weekNum%2 == 0 {
// 				// Determine assignment FileID based on class tag and week
// 				switch classData.Tag {
// 				case "EG": // English Grammar (Class 1)
// 					if weekNum == 2 {
// 						assignmentFileID = "68525b244df1e68bde926058"
// 					} else if weekNum == 4 {
// 						assignmentFileID = "68525b614df1e68bde92605b"
// 					}
// 				case "EC": // English Conversation (Class 2)
// 					if weekNum == 2 {
// 						assignmentFileID = "68525b7f4df1e68bde926060"
// 					} else if weekNum == 4 {
// 						assignmentFileID = "68525b8c4df1e68bde926063"
// 					}
// 				case "EL": // English Literature (Class 3)
// 					if weekNum == 2 {
// 						assignmentFileID = "68525ba44df1e68bde926068"
// 					} else if weekNum == 4 {
// 						assignmentFileID = "68525bb34df1e68bde92606b"
// 					}
// 				}

// 				assignment := entities.Assignment{
// 					Title:       fmt.Sprintf("Tugas %s - Minggu %d", classData.Name, weekNum),
// 					Description: fmt.Sprintf("Tugas praktikum untuk minggu ke-%d mata kuliah %s", weekNum, classData.Name),
// 					Deadline:    time.Now().AddDate(0, 0, 7),
// 					FileName:    fmt.Sprintf("tugas_%s_week_%d.pdf", classData.Tag, weekNum),
// 					FileId:      assignmentFileID,
// 					WeekID:      week.ID,
// 				}
// 				db.Create(&assignment)

// 				// Seed AssignmentSubmissions using shared user data
// 				studentsInClass := make([]SharedUserData, 0)
// 				for _, user := range sharedUsers {
// 					if user.ClassID == classData.ClassID && user.Role == entities.MemberRoleStudent {
// 						studentsInClass = append(studentsInClass, user)
// 					}
// 				}

// 				// Create submissions for first 3 students in class
// 				for j := 0; j < 3 && j < len(studentsInClass); j++ {
// 					submission := entities.AssignmentSubmission{
// 						AssignmentID: int(assignment.ID),
// 						UserID:       studentsInClass[j].UserID,
// 						IDFile:       "68429cdb04d68646d09139b8",
// 						FileName:     fmt.Sprintf("submission_%s.pdf", studentsInClass[j].Username),
// 						Score:        85 + j*5,
// 						Status:       entities.StatusSubmitted,
// 						CreatedAt:    time.Now(),
// 						UpdatedAt:    time.Now(),
// 					}
// 					db.Create(&submission)
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("Class Control data seeded successfully!")
// }

// // ========== SUMMARY ==========
// func PrintStaticUUIDs() {
// 	fmt.Println("\n========== STATIC UUID REFERENCES ==========")
// 	fmt.Println("CLASS IDs (untuk Assessment.ClassID):")
// 	fmt.Printf("  - English Grammar: %s\n", ClassWebProgID)
// 	fmt.Printf("  - English Conversation: %s\n", ClassDatabaseID)
// 	fmt.Printf("  - English Literature: %s\n", ClassAlgorithmID)

// 	fmt.Println("\nTEACHER IDs:")
// 	fmt.Printf("  - Ahmad Susanto: %s\n", TeacherAhmadID)
// 	fmt.Printf("  - Siti Nurhaliza: %s\n", TeacherSitiID)
// 	fmt.Printf("  - Budi Santoso: %s\n", TeacherBudiID)

// 	fmt.Println("\nSTUDENT IDs (untuk Assessment.Submission.UserID):")
// 	fmt.Printf("  - Alice Johnson: %s\n", StudentAliceID)
// 	fmt.Printf("  - Bob Smith: %s\n", StudentBobID)
// 	fmt.Printf("  - Charlie Brown: %s\n", StudentCharlieID)
// 	fmt.Printf("  - Diana Prince: %s\n", StudentDianaID)
// 	fmt.Printf("  - Edward Norton: %s\n", StudentEdwardID)
// 	fmt.Printf("  - Fiona Green: %s\n", StudentFionaID)
// 	fmt.Printf("  - George Washington: %s\n", StudentGeorgeID)
// 	fmt.Printf("  - Hannah Montana: %s\n", StudentHannahID)
// 	fmt.Printf("  - Ivan Petrov: %s\n", StudentIvanID)
// 	fmt.Printf("  - Julia Roberts: %s\n", StudentJuliaID)
// 	fmt.Printf("  - Kevin Hart: %s\n", StudentKevinID)
// 	fmt.Printf("  - Linda Hamilton: %s\n", StudentLindaID)
// 	fmt.Printf("  - Michael Jordan: %s\n", StudentMichaelID)
// 	fmt.Printf("  - Nancy Drew: %s\n", StudentNancyID)
// 	fmt.Printf("  - Oscar Wilde: %s\n", StudentOscarID)

// 	fmt.Println("\nTOTAL DATA SUMMARY:")
// 	allStudents := generateAllStudentUUIDs()
// 	fmt.Printf("  - Teachers: 3\n")
// 	fmt.Printf("  - Students: %d\n", len(allStudents))
// 	fmt.Printf("  - Total Users: %d\n", 3+len(allStudents))
// 	fmt.Printf("  - Classes: 3\n")

// 	fmt.Println("\nRANDOM YOUTUBE URLs USED:")
// 	for i, url := range randomYouTubeURLs {
// 		fmt.Printf("  - URL %d: %s\n", i+1, url)
// 	}
// }

package seed

import (
	"fmt"
	"time"

	database "LMSGo/config"
	entities "LMSGo/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeder() error {
	// Inisialisasi koneksi ke database
	db := database.SetUpDatabaseConnection()
	db.AutoMigrate(&entities.Kelas{}, &entities.Member{}, &entities.ItemPembelajaran{}, &entities.Assignment{}, &entities.AssignmentSubmission{}, &entities.Week{})
	
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

	// Teacher IDs (matching Python seeder)
	TeacherAhmadID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440101")
	TeacherSitiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440102")
	TeacherBudiID  = uuid.MustParse("550e8400-e29b-41d4-a716-446655440103")

	// Student IDs (matching Python seeder - first 15 named students)
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
	index := (weekNum - 1) % len(randomYouTubeURLs)
	return randomYouTubeURLs[index]
}

// Generate all student UUIDs (including the 1000 generated students)
func generateAllStudentUUIDs() []uuid.UUID {
	var studentUUIDs []uuid.UUID
	
	// Add the first 15 named students
	namedStudents := []uuid.UUID{
		StudentAliceID, StudentBobID, StudentCharlieID, StudentDianaID, StudentEdwardID,
		StudentFionaID, StudentGeorgeID, StudentHannahID, StudentIvanID, StudentJuliaID,
		StudentKevinID, StudentLindaID, StudentMichaelID, StudentNancyID, StudentOscarID,
	}
	studentUUIDs = append(studentUUIDs, namedStudents...)
	
	// Add the remaining 985 students (216-1215 from Python seeder)
	for i := 216; i <= 1215; i++ {
		uuidStr := fmt.Sprintf("550e8400-e29b-41d4-a716-44665544%04d", i)
		studentID := uuid.MustParse(uuidStr)
		studentUUIDs = append(studentUUIDs, studentID)
	}
	
	return studentUUIDs
}

func GenerateStaticData() ([]SharedClassData, []SharedUserData) {
	// Generate consistent class data with static UUIDs
	classes := []SharedClassData{
		{
			ClassID:     ClassWebProgID,
			Name:        "English Grammar",
			Tag:         "EG",
			Description: "Mata kuliah tata bahasa Inggris untuk pemahaman struktur kalimat",
			Teacher:     "Ahmad Susanto",
			TeacherID:   TeacherAhmadID,
		},
		{
			ClassID:     ClassDatabaseID,
			Name:        "English Conversation",
			Tag:         "EC",
			Description: "Mata kuliah percakapan bahasa Inggris untuk komunikasi sehari-hari",
			Teacher:     "Siti Nurhaliza",
			TeacherID:   TeacherSitiID,
		},
		{
			ClassID:     ClassAlgorithmID,
			Name:        "English Literature",
			Tag:         "EL",
			Description: "Mata kuliah sastra Inggris untuk pemahaman karya sastra klasik dan modern",
			Teacher:     "Budi Santoso",
			TeacherID:   TeacherBudiID,
		},
	}

	// Generate user data with static UUIDs
	var users []SharedUserData

	// Add teachers (matching Python seeder names)
	users = append(users, []SharedUserData{
		{UserID: TeacherAhmadID, Username: "Ahmad Susanto", Role: entities.MemberRoleTeacher, ClassID: ClassWebProgID},
		{UserID: TeacherSitiID, Username: "Siti Nurhaliza", Role: entities.MemberRoleTeacher, ClassID: ClassDatabaseID},
		{UserID: TeacherBudiID, Username: "Budi Santoso", Role: entities.MemberRoleTeacher, ClassID: ClassAlgorithmID},
	}...)

	// Get all student UUIDs
	allStudentUUIDs := generateAllStudentUUIDs()
	
	// Named students data (matching Python seeder exactly)
	namedStudents := []struct {
		ID       uuid.UUID
		Username string
	}{
		{StudentAliceID, "Alice Johnson"},
		{StudentBobID, "Bob Smith"},
		{StudentCharlieID, "Charlie Brown"},
		{StudentDianaID, "Diana Prince"},
		{StudentEdwardID, "Edward Norton"},
		{StudentFionaID, "Fiona Green"},
		{StudentGeorgeID, "George Washington"},
		{StudentHannahID, "Hannah Montana"},
		{StudentIvanID, "Ivan Petrov"},
		{StudentJuliaID, "Julia Roberts"},
		{StudentKevinID, "Kevin Hart"},
		{StudentLindaID, "Linda Hamilton"},
		{StudentMichaelID, "Michael Jordan"},
		{StudentNancyID, "Nancy Drew"},
		{StudentOscarID, "Oscar Wilde"},
	}

	// ASSIGN ALL STUDENTS TO ONE CLASS (English Grammar - First Class)
	targetClass := classes[0] // All students will be assigned to English Grammar class
	
	fmt.Printf("🎯 Assigning ALL %d students to class: %s\n", len(allStudentUUIDs), targetClass.Name)
	
	for studentIndex, studentID := range allStudentUUIDs {
		var username string
		
		// Use named student username if it's one of the first 15
		if studentIndex < len(namedStudents) {
			username = namedStudents[studentIndex].Username
		} else {
			// Generate username for numbered students (matching Python pattern)
			studentNum := studentIndex + 216 - 15 // Adjust for the offset
			username = fmt.Sprintf("Student%d Test", studentNum)
		}
		
		users = append(users, SharedUserData{
			UserID:   studentID,
			Username: username,
			Role:     entities.MemberRoleStudent,
			ClassID:  targetClass.ClassID, // All students assigned to English Grammar
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
	fmt.Printf("Seeding %d members...\n", len(sharedUsers))
	for i, userData := range sharedUsers {
		member := entities.Member{
			Username:      userData.Username,
			Role:          userData.Role,
			User_userID:   userData.UserID,
			Kelas_kelasID: userData.ClassID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		db.Create(&member)
		
		// Progress indicator for large datasets
		if (i+1)%100 == 0 {
			fmt.Printf("Seeded %d/%d members...\n", i+1, len(sharedUsers))
		}
	}

	// Seed Weeks and related data
	for _, classData := range sharedClasses {
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
				UrlVideo:         getRandomYouTubeURL(weekNum),
				FileName:         fmt.Sprintf("materi_%s_week_%d.pdf", classData.Tag, weekNum),
				FileId:           "6852508f4df1e68bde926047",
			}
			db.Create(&itemPembelajaran)

			// Seed Assignment (every 2 weeks)
			if weekNum%2 == 0 {
				// Determine assignment FileID based on class tag and week
				switch classData.Tag {
				case "EG": // English Grammar (Class 1)
					if weekNum == 2 {
						assignmentFileID = "68525b244df1e68bde926058"
					} else if weekNum == 4 {
						assignmentFileID = "68525b614df1e68bde92605b"
					}
				case "EC": // English Conversation (Class 2)
					if weekNum == 2 {
						assignmentFileID = "68525b7f4df1e68bde926060"
					} else if weekNum == 4 {
						assignmentFileID = "68525b8c4df1e68bde926063"
					}
				case "EL": // English Literature (Class 3)
					if weekNum == 2 {
						assignmentFileID = "68525ba44df1e68bde926068"
					} else if weekNum == 4 {
						assignmentFileID = "68525bb34df1e68bde92606b"
					}
				}

				assignment := entities.Assignment{
					Title:       fmt.Sprintf("Tugas %s - Minggu %d", classData.Name, weekNum),
					Description: fmt.Sprintf("Tugas praktikum untuk minggu ke-%d mata kuliah %s", weekNum, classData.Name),
					Deadline:    time.Now().AddDate(0, 0, 7),
					FileName:    fmt.Sprintf("tugas_%s_week_%d.pdf", classData.Tag, weekNum),
					FileId:      assignmentFileID,
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
	fmt.Printf("  - English Grammar: %s (⭐ ALL STUDENTS ASSIGNED HERE)\n", ClassWebProgID)
	fmt.Printf("  - English Conversation: %s (Empty - no students assigned)\n", ClassDatabaseID)
	fmt.Printf("  - English Literature: %s (Empty - no students assigned)\n", ClassAlgorithmID)

	fmt.Println("\nTEACHER IDs:")
	fmt.Printf("  - Ahmad Susanto (English Grammar): %s\n", TeacherAhmadID)
	fmt.Printf("  - Siti Nurhaliza (English Conversation): %s\n", TeacherSitiID)
	fmt.Printf("  - Budi Santoso (English Literature): %s\n", TeacherBudiID)

	fmt.Println("\nSTUDENT IDs (untuk Assessment.Submission.UserID):")
	fmt.Printf("  - Alice Johnson: %s\n", StudentAliceID)
	fmt.Printf("  - Bob Smith: %s\n", StudentBobID)
	fmt.Printf("  - Charlie Brown: %s\n", StudentCharlieID)
	fmt.Printf("  - Diana Prince: %s\n", StudentDianaID)
	fmt.Printf("  - Edward Norton: %s\n", StudentEdwardID)
	fmt.Printf("  - Fiona Green: %s\n", StudentFionaID)
	fmt.Printf("  - George Washington: %s\n", StudentGeorgeID)
	fmt.Printf("  - Hannah Montana: %s\n", StudentHannahID)
	fmt.Printf("  - Ivan Petrov: %s\n", StudentIvanID)
	fmt.Printf("  - Julia Roberts: %s\n", StudentJuliaID)
	fmt.Printf("  - Kevin Hart: %s\n", StudentKevinID)
	fmt.Printf("  - Linda Hamilton: %s\n", StudentLindaID)
	fmt.Printf("  - Michael Jordan: %s\n", StudentMichaelID)
	fmt.Printf("  - Nancy Drew: %s\n", StudentNancyID)
	fmt.Printf("  - Oscar Wilde: %s\n", StudentOscarID)

	allStudents := generateAllStudentUUIDs()
	fmt.Println("\nSTUDENT DISTRIBUTION:")
	fmt.Printf("  🎯 English Grammar Class: %d students (ALL STUDENTS)\n", len(allStudents))
	fmt.Printf("  📚 English Conversation Class: 0 students\n")
	fmt.Printf("  📖 English Literature Class: 0 students\n")

	fmt.Println("\nTOTAL DATA SUMMARY:")
	fmt.Printf("  - Teachers: 3\n")
	fmt.Printf("  - Students: %d\n", len(allStudents))
	fmt.Printf("  - Total Users: %d\n", 3+len(allStudents))
	fmt.Printf("  - Classes: 3 (but all students in 1 class)\n")

	fmt.Println("\nRANDOM YOUTUBE URLs USED:")
	for i, url := range randomYouTubeURLs {
		fmt.Printf("  - URL %d: %s\n", i+1, url)
	}
}