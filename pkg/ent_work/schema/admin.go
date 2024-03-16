package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Admin struct {
	ent.Schema
}

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

func (Admin) Edges() []ent.Edge {
	return nil
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
