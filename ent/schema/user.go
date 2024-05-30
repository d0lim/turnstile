package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("oauth_id").NotEmpty(),
		field.String("oauth_provider").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("profile_image_url").Optional(),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Mixin for the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimestampMixin{},
	}
}
