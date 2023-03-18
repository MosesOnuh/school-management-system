package handlers

import (
	"net/http"
	"management-system/domain"
	"management-system/utils/error_utils"
	"strconv"

	"management-system/services"

	"github.com/gin-gonic/gin"
)


func convertIdType(id string) (int64, error_utils.AppErr) {
	converteId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
			return 0, error_utils.AppBadRequestError("id must be a number")
	}

	return converteId, nil
}

func DropStudentTable(c *gin.Context){
	err := services.TablesService.DropStudentTable()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully dropped student table",
	})
}

func CreateStudentCoursesTable(c *gin.Context){
	err := services.TablesService.CreateStudentCoursesTable()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created student table",
	})
}

func CreateStudentTable(c *gin.Context){
	err := services.TablesService.CreateStudentTable()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created student table",
	})
}



func CreateCourseTable(c *gin.Context){
	err := services.TablesService.CreateCourseTable()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created course table",
	})
}


func GetAllStudentCourses(c *gin.Context) {
	courses, err := services.StudentCourseService.GetAllStudentCourses()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got all student_courses",
		"data": courses,
	})
}




func GetAStudentCourses(c *gin.Context){
	id := c.Param("id")
	studentCourseId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	studentCourses, err := services.StudentCourseService.GetAStudentCourses(studentCourseId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got student courses",
		"data": studentCourses,
	})
}

func GetCourse(c *gin.Context){
	id := c.Param("id")
	courseId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	course, err := services.CourseService.GetCourse(courseId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got course",
		"data": course,
	})
}


func AddStudentCourse(c *gin.Context){
	var studentCourse domain.StudentCourse
	bindErr := c.ShouldBindJSON(&studentCourse)
	if bindErr != nil {
		Err := error_utils.AppUnprocessibleEntityError("invalid json body")
		c.JSON(Err.Status(), Err)
		return
	} 

	addedStudentCourse, err := services.StudentCourseService.AddStudentCourse(&studentCourse)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully added student course",
		"data": addedStudentCourse,
	})
}

func GetAllCourses(c *gin.Context) {
	courses, err := services.CourseService.GetAllCourses()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got all courses",
		"data": courses,
	})
}

func CreateCourse(c *gin.Context){
	var course domain.Course
	bindErr := c.ShouldBindJSON(&course)
	if bindErr != nil {
		Err := error_utils.AppUnprocessibleEntityError("invalid json body")
		c.JSON(Err.Status(), Err)
		return
	} 

	createdCourse, err := services.CourseService.CreateCourse(&course)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created course",
		"data": createdCourse,
	})
}

func UpdateCourse(c *gin.Context){
	id := c.Param("id")
	courseId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var course domain.Course
	bindErr := c.ShouldBindJSON(&course)
	if bindErr != nil {
		Err := error_utils.AppUnprocessibleEntityError("invalid json body")
		c.JSON(Err.Status(), Err)
		return
	} 

	course.Id = courseId
	updatedCourse, err := services.CourseService.UpdateCourse(&course)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated course",
		"data": updatedCourse,
	})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	courseId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	err = services.CourseService.DeleteCourse(courseId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}

/////////student ////////////////////////////
func GetStudent(c *gin.Context){
	id := c.Param("id")
	studentId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	student, err := services.StudentService.GetStudent(studentId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got student",
		"data": student,
	})
}

func GetAllStudents(c *gin.Context) {
	students, err := services.StudentService.GetAllStudents()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully got all students",
		"data": students,
	})
}

func CreateStudent(c *gin.Context){
	var student domain.Student
	bindErr := c.ShouldBindJSON(&student)
	if bindErr != nil {
		Err := error_utils.AppUnprocessibleEntityError("invalid json body")
		c.JSON(Err.Status(), Err)
		return
	} 

	std, err := services.StudentService.CreateStudent(&student)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created student",
		"data": std,
	})
}

func UpdateStudent(c *gin.Context){
	id := c.Param("id")
	studentId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var student domain.Student
	bindErr := c.ShouldBindJSON(&student)
	if bindErr != nil {
		Err := error_utils.AppUnprocessibleEntityError("invalid json body")
		c.JSON(Err.Status(), Err)
		return
	} 

	student.Id = studentId
	std, err := services.StudentService.UpdateStudent(&student)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated Student",
		"data": std,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	studentId, err := convertIdType(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	err = services.StudentService.DeleteStudent(studentId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}