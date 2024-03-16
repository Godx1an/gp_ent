package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type FitnessTestItem struct {
	ent.Schema
}

func (FitnessTestItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("item").
			Unique(),
	}
}

func (FitnessTestItem) Edges() []ent.Edge {
	return nil
}

func (FitnessTestItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
