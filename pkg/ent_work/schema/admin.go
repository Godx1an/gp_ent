package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Admin struct {
	ent.Schema
}

// Fields of the User.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").
			Unique().
			MinLen(11).
			MaxLen(11),
		field.String("nickname"),
		field.String("password"),
		field.String("school").Default(""),
	}
}

// Edges of the User.
func (Admin) Edges() []ent.Edge {
	return nil
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
