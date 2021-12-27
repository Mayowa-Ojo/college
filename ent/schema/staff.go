package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Unique().Immutable(),
		field.String("firstname").MaxLen(250),
		field.String("lastname").MaxLen(250),
		field.String("email").Match(regexp.MustCompile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)).Unique(),
		field.String("telephone").MaxLen(15),
		field.Int("salary").NonNegative(),
		field.Enum("role").GoType(StaffRole("")),
		field.Enum("rank").GoType(StaffRank("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).
			Ref("staffs").
			Unique(),
		edge.From("classes", Class.Type).Ref("instructors"),
	}
}

// Enum implementation with custiom type
type (
	StaffRank string
	StaffRole string
)

const (
	LecturerI         StaffRank = "Lecturer-I"
	LecturerII        StaffRank = "Lecturer-II"
	SeniorLecturerI   StaffRank = "Senior-Lecturer-I"
	SeniorLecturerII  StaffRank = "Senior-Lecturer-II"
	PrincipalLecturer StaffRank = "Principal-Lecturer"
	HOD               StaffRank = "HOD"
)

const (
	Academic    StaffRole = "Academic-Staff"
	NonAcademic StaffRole = "Non-Academic-Staff"
)

func (StaffRank) Values() (kinds []string) {
	for _, s := range []StaffRank{
		LecturerI, LecturerII,
		SeniorLecturerI, SeniorLecturerII,
		PrincipalLecturer, HOD,
	} {
		kinds = append(kinds, string(s))
	}

	return
}

func (StaffRole) Values() (kinds []string) {
	for _, s := range []StaffRole{Academic, NonAcademic} {
		kinds = append(kinds, string(s))
	}

	return
}
