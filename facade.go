package college

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Facade struct {
	sh *StudentHandler
	dh *DepartmentHandler
	ch *ClassHandler
}

func NewFacade(sh *StudentHandler, dh *DepartmentHandler, ch *ClassHandler) *Facade {
	return &Facade{sh, dh, ch}
}

/**
Student Facade
*/
func (f *Facade) GetStudents(ctx *gin.Context) {
	var vm GetStudentsVM

	if err := ctx.ShouldBindQuery(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})
	}

	students, paginator, err := f.sh.GetStudents(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
			"meta":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    students,
		"meta":    paginator,
	})
}

func (f *Facade) GetStudentDetails(ctx *gin.Context) {
	sId := ctx.Param("id")

	student, err := f.sh.GetStudentDetails(ctx, sId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    student,
	})
}

func (f *Facade) CreateStudent(ctx *gin.Context) {
	var vm CreateStudentVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	student, err := f.sh.CreateStudent(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    student,
	})
}

func (f *Facade) UpdateStudentDetails(ctx *gin.Context) {
	var vm UpdateStudentDetailsVM
	ID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	student, err := f.sh.UpdateStudentDetails(ctx, vm, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    student,
	})
}

func (f *Facade) UpdateStudentDepartment(ctx *gin.Context) {
	var vm UpdateStudentDepartmentVM
	ID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	student, err := f.sh.UpdateStudentDepartment(ctx, vm, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    student,
	})
}

// ClassRegistration godoc
// @Summary Class registration endpoint
// @Accept  json
// @Produce  json
// @Param ClassRegistrationVM body ClassRegistrationVM true "success"
// @Success 200 {object} APIResponse
// @Router /students/:id [get]
func (f *Facade) ClassRegistration(ctx *gin.Context) {
	var vm ClassRegistrationVM
	ID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	student, err := f.sh.ClassRegistration(ctx, vm, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    student,
	})
}
func (f *Facade) ClassUnregistration(ctx *gin.Context) {}

func (f *Facade) DeleteStudent(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := f.sh.DeleteStudent(ctx, ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    nil,
	})
}

/**
Department Facade
*/
func (f *Facade) CreateDepartment(ctx *gin.Context) {
	var vm CreateDepartmentVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	department, err := f.dh.CreateDepartment(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    department,
	})
}

/**
Class Facade
*/
func (f *Facade) GetClasses(ctx *gin.Context) {
	var vm GetClassesVM

	if err := ctx.ShouldBindQuery(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})
	}

	classes, paginator, err := f.ch.GetClasses(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
			"meta":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    classes,
		"meta":    paginator,
	})
}
