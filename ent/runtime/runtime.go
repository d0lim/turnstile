// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/d0lim/turnstile/ent/schema"
	"github.com/d0lim/turnstile/ent/user"
	"github.com/d0lim/turnstile/ent/useroauthprovider"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	useroauthproviderFields := schema.UserOauthProvider{}.Fields()
	_ = useroauthproviderFields
	// useroauthproviderDescOauthProvider is the schema descriptor for oauth_provider field.
	useroauthproviderDescOauthProvider := useroauthproviderFields[2].Descriptor()
	// useroauthprovider.OauthProviderValidator is a validator for the "oauth_provider" field. It is called by the builders before save.
	useroauthprovider.OauthProviderValidator = useroauthproviderDescOauthProvider.Validators[0].(func(string) error)
	// useroauthproviderDescOauthUserID is the schema descriptor for oauth_user_id field.
	useroauthproviderDescOauthUserID := useroauthproviderFields[3].Descriptor()
	// useroauthprovider.OauthUserIDValidator is a validator for the "oauth_user_id" field. It is called by the builders before save.
	useroauthprovider.OauthUserIDValidator = useroauthproviderDescOauthUserID.Validators[0].(func(string) error)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)