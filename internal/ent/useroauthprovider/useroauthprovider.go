// Code generated by ent, DO NOT EDIT.

package useroauthprovider

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the useroauthprovider type in the database.
	Label = "user_oauth_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldOauthProvider holds the string denoting the oauth_provider field in the database.
	FieldOauthProvider = "oauth_provider"
	// FieldOauthUserID holds the string denoting the oauth_user_id field in the database.
	FieldOauthUserID = "oauth_user_id"
	// Table holds the table name of the useroauthprovider in the database.
	Table = "user_oauth_providers"
)

// Columns holds all SQL columns for useroauthprovider fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldOauthProvider,
	FieldOauthUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OauthProviderValidator is a validator for the "oauth_provider" field. It is called by the builders before save.
	OauthProviderValidator func(string) error
	// OauthUserIDValidator is a validator for the "oauth_user_id" field. It is called by the builders before save.
	OauthUserIDValidator func(string) error
)

// OrderOption defines the ordering options for the UserOauthProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByOauthProvider orders the results by the oauth_provider field.
func ByOauthProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauthProvider, opts...).ToFunc()
}

// ByOauthUserID orders the results by the oauth_user_id field.
func ByOauthUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauthUserID, opts...).ToFunc()
}
