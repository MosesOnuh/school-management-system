package domain

import (
	"database/sql"
	"fmt"
	"log"
	"management-system/utils/error_utils"
	"context"
	"time"
)

type managementRepo struct {
	db *sql.DB
}

const (
	queryCreateStudentTable = "CREATE TABLE students (id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100)  NOT NULL, sex VARCHAR(10) NOT NULL,created_at INT NOT NULL);"
	queryCreateCourseTable = "CREATE TABLE courses ( id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL, unit INT NOT NULL, semester INT NOT NULL, created_at INT NOT NULL, UNIQUE KEY courseInfo (title, unit, semester), CHECK (unit BETWEEN 0 AND 5 AND semester BETWEEN 1 AND 2));"
	queryCreateStudentCoursesTable= "CREATE TABLE student_courses (studentId INT NOT NULL, courseId INT NOT NULL, FOREIGN KEY (studentId) REFERENCES students(id) ON DELETE CASCADE, FOREIGN KEY (courseId) REFERENCES courses(id) ON DELETE CASCADE, PRIMARY KEY(studentId, courseId) );"
	queryDropStudentTable = "DROP TABLE students;"

	queryInsertStudent  = "INSERT INTO students (name, sex, created_at) VALUES(?, ?, ?);"
	queryGetAllStudents    = "SELECT id, name, sex, created_at  FROM students;"
	queryGetStudent     = "SELECT * FROM students WHERE id = ?;"
	queryUpdateAStudent = "UPDATE students SET name = ?, sex = ? WHERE id =?;"
	queryDeleteStudent  = "DELETE FROM students WHERE id = ?;"

	queryInsertCourse  = "INSERT INTO courses ( title, unit, semester, created_at) VALUES(?, ?, ?, ?);"
	queryGetAllCourses    = "SELECT * FROM courses;"
	queryGetCourse     = "SELECT * FROM courses WHERE id = ?;"
	queryUpdateACourse = "UPDATE courses SET title = ?, unit =?, semester = ?;"
	queryDeleteCourse  = "DELETE FROM courses WHERE id = ?;"

	queryGetStudentCourses        = "SELECT * FROM students INNER JOIN student_courses ON students.id = student_courses.studentId INNER JOIN courses ON student_courses.courseId = courses.id ;"

	queryGetAStudentCourses        = "SELECT * FROM students INNER JOIN student_courses ON students.id = student_courses.studentId INNER JOIN courses ON student_courses.courseId = courses.id WHERE students.id =?"

	queryAddToAStudentCourses      = "INSERT INTO student_courses (studentId, CourseId) VALUES(?, ?);"
	// queryRemoveFromAStudentCourses = "DELETE FROM student_courses WHERE studentId = ? AND courseId = ?;"
)

var (
	DbRepo TableInterface = &managementRepo{}
)

type TableInterface interface {
	Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string ) *sql.DB
	DropStudentTable() error_utils.AppErr
	CreateStudentTable() error_utils.AppErr
	CreateCourseTable() error_utils.AppErr
	CreateStudentCoursesTable() error_utils.AppErr
	CreateStudent(student *Student) (*Student, error_utils.AppErr)
	GetAllStudents() ([]Student, error_utils.AppErr)
	GetStudent(id int64) (*Student, error_utils.AppErr)
	UpdateStudent(student *Student) (*Student, error_utils.AppErr)
	DeleteStudent(id int64) error_utils.AppErr
	CreateCourse(course *Course) (*Course, error_utils.AppErr)
	GetAllCourses() ([]Course, error_utils.AppErr)
	GetCourse(id int64) (*Course, error_utils.AppErr)
	UpdateCourse(course *Course) (*Course, error_utils.AppErr)
	DeleteCourse(id int64) error_utils.AppErr
	GetAllStudentCourses() ([]StudentCouresResp, error_utils.AppErr)
	GetAStudentCourses(studentId int64) ([]StudentCouresResp, error_utils.AppErr)
	AddStudentCourse (studentCourse *StudentCourse) (*StudentCourse, error_utils.AppErr)
}


