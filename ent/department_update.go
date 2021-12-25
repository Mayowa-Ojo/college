// Code generated by entc, DO NOT EDIT.

package ent

import (
	"college/ent/department"
	"college/ent/predicate"
	"college/ent/student"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// Where appends a list predicates to the DepartmentUpdate builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DepartmentUpdate) SetName(s string) *DepartmentUpdate {
	du.mutation.SetName(s)
	return du
}

// SetCode sets the "code" field.
func (du *DepartmentUpdate) SetCode(s string) *DepartmentUpdate {
	du.mutation.SetCode(s)
	return du
}

// SetTelephone sets the "telephone" field.
func (du *DepartmentUpdate) SetTelephone(s string) *DepartmentUpdate {
	du.mutation.SetTelephone(s)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DepartmentUpdate) SetCreatedAt(t time.Time) *DepartmentUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableCreatedAt(t *time.Time) *DepartmentUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DepartmentUpdate) SetUpdatedAt(t time.Time) *DepartmentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (du *DepartmentUpdate) AddStudentIDs(ids ...uuid.UUID) *DepartmentUpdate {
	du.mutation.AddStudentIDs(ids...)
	return du
}

// AddStudents adds the "students" edges to the Student entity.
func (du *DepartmentUpdate) AddStudents(s ...*Student) *DepartmentUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddStudentIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (du *DepartmentUpdate) ClearStudents() *DepartmentUpdate {
	du.mutation.ClearStudents()
	return du
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (du *DepartmentUpdate) RemoveStudentIDs(ids ...uuid.UUID) *DepartmentUpdate {
	du.mutation.RemoveStudentIDs(ids...)
	return du
}

// RemoveStudents removes "students" edges to Student entities.
func (du *DepartmentUpdate) RemoveStudents(s ...*Student) *DepartmentUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveStudentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DepartmentUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := department.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DepartmentUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := du.mutation.Code(); ok {
		if err := department.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := du.mutation.Telephone(); ok {
		if err := department.TelephoneValidator(v); err != nil {
			return &ValidationError{Name: "telephone", err: fmt.Errorf("ent: validator failed for field \"telephone\": %w", err)}
		}
	}
	return nil
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: department.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldName,
		})
	}
	if value, ok := du.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldCode,
		})
	}
	if value, ok := du.mutation.Telephone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldTelephone,
		})
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldUpdatedAt,
		})
	}
	if du.mutation.StudentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !du.mutation.StudentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.StudentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DepartmentMutation
}

// SetName sets the "name" field.
func (duo *DepartmentUpdateOne) SetName(s string) *DepartmentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetCode sets the "code" field.
func (duo *DepartmentUpdateOne) SetCode(s string) *DepartmentUpdateOne {
	duo.mutation.SetCode(s)
	return duo
}

// SetTelephone sets the "telephone" field.
func (duo *DepartmentUpdateOne) SetTelephone(s string) *DepartmentUpdateOne {
	duo.mutation.SetTelephone(s)
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DepartmentUpdateOne) SetCreatedAt(t time.Time) *DepartmentUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableCreatedAt(t *time.Time) *DepartmentUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DepartmentUpdateOne) SetUpdatedAt(t time.Time) *DepartmentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (duo *DepartmentUpdateOne) AddStudentIDs(ids ...uuid.UUID) *DepartmentUpdateOne {
	duo.mutation.AddStudentIDs(ids...)
	return duo
}

// AddStudents adds the "students" edges to the Student entity.
func (duo *DepartmentUpdateOne) AddStudents(s ...*Student) *DepartmentUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddStudentIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (duo *DepartmentUpdateOne) ClearStudents() *DepartmentUpdateOne {
	duo.mutation.ClearStudents()
	return duo
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (duo *DepartmentUpdateOne) RemoveStudentIDs(ids ...uuid.UUID) *DepartmentUpdateOne {
	duo.mutation.RemoveStudentIDs(ids...)
	return duo
}

// RemoveStudents removes "students" edges to Student entities.
func (duo *DepartmentUpdateOne) RemoveStudents(s ...*Student) *DepartmentUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveStudentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DepartmentUpdateOne) Select(field string, fields ...string) *DepartmentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Department entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {
	var (
		err  error
		node *Department
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DepartmentUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := department.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DepartmentUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Code(); ok {
		if err := department.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Telephone(); ok {
		if err := department.TelephoneValidator(v); err != nil {
			return &ValidationError{Name: "telephone", err: fmt.Errorf("ent: validator failed for field \"telephone\": %w", err)}
		}
	}
	return nil
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (_node *Department, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: department.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Department.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for _, f := range fields {
			if !department.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldName,
		})
	}
	if value, ok := duo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldCode,
		})
	}
	if value, ok := duo.mutation.Telephone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldTelephone,
		})
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: department.FieldUpdatedAt,
		})
	}
	if duo.mutation.StudentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !duo.mutation.StudentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.StudentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Department{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
