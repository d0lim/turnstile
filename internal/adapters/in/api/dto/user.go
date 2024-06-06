package dto

type RedirectUriResponse struct {
	RedirectUri string `json:"redirect_uri"`
}

type GoogleUserResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type InternalUserResponse struct {
	ID              int64   `json:"id"`
	OAuthProvider   string  `json:"oauth_provider"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	ProfileImageUrl *string `json:"profile_image_url"`
}
