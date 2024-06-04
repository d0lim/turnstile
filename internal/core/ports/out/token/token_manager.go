package token

import "github.com/d0lim/turnstile/internal/core/domain"

type TokenManager interface {
	IssueAccessToken(sub string) (string, error)
	IssueRefreshToken(sub string) (string, error)
	ValidateAccessToken(token string) (domain.DecodedToken, error)
	ValidateRefreshToken(token string) (domain.DecodedToken, error)
}
