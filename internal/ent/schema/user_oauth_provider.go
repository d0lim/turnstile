package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserOauthProvider holds the schema definition for the UserOauthProvider entity.
type UserOauthProvider struct {
	ent.Schema
}

// Fields of the UserOauthProvider.
func (UserOauthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("user_id"),
		field.String("oauth_provider").NotEmpty(),
		field.String("oauth_user_id").NotEmpty(),
	}
}
