package college

import (
	"college/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, client *ent.Client) {
	classRepository := NewClassService(client)
	classHandler := NewClassHandler(classRepository)

	studentRepository := NewStudentService(client)
	studentHandler := NewStudentHandler(studentRepository, classRepository)

	departmentRepository := NewDepartmentService(client)
	departmentHandler := NewDepartmentHandler(departmentRepository)

	facade := NewFacade(studentHandler, departmentHandler, classHandler)

	// health-check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "up and running...",
			"data":    nil,
		})
	})

	// API v1
	{
		v1 := router.Group("/api/v1")

		v1.POST("/students", facade.CreateStudent)
		v1.GET("/students", facade.GetStudents)
		v1.PATCH("/students/:id/details", facade.UpdateStudentDetails)
		v1.PATCH("/students/:id/department", facade.UpdateStudentDepartment)
		v1.DELETE("/students/:id", facade.DeleteStudent)
		v1.GET("/students/:id", facade.GetStudentDetails)
		v1.PATCH("/students/:id/class-registration", facade.ClassRegistration)

		v1.POST("/departments", facade.CreateDepartment)

		v1.GET("/classes", facade.GetClasses)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "resource does not exist",
		})
	})
}
