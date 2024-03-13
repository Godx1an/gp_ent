package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type School struct {
	ent.Schema
}

// Fields of the User.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
	}
}

// Edges of the User.
func (School) Edges() []ent.Edge {
	return nil
}

func (School) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