func (r *managementRepo) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string ) *sql.DB {
	DbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbName)
	
	var err error
	r.db, err = sql.Open(Dbdriver, DbDSN)
	if err != nil {
		log.Fatal("Error occurred when connecting to database:", err)
	}
	fmt.Printf("Successfully connected to the %s database", Dbdriver)
	return r.db
}

func (r *managementRepo) DropStudentTable() error_utils.AppErr {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancelfunc()

	_, err := r.db.ExecContext(ctx, queryDropStudentTable)  
	if err != nil {  
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when droping student table: %s", err.Error()))
	} 
	return nil
}

func (r *managementRepo) CreateStudentCoursesTable() error_utils.AppErr {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancelfunc()

	_, err := r.db.ExecContext(ctx, queryCreateStudentCoursesTable)  
	if err != nil {  
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when creating student_courses table: %s", err.Error()))
	} 
	return nil
}

func (r *managementRepo) CreateStudentTable() error_utils.AppErr {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancelfunc()

	_, err := r.db.ExecContext(ctx, queryCreateStudentTable)  
	if err != nil {  
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when creating student table: %s", err.Error()))
	} 
	return nil
}

func (r *managementRepo) CreateCourseTable() error_utils.AppErr {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancelfunc()

	_, err := r.db.ExecContext(ctx, queryCreateCourseTable)  
	if err != nil {  
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when creating course table: %s", err.Error()))
	} 
	return nil
}

func (r *managementRepo) GetAllStudentCourses() ([]StudentCouresResp, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryGetStudentCourses )
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing student_courses to get: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	defer rows.Close()

	studentCourses := make([]StudentCouresResp, 0)

	for rows.Next() {
		var studentCourse StudentCouresResp

		err = rows.Scan(&studentCourse.StudentId, &studentCourse.Name, &studentCourse.Sex, &studentCourse.StudentCreatedAt, &studentCourse.CourseId, &studentCourse.Title, &studentCourse.Unit, &studentCourse.Semester, &studentCourse.CourseCreatedAt )
		if err != nil {
			return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occurred when getting student_courses: %s", err.Error()))
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	if len(studentCourses) == 0 {
		return nil, error_utils.AppNotFoundError("no records found")
	}
	return studentCourses, nil
}

func (r *managementRepo) GetAStudentCourses(studentId int64) ([]StudentCouresResp, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryGetAStudentCourses )
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing student_courses to get: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query(studentId)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	defer rows.Close()

	studentCourses := make([]StudentCouresResp, 0)

	for rows.Next() {
		var studentCourse StudentCouresResp

		err = rows.Scan(&studentCourse.StudentId, &studentCourse.Name, &studentCourse.Sex, &studentCourse.StudentCreatedAt, &studentCourse.CourseId, &studentCourse.Title, &studentCourse.Unit, &studentCourse.Semester, &studentCourse.CourseCreatedAt )
		if err != nil {
			return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occurred when getting student_courses: %s", err.Error()))
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	if len(studentCourses) == 0 {
		return nil, error_utils.AppNotFoundError("no records found")
	}
	return studentCourses, nil
}

func (r *managementRepo) AddStudentCourse (studentCourse *StudentCourse) (*StudentCourse, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryAddToAStudentCourses)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("error occured when preparing to add studentCourse: %s", err.Error()))
	}

	defer stmt.Close()

	insertedResult, err := stmt.Exec(studentCourse.StudentId, studentCourse.CourseId)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	_, err = insertedResult.LastInsertId()
	if err != nil {
		return nil, error_utils.AppInternalServerError((fmt.Sprintf("error occured when saving student_course: %s", err.Error())))
	}

	return studentCourse, nil
}

