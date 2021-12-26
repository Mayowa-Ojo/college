package college

import (
	"college/ent"
	"college/ent/department"
	"context"
	"time"

	"github.com/google/uuid"
)

/**
Model
*/
type Department struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Telephone string    `json:"telephone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func parseToDepartment(model *ent.Department) *Department {
	return &Department{
		ID:        model.ID,
		Name:      model.Name,
		Code:      model.Code,
		Telephone: model.Telephone,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func NewDepartmentService(client *ent.Client) DepartmentRepository {
	return &Datastore{c: client}
}

/**
Repository
*/
func (d *Datastore) FindListDepartments(ctx context.Context) ([]*ent.Department, error) {
	entities, err := d.c.Department.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (d *Datastore) FindListDepartmentsWithStudents(ctx context.Context) ([]*ent.Department, error) {
	entities, err := d.c.Department.Query().WithStudents().All(ctx)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (d *Datastore) FindDepartmentByID(ctx context.Context, ID uuid.UUID) (*ent.Department, error) {
	entity, err := d.c.Department.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindDepartmentByName(ctx context.Context, name string) (*ent.Department, error) {
	entity, err := d.c.Department.Query().Where(department.NameEqualFold(name)).First(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) CreateDepartment(ctx context.Context, dpt *Department) (*ent.Department, error) {
	entity, err := d.c.Department.
		Create().
		SetName(dpt.Name).
		SetTelephone(dpt.Telephone).
		SetCode(dpt.Code).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

/**
Handlers
*/
type DepartmentHandler struct {
	dr DepartmentRepository
}

func NewDepartmentHandler(dr DepartmentRepository) *DepartmentHandler {
	return &DepartmentHandler{dr}
}

func (dh *DepartmentHandler) GetDepartments() {}

func (dh *DepartmentHandler) CreateDepartment(ctx context.Context, vm CreateDepartmentVM) (*Department, error) {
	_, err := dh.dr.FindDepartmentByName(ctx, vm.Name)
	if err == nil {
		return nil, ErrDepartmentAlreadyExists
	}

	d := &Department{
		Name:      vm.Name,
		Code:      vm.Code,
		Telephone: vm.Telephone,
	}

	entity, err := dh.dr.CreateDepartment(ctx, d)
	if err != nil {
		return nil, err
	}

	return parseToDepartment(entity), nil
}
