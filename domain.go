package main

import (
	"context"
	"ent-demo/ent"
)

type Datastore struct {
	c *ent.Client
}

type StudentRepository interface {
	FindListStudents(context.Context) ([]*ent.Student, error)
}

type DepartmentRepository interface {
	FindListDepartments()
}

type CourseRepository interface {
	// FindListStudents()
}

type LecturerRepository interface {
	// FindListStudents()
}
