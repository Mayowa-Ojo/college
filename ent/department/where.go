// Code generated by entc, DO NOT EDIT.

package department

import (
	"college/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephone), v))
	})
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelephone), v...))
	})
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelephone), v...))
	})
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephone), v))
	})
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephone), v))
	})
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephone), v))
	})
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephone), v))
	})
}

// TelephoneContains applies the Contains predicate on the "telephone" field.
func TelephoneContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTelephone), v))
	})
}

// TelephoneHasPrefix applies the HasPrefix predicate on the "telephone" field.
func TelephoneHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTelephone), v))
	})
}

// TelephoneHasSuffix applies the HasSuffix predicate on the "telephone" field.
func TelephoneHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTelephone), v))
	})
}

// TelephoneEqualFold applies the EqualFold predicate on the "telephone" field.
func TelephoneEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTelephone), v))
	})
}

// TelephoneContainsFold applies the ContainsFold predicate on the "telephone" field.
func TelephoneContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTelephone), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasStudents applies the HasEdge predicate on the "students" edge.
func HasStudents() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentsWith applies the HasEdge predicate on the "students" edge with a given conditions (other predicates).
func HasStudentsWith(preds ...predicate.Student) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStaffs applies the HasEdge predicate on the "staffs" edge.
func HasStaffs() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaffsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StaffsTable, StaffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaffsWith applies the HasEdge predicate on the "staffs" edge with a given conditions (other predicates).
func HasStaffsWith(preds ...predicate.Staff) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaffsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StaffsTable, StaffsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		p(s.Not())
	})
}
