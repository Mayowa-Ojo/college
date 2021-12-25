package schema

import "entgo.io/ent"

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return nil
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return nil
}
