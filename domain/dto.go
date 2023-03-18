package domain

type Student struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	CreatedAt int64  `json:"created_at"`
}

type Course struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Unit      int64  `json:"unit"`
	Semester  int64  `json:"semester"`
	CreatedAt int64  `json:"created_at"`
}

type StudentCouresResp struct {
	StudentId int64  `json:"studentId"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	StudentCreatedAt string `json:"student_created_at"`
	CourseId  int64  `json:"courseId"`
	Title     string `json:"title"`
	Unit      int64  `json:"unit"`
	Semester  int64  `json:"semester"`
	CourseCreatedAt string `json:"course_created_at"`
}

type StudentCourse struct {
	StudentId int64     `json:"student_id"`
	CourseId  int64     `json:"course_id"`
}
