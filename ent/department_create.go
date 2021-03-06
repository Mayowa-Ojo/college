// Code generated by entc, DO NOT EDIT.

package ent

import (
	"college/ent/department"
	"college/ent/staff"
	"college/ent/student"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetCode sets the "code" field.
func (dc *DepartmentCreate) SetCode(s string) *DepartmentCreate {
	dc.mutation.SetCode(s)
	return dc
}

// SetTelephone sets the "telephone" field.
func (dc *DepartmentCreate) SetTelephone(s string) *DepartmentCreate {
	dc.mutation.SetTelephone(s)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DepartmentCreate) SetCreatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DepartmentCreate) SetUpdatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DepartmentCreate) SetID(u uuid.UUID) *DepartmentCreate {
	dc.mutation.SetID(u)
	return dc
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (dc *DepartmentCreate) AddStudentIDs(ids ...uuid.UUID) *DepartmentCreate {
	dc.mutation.AddStudentIDs(ids...)
	return dc
}

// AddStudents adds the "students" edges to the Student entity.
func (dc *DepartmentCreate) AddStudents(s ...*Student) *DepartmentCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddStudentIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the Staff entity by IDs.
func (dc *DepartmentCreate) AddStaffIDs(ids ...uuid.UUID) *DepartmentCreate {
	dc.mutation.AddStaffIDs(ids...)
	return dc
}

// AddStaffs adds the "staffs" edges to the Staff entity.
func (dc *DepartmentCreate) AddStaffs(s ...*Staff) *DepartmentCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddStaffIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	var (
		err  error
		node *Department
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := department.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := department.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := department.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "code"`)}
	}
	if v, ok := dc.mutation.Code(); ok {
		if err := department.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "code": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Telephone(); !ok {
		return &ValidationError{Name: "telephone", err: errors.New(`ent: missing required field "telephone"`)}
	}
	if v, ok := dc.mutation.Telephone(); ok {
		if err := department.TelephoneValidator(v); err != nil {
			return &ValidationError{Name: "telephone", err: fmt.Errorf(`ent: validator failed for field "telephone": %w`, err)}
		}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: department.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: department.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := dc.mutation.Telephone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldTelephone,
		})
		_node.Telephone = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := dc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.StudentsTable,
			Columns: []string{department.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.StaffsTable,
			Columns: []string{department.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: staff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
