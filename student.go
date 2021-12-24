package main

import (
	"context"
	"ent-demo/ent"
	"time"
)

/**
Model
*/
type Student struct {
	ID              int       `json:"id"`
	Firstname       string    `json:"first_name"`
	Lastname        string    `json:"last_name"`
	Email           string    `json:"email"`
	AdmissionNumber string    `json:"admission_number"`
	Year            int       `json:"year"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func parseToStudent(model *ent.Student) *Student {
	return nil
}

type StudentHandler struct {
	sr StudentRepository
}

func NewStudentService(client *ent.Client) StudentRepository {
	return &Datastore{c: client}
}

func NewStudentHandler(sr StudentRepository) *StudentHandler {
	return &StudentHandler{sr}
}

/**
Repository
*/
func (d *Datastore) FindListStudents(ctx context.Context) ([]*ent.Student, error) {
	entities, err := d.c.Student.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

/**
Handlers
*/
func (sh *StudentHandler) GetStudents(ctx context.Context) {
	sh.sr.FindListStudents(ctx)
}
