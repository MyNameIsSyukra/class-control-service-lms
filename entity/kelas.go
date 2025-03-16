package entities

type Kelas struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
	TeacherID   int    `json:"teacher_id"`
}
