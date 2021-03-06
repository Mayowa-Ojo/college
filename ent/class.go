// Code generated by entc, DO NOT EDIT.

package ent

import (
	"college/ent/class"
	"college/ent/schema"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit int `json:"unit,omitempty"`
	// Semester holds the value of the "semester" field.
	Semester schema.Semester `json:"semester,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges ClassEdges `json:"edges"`
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Student holds the value of the student edge.
	Student []*Student `json:"student,omitempty"`
	// Instructors holds the value of the instructors edge.
	Instructors []*Staff `json:"instructors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StudentOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// InstructorsOrErr returns the Instructors value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) InstructorsOrErr() ([]*Staff, error) {
	if e.loadedTypes[1] {
		return e.Instructors, nil
	}
	return nil, &NotLoadedError{edge: "instructors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldUnit:
			values[i] = new(sql.NullInt64)
		case class.FieldTitle, class.FieldCode, class.FieldSemester, class.FieldLocation:
			values[i] = new(sql.NullString)
		case class.FieldCreatedAt, class.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case class.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Class", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case class.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case class.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		case class.FieldUnit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				c.Unit = int(value.Int64)
			}
		case class.FieldSemester:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field semester", values[i])
			} else if value.Valid {
				c.Semester = schema.Semester(value.String)
			}
		case class.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				c.Location = value.String
			}
		case class.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case class.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryStudent queries the "student" edge of the Class entity.
func (c *Class) QueryStudent() *StudentQuery {
	return (&ClassClient{config: c.config}).QueryStudent(c)
}

// QueryInstructors queries the "instructors" edge of the Class entity.
func (c *Class) QueryInstructors() *StaffQuery {
	return (&ClassClient{config: c.config}).QueryInstructors(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return (&ClassClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", title=")
	builder.WriteString(c.Title)
	builder.WriteString(", code=")
	builder.WriteString(c.Code)
	builder.WriteString(", unit=")
	builder.WriteString(fmt.Sprintf("%v", c.Unit))
	builder.WriteString(", semester=")
	builder.WriteString(fmt.Sprintf("%v", c.Semester))
	builder.WriteString(", location=")
	builder.WriteString(c.Location)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class

func (c Classes) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
