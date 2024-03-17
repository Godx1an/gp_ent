package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SchoolFitnessTestItem holds the schema definition for the SchoolFitnessTestItem entity.
type SchoolFitnessTestItem struct {
	ent.Schema
}

// Fields of the SchoolFitnessTestItem.
func (SchoolFitnessTestItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("max_participants").
			Default(0),
		field.Int("avg_time_per_person").
			Default(0),
		field.Int64("school_id").
			Default(0),
		field.Int64("item_id").
			Default(0),
		field.String("school").
			Default(""),
		field.String("item").
			Default(""),
	}
}

// Edges of the SchoolFitnessTestItem.
func (SchoolFitnessTestItem) Edges() []ent.Edge {
	return nil
}

func (SchoolFitnessTestItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
