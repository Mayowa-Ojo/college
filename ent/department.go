// Code generated by entc, DO NOT EDIT.

package ent

import (
	"college/ent/department"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Telephone holds the value of the "telephone" field.
	Telephone string `json:"telephone,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges DepartmentEdges `json:"edges"`
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// Staffs holds the value of the staffs edge.
	Staffs []*Staff `json:"staffs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) StaffsOrErr() ([]*Staff, error) {
	if e.loadedTypes[1] {
		return e.Staffs, nil
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldName, department.FieldCode, department.FieldTelephone:
			values[i] = new(sql.NullString)
		case department.FieldCreatedAt, department.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case department.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Department", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case department.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				d.Code = value.String
			}
		case department.FieldTelephone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field telephone", values[i])
			} else if value.Valid {
				d.Telephone = value.String
			}
		case department.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case department.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryStudents queries the "students" edge of the Department entity.
func (d *Department) QueryStudents() *StudentQuery {
	return (&DepartmentClient{config: d.config}).QueryStudents(d)
}

// QueryStaffs queries the "staffs" edge of the Department entity.
func (d *Department) QueryStaffs() *StaffQuery {
	return (&DepartmentClient{config: d.config}).QueryStaffs(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return (&DepartmentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteString(", code=")
	builder.WriteString(d.Code)
	builder.WriteString(", telephone=")
	builder.WriteString(d.Telephone)
	builder.WriteString(", created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department

func (d Departments) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
