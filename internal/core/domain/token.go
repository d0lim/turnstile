package domain

type Token struct {
	Sub int64
	Iat int64
	Exp int64
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