func (r *managementRepo) CreateStudent (student *Student) (*Student, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryInsertStudent)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("error occured when preparing to create student: %s", err.Error()))
	}

	defer stmt.Close()

	insertedResult, err := stmt.Exec(student.Name, student.Sex, student.CreatedAt)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	studentId, err := insertedResult.LastInsertId()
	if err != nil {
		return nil, error_utils.AppInternalServerError((fmt.Sprintf("error occured when saving student: %s", err.Error())))
	}

	student.Id = studentId
	return student, nil
}

func (r *managementRepo) GetAllStudents() ([]Student, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryGetAllStudents )
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing students to get: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	defer rows.Close()

	students :=make([]Student, 0)

	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Id, &student.Name, &student.Sex, &student.CreatedAt)
		if err != nil {
			return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occurred when getting student: %s", err.Error()))
		}
		students = append(students, student)
	}

	if len(students) == 0 {
		return nil, error_utils.AppNotFoundError("no records found")
	}
	return students, nil
}

func (r *managementRepo) GetStudent (id int64) (*Student, error_utils.AppErr){
	stmt, err := r.db.Prepare(queryGetStudent)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing student: %s", err.Error()))
	}
	defer stmt.Close()

	var student Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Sex, &student.CreatedAt);

	if err != nil {
		return nil, error_utils.ParseErr(err)
	}

	return &student, nil
}



func (r *managementRepo) UpdateStudent(student *Student) (*Student, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryUpdateAStudent)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing student to update: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.Name, student.Sex)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	return student, nil
}

func (r *managementRepo) DeleteStudent(id int64) error_utils.AppErr {
	stmt, err := r.db.Prepare(queryDeleteStudent)
	if err != nil {
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when deleting student: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when deleting a student %s", err.Error()))
	}
	return nil
}

func (r *managementRepo) CreateCourse (course *Course) (*Course, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryInsertCourse)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("error occured when preparing to create course: %s", err.Error()))
	}

	defer stmt.Close()

	insertedResult, err := stmt.Exec(course.Title, course.Unit, course.Semester, course.CreatedAt)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	courseId, err := insertedResult.LastInsertId()
	if err != nil {
		return nil, error_utils.AppInternalServerError((fmt.Sprintf("error occured when saving course: %s", err.Error())))
	}

	course.Id = courseId
	return course, nil
}

func (r *managementRepo) GetAllCourses() ([]Course, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryGetAllCourses)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing course: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	defer rows.Close()

	courses :=make([]Course, 0)

	for rows.Next() {
		var course Course
		err = rows.Scan(&course.Id, &course.Title, &course.Unit, &course.Semester, &course.CreatedAt)
		if err != nil {
			return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occurred when getting student: %s", err.Error()))
		}
		courses = append(courses, course)
	}

	if len(courses) == 0 {
		return nil, error_utils.AppNotFoundError("no records found")
	}
	return courses, nil
}

func (r *managementRepo) GetCourse (id int64) (*Course, error_utils.AppErr){
	stmt, err := r.db.Prepare(queryGetCourse)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing course to get: %s", err.Error()))
	}
	defer stmt.Close()

	var course Course
	err = stmt.QueryRow(id).Scan(&course.Id, &course.Title, &course.Unit, &course.Semester, &course.CreatedAt);
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}

	return &course, nil
}

func (r *managementRepo) 	UpdateCourse(course *Course) (*Course, error_utils.AppErr) {
	stmt, err := r.db.Prepare(queryUpdateACourse)
	if err != nil {
		return nil, error_utils.AppInternalServerError(fmt.Sprintf("Error occured when preparing course to update: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Title, course.Unit, course.Semester)
	if err != nil {
		return nil, error_utils.ParseErr(err)
	}
	return course, nil
}

func (r *managementRepo) DeleteCourse(id int64) error_utils.AppErr{
	stmt, err := r.db.Prepare(queryDeleteCourse)
	if err != nil {
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when deleting course: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return error_utils.AppInternalServerError(fmt.Sprintf("error occured when deleting a course %s", err.Error()))
	}
	return nil
}


