package schema

import "entgo.io/ent"

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return nil
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return nil
}
