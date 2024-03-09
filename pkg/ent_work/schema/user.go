package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("请设置昵称").StructTag(`json:"name"`).Comment("用户名"),
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("头像"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("用户的手机号"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.Bool("is_frozen").Default(false).StructTag(`json:"is_frozen"`).Comment("是否冻结"),
		field.Bool("is_recharge").Default(false).StructTag(`json:"is_recharge"`).Comment("是否充值过"),
		field.Enum("user_type").Values("personal", "enterprise").Default("personal").StructTag(`json:"user_type"`).Comment("用户类型"),
		field.String("pop_version").Default("").StructTag(`json:"pop_version"`).Comment("用户最新弹窗版本"),
		field.String("area_code").Default("+86").StructTag(`json:"area_code"`).Comment("国家区号"),
		field.String("email").Default("").StructTag(`json:"email"'`).Comment("邮箱"),
		field.Int64("cloud_space").Default(0).StructTag(`json:"cloud_space"`).Comment("云盘空间"),
	}
}

// Mixin of User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户表"),
	}
}

// Indexes of User
func (User) Indexes() []ent.Index {
	return []ent.Index{}
}
