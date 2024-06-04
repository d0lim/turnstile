package token

import "github.com/d0lim/turnstile/internal/core/domain"

type TokenManager interface {
	IssueAccessToken(sub string) (string, error)
	IssueRefreshToken(sub string) (string, error)
	VerifyAccessToken(token string) (domain.Token, error)
	VerifyRefreshToken(token string) (domain.Token, error)
}
