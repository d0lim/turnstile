package domain

type Token struct {
	Sub string
	Iat int64
	Exp int64
}
