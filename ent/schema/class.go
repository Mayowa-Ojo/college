package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Unique().Immutable(),
		field.String("title").MaxLen(250),
		field.String("code").MaxLen(8),
		field.Int("unit").Max(5).Min(1),
		field.Enum("semester").GoType(Semester("")),
		field.String("location").MaxLen(250),
		// instructor
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).
			StorageKey(
				// set join-table and column names
				edge.Table("class_student"),
				edge.Columns("class_id", "student_id"),
			),
	}
}

// Enum implementation with custiom type
type Semester string

const (
	FirstSemester  Semester = "FIRST"
	SecondSemester Semester = "SECOND"
)

func (Semester) Values() (kinds []string) {
	for _, s := range []Semester{FirstSemester, SecondSemester} {
		kinds = append(kinds, string(s))
	}

	return
}
