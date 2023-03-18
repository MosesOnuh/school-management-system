package services

import (
	"management-system/domain"
	"management-system/utils/error_utils"
	"time"
)

type managementService struct{}

var (
	StudentService studentServiceInterface = &managementService{}
	CourseService  courseServiceInterface  = &managementService{}
	TablesService  tablesInterface         = &managementService{}
	StudentCourseService studentCoursesServiceInterface = &managementService{}
)

type tablesInterface interface {
	CreateStudentTable() error_utils.AppErr
	CreateCourseTable() error_utils.AppErr
	CreateStudentCoursesTable() error_utils.AppErr
	DropStudentTable() error_utils.AppErr
}

type courseServiceInterface interface {
	CreateCourse(course *domain.Course) (*domain.Course, error_utils.AppErr)
	GetAllCourses() ([]domain.Course, error_utils.AppErr)
	GetCourse(id int64) (*domain.Course, error_utils.AppErr)
	UpdateCourse(course *domain.Course) (*domain.Course, error_utils.AppErr)
	DeleteCourse(id int64) error_utils.AppErr
}

type studentServiceInterface interface {
	CreateStudent(student *domain.Student) (*domain.Student, error_utils.AppErr)
	GetAllStudents() ([]domain.Student, error_utils.AppErr)
	GetStudent(id int64) (*domain.Student, error_utils.AppErr)
	UpdateStudent(student *domain.Student) (*domain.Student, error_utils.AppErr)
	DeleteStudent(id int64) error_utils.AppErr
}

type studentCoursesServiceInterface interface {
	GetAllStudentCourses() ([]domain.StudentCouresResp, error_utils.AppErr)
	GetAStudentCourses(studentId int64) ([]domain.StudentCouresResp, error_utils.AppErr)
	AddStudentCourse (studentCourse *domain.StudentCourse) (*domain.StudentCourse, error_utils.AppErr)
}

func (m *managementService) DropStudentTable() error_utils.AppErr {
	err := domain.DbRepo.DropStudentTable()
	if err != nil {
		return err
	}
	return nil
}

func (m *managementService) CreateStudentCoursesTable() error_utils.AppErr {
	err := domain.DbRepo.CreateStudentCoursesTable()
	if err != nil {
		return err
	}
	return nil
}

func (m *managementService) CreateStudentTable() error_utils.AppErr {
	err := domain.DbRepo.CreateStudentTable()
	if err != nil {
		return err
	}
	return nil
}

func (m *managementService) CreateCourseTable() error_utils.AppErr {
	err := domain.DbRepo.CreateCourseTable()
	if err != nil {
		return err
	}
	return nil
}

func (m *managementService) CreateStudent(student *domain.Student) (*domain.Student, error_utils.AppErr) {
	student.CreatedAt = time.Now().Unix()
	createdStudent, err := domain.DbRepo.CreateStudent(student)
	if err != nil {
		return nil, err
	}

	return createdStudent, nil
}

func (m *managementService)GetAllStudentCourses() ([]domain.StudentCouresResp, error_utils.AppErr){
	studentCourses, err := domain.DbRepo.GetAllStudentCourses()
	if err != nil {
		return nil, err
	}

	return studentCourses, nil
}

func (m *managementService) GetAStudentCourses(studentId int64) ([]domain.StudentCouresResp, error_utils.AppErr){
	studentCourses, err := domain.DbRepo.GetAStudentCourses(studentId)
	if err != nil {
		return nil, err
	}

	return studentCourses, nil
}

func (m *managementService) AddStudentCourse (studentCourse *domain.StudentCourse) (*domain.StudentCourse, error_utils.AppErr) {
	createdStudent, err := domain.DbRepo. AddStudentCourse(studentCourse)
	if err != nil {
		return nil, err
	}

	return createdStudent, nil
}

func (m *managementService) GetAllStudents() ([]domain.Student, error_utils.AppErr) {
	students, err := domain.DbRepo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (m *managementService) GetStudent(id int64) (*domain.Student, error_utils.AppErr) {
	student, err := domain.DbRepo.GetStudent(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (m *managementService) UpdateStudent(student *domain.Student) (*domain.Student, error_utils.AppErr) {
	currentStudent, err := domain.DbRepo.GetStudent(student.Id)
	if err != nil {
		return nil, err
	}

	currentStudent.Name = student.Name
	currentStudent.Sex = student.Sex

	updatedStudent, err := domain.DbRepo.UpdateStudent(currentStudent)
	if err != nil {
		return nil, err
	}

	return updatedStudent, nil
}

func (m *managementService) DeleteStudent(id int64) error_utils.AppErr {
	student, err := domain.DbRepo.GetStudent(id)
	if err != nil {
		return err
	}

	err = domain.DbRepo.DeleteStudent(student.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *managementService) CreateCourse(course *domain.Course) (*domain.Course, error_utils.AppErr) {
	course.CreatedAt = time.Now().Unix()
	createdCourse, err := domain.DbRepo.CreateCourse(course)
	if err != nil {
		return nil, err
	}

	return createdCourse, nil
}

func (m *managementService) GetAllCourses() ([]domain.Course, error_utils.AppErr) {
	courses, err := domain.DbRepo.GetAllCourses()
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (m *managementService) GetCourse(id int64) (*domain.Course, error_utils.AppErr) {
	course, err := domain.DbRepo.GetCourse(id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (m *managementService) UpdateCourse(course *domain.Course) (*domain.Course, error_utils.AppErr) {
	currentCourse, err := domain.DbRepo.GetCourse(course.Id)
	if err != nil {
		return nil, err
	}

	currentCourse.Title = course.Title
	currentCourse.Unit = course.Unit
	currentCourse.Semester = course.Semester

	updatedCourse, err := domain.DbRepo.UpdateCourse(currentCourse)
	if err != nil {
		return nil, err
	}

	return updatedCourse, nil
}

func (m *managementService) DeleteCourse(id int64) error_utils.AppErr {
	course, err := domain.DbRepo.GetCourse(id)
	if err != nil {
		return err
	}

	err = domain.DbRepo.DeleteCourse(course.Id)
	if err != nil {
		return err
	}
	return nil
}
