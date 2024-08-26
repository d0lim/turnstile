package domain

type User struct {
	ID              int64
	OAuthId         string
	OAuthProvider   string
	Name            string
	Email           string
	ProfileImageUrl *string
}
