package examManage

import "time"

type ExamInfo struct {
	Id           int
	ExamPlanID   int
	TeachClassID int
	ClassName    string
	StudentID    int
	StudentName  string
	TeacherID    int
	Screenshot   string
	IPAddress    string
	UploadTime   time.Time
}
