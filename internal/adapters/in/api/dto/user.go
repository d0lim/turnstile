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
