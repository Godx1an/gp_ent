package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").
			Unique().
			MinLen(11).
			MaxLen(11),
		field.String("nickname"),
		field.String("password"),
		field.String("school").Default(""),
		field.Time("next_update_time").Optional(),
		field.String("email").Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
