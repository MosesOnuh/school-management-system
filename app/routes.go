package app

import (
	"management-system/handlers"
)

func routes(){
	tableRouter := router.Group("tables")
	studentRouter := router.Group("student")
	courseRouter := router.Group("course")
	studentCourses := router.Group("studentCourses")

	tableRouter.POST("/students", handlers.CreateStudentTable) 
	tableRouter.POST("/students/remove", handlers.DropStudentTable) 
	tableRouter.POST("/course", handlers.CreateCourseTable) 
	tableRouter.POST("/studentCourses", handlers.CreateStudentCoursesTable) 

	studentRouter.GET("/:id", handlers.GetStudent)
	studentRouter.GET("/", handlers.GetAllStudents)
	studentRouter.POST("/", handlers.CreateStudent)
	studentRouter.PUT("/:id", handlers.UpdateStudent)
	studentRouter.DELETE("/:id", handlers.DeleteStudent)

	courseRouter.GET("/:id", handlers.GetCourse)
	courseRouter.GET("/", handlers.GetAllCourses)
	courseRouter.POST("/", handlers.CreateCourse)
	courseRouter.PUT("/:id", handlers.UpdateCourse)
	courseRouter.DELETE("/:id", handlers.DeleteCourse)

	studentCourses.GET("/", handlers.GetAllStudentCourses)
	studentCourses.GET("/:id", handlers.GetAStudentCourses)
	studentCourses.POST("/", handlers.AddStudentCourse)
}


