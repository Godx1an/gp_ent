package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("请设置昵称").StructTag(`json:"name"`).Comment("用户名"),
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("头像"),
	}
}

// Mixin of User
func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户表"),
	}
}

// Indexes of User
func (Student) Indexes() []ent.Index {
	return []ent.Index{}
}
