package token

import "github.com/d0lim/turnstile/internal/core/domain"

type TokenManager interface {
	IssueAccessToken(sub string) (string, error)
	IssueRefreshToken(sub string) (string, error)
	ValidateAccessToken(token string) (domain.Token, error)
	ValidateRefreshToken(token string) (domain.Token, error)
}
